package mss

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/HGV/mss-go/request"
	"github.com/HGV/mss-go/response"
)

type Client struct {
	User     string
	Password string
	Source   string
}

func (settings Client) Request(callback func(request.Root) request.Root) (response.Root, error) {
	requestRoot := request.Root{
		Version: "2.0",
		Header: request.Header{
			Credentials: request.Credentials{
				User:     settings.User,
				Password: settings.Password,
				Source:   settings.Source,
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

	return sendRequest(transformedRequestRoot)
}

func sendRequest(requestRoot request.Root) (response.Root, error) {
	requestXMLRoot, err := xml.Marshal(requestRoot)

	if err != nil {
		return response.Root{}, err
	}

	resp, err := http.Post(
		"https://easychannel.it/mss/mss_service.php",
		"text/xml",
		strings.NewReader(xml.Header+string(requestXMLRoot)),
	)

	if err != nil {
		return response.Root{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return response.Root{},
			fmt.Errorf("request to MSS failed with HTTP status code %v", resp.StatusCode)
	}

	rawDec := xml.NewDecoder(resp.Body)
	dec := xml.NewTokenDecoder(normalizer{rawDec})

	var responseRoot response.Root
	err = dec.Decode(&responseRoot)

	if err != nil {
		return response.Root{}, err
	}

	if responseRoot.Header.Error.Code != 0 {
		return response.Root{},
			fmt.Errorf("MSS returned an error\nCode: %v,\nMessage: %v",
				responseRoot.Header.Error.Code,
				responseRoot.Header.Error.Message,
			)
	}

	return responseRoot, nil
}

type normalizer struct {
	dec *xml.Decoder
}

// Match Unicode Private Use Areas
var reg = regexp.MustCompile(`\p{Co}`)

// MSS could potentially output trailing spaces and private Unicode characters.
// This trims all leading and trailing whitespace inside XML elements and
// removes Unicode Private Use characters.
func (n normalizer) Token() (xml.Token, error) {
	t, err := n.dec.Token()

	if cd, ok := t.(xml.CharData); ok {
		replaced := reg.ReplaceAll(cd, []byte(""))
		trimmed := bytes.TrimSpace(replaced)
		t = xml.CharData(trimmed)
	}
	return t, err
}
