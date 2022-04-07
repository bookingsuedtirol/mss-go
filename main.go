package mss

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"

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

func NewClient(c Credentials) Client {
	return Client{
		httpClient:  http.Client{Timeout: 20 * time.Second},
		credentials: c,
	}
}

func (c Client) Request(callback func(request.Root) request.Root) (*response.Root, *response.MSSError) {
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

	return c.sendRequest(transformedRequestRoot)
}

func (c Client) sendRequest(requestRoot request.Root) (*response.Root, *response.MSSError) {
	requestXMLRoot, err := xml.Marshal(requestRoot)

	if err != nil {
		return nil, &response.MSSError{Err: err}
	}

	resp, err := c.httpClient.Post(
		"https://easychannel.it/mss/mss_service.php",
		"text/xml",
		strings.NewReader(xml.Header+string(requestXMLRoot)),
	)

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
	dec := xml.NewTokenDecoder(normalizer{rawDec})

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
	dec *xml.Decoder
}

// Match Unicode Private Use Areas
var reg = regexp.MustCompile(`\p{Co}`)

// Fixes some inconveniencies of the MSS output.
// They would otherwise produce warnings in the W3C HTML validator.
// - Trims all leading and trailing whitespace.
// - Removes Unicode Private Use characters.
// - Runs text through Unicode normalization form NFC.
func (n normalizer) Token() (xml.Token, error) {
	t, err := n.dec.Token()

	if cd, ok := t.(xml.CharData); ok {
		replaced := reg.ReplaceAll(cd, []byte(""))
		trimmed := bytes.TrimSpace(replaced)
		normalized := norm.NFC.Bytes(trimmed)
		t = xml.CharData(normalized)
	}
	return t, err
}
