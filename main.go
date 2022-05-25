package mss

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/HGV/mss-go/request"
	"github.com/HGV/mss-go/response"
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

func (c Client) Request(
	ctx context.Context,
	callback func(request.Root) request.Root) (*response.Root, *response.MSSError) {
	requestRoot := request.Root{
		Version: "2.0",
		Header: request.Header{
			Credentials: request.Credentials{
				User:     c.credentials.User,
				Password: c.credentials.Password,
				Source:   c.credentials.Source,
			},
		},
	}

	transformedRequestRoot := callback(requestRoot)

	if transformedRequestRoot.Request.Search == nil {
		transformedRequestRoot.Request.Search = &request.Search{}
	}

	// Set a default value for Lang because it’s required by the MSS
	if transformedRequestRoot.Request.Search.Lang == "" {
		transformedRequestRoot.Request.Search.Lang = "de"
	}

	return c.sendRequest(ctx, transformedRequestRoot)
}

func (c Client) sendRequest(
	ctx context.Context,
	requestRoot request.Root,
) (*response.Root, *response.MSSError) {
	requestXMLRoot, err := xml.Marshal(requestRoot)

	if err != nil {
		return nil, &response.MSSError{Err: err}
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		"https://easychannel.it/mss/mss_service.php",
		strings.NewReader(xml.Header+string(requestXMLRoot)),
	)

	// The request needs to be closed every time because MSS does not seem to handle the
	// "Connection: Keep-Alive" header correctly.
	// For reference: https://stackoverflow.com/q/17714494, https://stackoverflow.com/a/34474535
	if req != nil {
		req.Close = true
	}

	if err != nil {
		return nil, &response.MSSError{Err: err}
	}

	req.Header.Set("Content-Type", "text/xml")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, &response.MSSError{Err: err}
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, &response.MSSError{
			Err: fmt.Errorf(
				"request to MSS failed with HTTP status code %v", resp.StatusCode,
			),
			StatusCode: resp.StatusCode,
		}
	}

	rawDec := xml.NewDecoder(resp.Body)
	dec := xml.NewTokenDecoder(newNormalizer(rawDec))

	var responseRoot response.Root
	err = dec.Decode(&responseRoot)

	if err != nil {
		return nil, &response.MSSError{Err: err}
	}

	if responseRoot.Header.Error.Code != 0 {
		return nil,
			&response.MSSError{
				Err: fmt.Errorf(
					"%v, code: %v",
					responseRoot.Header.Error.Message,
					responseRoot.Header.Error.Code,
				),
				Code: responseRoot.Header.Error.Code,
			}
	}

	return &responseRoot, nil
}

type normalizer struct {
	dec   *xml.Decoder
	regex *regexp.Regexp
}

func newNormalizer(dec *xml.Decoder) normalizer {
	// Match Unicode Private Use Areas
	regex := regexp.MustCompile(`\p{Co}`)
	return normalizer{dec, regex}
}

// Fixes some inconveniences of the MSS output.
// They would otherwise produce warnings in the W3C HTML validator.
// - Trims all leading and trailing whitespace.
// - Removes Unicode Private Use characters.
// - Runs text through Unicode normalization form NFC.
func (n normalizer) Token() (xml.Token, error) {
	t, err := n.dec.Token()

	if cd, ok := t.(xml.CharData); ok {
		replaced := n.regex.ReplaceAll(cd, []byte(""))
		trimmed := bytes.TrimSpace(replaced)
		normalized := norm.NFC.Bytes(trimmed)
		t = xml.CharData(normalized)
	}
	return t, err
}
