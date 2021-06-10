package mss

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/hgv/mss-go/request"
	"github.com/hgv/mss-go/response"
)

type Client struct {
	User     string
	Password string
	Source   string
}

func (settings Client) Request(callback func(request.Root) request.Root) response.Root {
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

	// Set a default value for Lang because it’s required by MSS
	if transformedRequestRoot.Request.Search.Lang == "" {
		transformedRequestRoot.Request.Search.Lang = "de"
	}

	return sendRequest(transformedRequestRoot)
}

func sendRequest(requestRoot request.Root) response.Root {
	requestXmlRoot, err := xml.Marshal(requestRoot)

	if err != nil {
		panic(err)
	}

	resp, err := http.Post(
		"https://easychannel.it/mss/mss_service.php",
		"text/xml",
		strings.NewReader(xml.Header+string(requestXmlRoot)),
	)

	if err != nil {
		panic(err)
	}

	if resp.StatusCode >= 400 {
		panic(fmt.Errorf("failed to request the API\nStatusCode %v", resp.StatusCode))
	}

	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	var responseRoot response.Root
	xml.Unmarshal(responseBody, &responseRoot)

	if responseRoot.Header.Error.Code != 0 {
		panic(fmt.Errorf("the API returned an error\nCode: %v,\nMessage: %v",
			responseRoot.Header.Error.Code,
			responseRoot.Header.Error.Message,
		))
	}

	return responseRoot
}
