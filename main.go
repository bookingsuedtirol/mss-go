package mss

import (
	"bytes"
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"unicode"

	"github.com/HGV/mss-go/request"
	"github.com/HGV/mss-go/response"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type Client struct {
	httpClient  http.Client
	credentials Credentials
}

type Credentials struct {
	User     string
	Password string
	Source   string
}

// NewClient creates a new client for requests to MSS.
// Make sure to pass an http.Client with a reasonable timeout, e.g. 10–20 seconds.
func NewClient(h http.Client, c Credentials) Client {
	return Client{h, c}
}

type Callback func(request.Root) request.Root

func (c Client) Request(ctx context.Context, cb Callback) (*response.Root, error) {
	body, err := c.RequestXML(ctx, cb)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	return c.decodeXMLResponse(body)
}

// RequestXML takes a callback to construct the request root,
// marshals it to XML, sends it to MSS and outputs the
// MSS XML response body as io.ReadCloser.
func (c Client) RequestXML(
	ctx context.Context, cb Callback,
) (io.ReadCloser, error) {
	reqRoot := c.getRequestRoot(cb)

	reqXML, err := xml.Marshal(reqRoot)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		"https://easychannel.it/mss/mss_service.php",
		strings.NewReader(xml.Header+string(reqXML)),
	)

	// The request needs to be closed every time because MSS does not seem to handle the
	// "Connection: Keep-Alive" header correctly.
	// For reference: https://stackoverflow.com/q/17714494, https://stackoverflow.com/a/34474535
	if req != nil {
		req.Close = true
	}

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "text/xml")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= http.StatusBadRequest {
		resp.Body.Close()
		return nil, &Error{
			Err: fmt.Errorf(
				"request to MSS failed with HTTP status code %v", resp.StatusCode,
			),
			StatusCode: http.StatusBadGateway,
		}
	}

	return resp.Body, nil
}

func (c Client) getRequestRoot(cb Callback) request.Root {
	root := request.Root{
		Version: "2.0",
		Header: request.Header{
			Credentials: request.Credentials{
				User:     c.credentials.User,
				Password: c.credentials.Password,
				Source:   c.credentials.Source,
			},
		},
	}

	newRoot := cb(root)

	if newRoot.Request.Search == nil {
		newRoot.Request.Search = &request.Search{}
	}

	// Set a default value for Lang because it’s required by the MSS
	if newRoot.Request.Search.Lang == "" {
		newRoot.Request.Search.Lang = "de"
	}

	return newRoot
}

func (c Client) decodeXMLResponse(body io.Reader) (*response.Root, error) {
	// Fix some MSS inconvienences before decoding XML.
	t := transform.NewReader(body,
		transform.Chain(
			// Normalize to Unicode NFC (MSS sometimes returns non-normalized Unicode sequences).
			norm.NFC,
			// Remove Unicode private use characters which MSS sometimes returns.
			// They shouldn’t appear in normal text.
			runes.Remove(runes.In(unicode.Co)),
		),
	)

	rawDec := xml.NewDecoder(t)
	dec := xml.NewTokenDecoder(newSpaceTrimmer(rawDec))

	var responseRoot response.Root

	if err := dec.Decode(&responseRoot); err != nil {
		return nil, err
	}

	if err := ErrorResponse(responseRoot.Header); err != nil {
		return nil, err
	}

	return &responseRoot, nil
}

// spaceTrimmer removes all leading and trailing whitespace inside XML elements.
type spaceTrimmer struct {
	dec *xml.Decoder
}

func newSpaceTrimmer(dec *xml.Decoder) spaceTrimmer {
	return spaceTrimmer{dec}
}

func (n spaceTrimmer) Token() (xml.Token, error) {
	t, err := n.dec.Token()

	if cd, ok := t.(xml.CharData); ok {
		t = xml.CharData(bytes.TrimSpace(cd))
	}
	return t, err
}

// ErrorResponse checks if MSS returned an error in its response and formats it accordingly.
func ErrorResponse(h response.Header) error {
	if h.Error.Code == 0 {
		return nil
	}

	return &Error{
		Err:        errors.New(h.Error.Message),
		Code:       h.Error.Code,
		StatusCode: mapStatusCode(h.Error.Code, h.Error.Message),
	}
}

func mapStatusCode(c response.ErrorCode, msg string) int {
	if strings.Contains(msg, "cannot cancel a cancelled booking") {
		return http.StatusBadRequest
	}

	// Handle storno_id was not found by getBooking or cancelBooking.
	if c == response.ErrorCodeInvalidMissingParameter &&
		strings.Contains(msg, "storno_id") {
		return http.StatusNotFound
	}

	switch c {
	case response.ErrorCodeInvalidXML,
		response.ErrorCodeInvalidMethod,
		response.ErrorCodeInvalidMissingParameter,
		response.ErrorCodeBookingValidationFailed,
		response.ErrorCodeMaxStayExceeded:
		return http.StatusBadRequest

	case response.ErrorCodeAuthenticationError:
		return http.StatusUnauthorized

	case response.ErrorCodePermissionsDenied:
		return http.StatusForbidden

	case response.ErrorCodeResultIDNotInCache:
		return http.StatusGone
	}

	return http.StatusBadGateway
}

type Error struct {
	Err        error
	Code       response.ErrorCode
	StatusCode int
}

func (e Error) Error() string {
	return e.Err.Error()
}

// AsError checks if the error matches mss.Error.
// If it matches, it returns the Error, otherwise it returns nil.
func AsError(err error) *Error {
	var e *Error
	errors.As(err, &e)
	return e
}
