package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/hgv/mss-go/request"
	"github.com/hgv/mss-go/response"
)

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

type ClientSettings struct {
	user     string
	password string
	source   string
}

func Client(settings ClientSettings) func(myFunc func(request.Root) request.Root) response.Root {
	requestRoot := request.Root{
		Version: "2.0",
		Header: request.Header{
			Credentials: request.Credentials{
				User:     settings.user,
				Password: settings.password,
				Source:   settings.source,
			},
		},
	}

	innerFunc := func(myFunc func(request.Root) request.Root) response.Root {
		transformedRequestRoot := myFunc(requestRoot)
		return sendRequest(transformedRequestRoot)
	}

	return innerFunc
}

func main() {
	settings := ClientSettings{
		user:     os.Getenv("MSS_USER"),
		password: os.Getenv("MSS_PASSWORD"),
		source:   os.Getenv("MSS_SOURCE"),
	}

	sendRequest := Client(settings)

	responseRoot := sendRequest(func(requestRoot request.Root) request.Root {
		requestRoot.Header.Method = request.Method.GetHotelList
		requestRoot.Request = request.Request{
			Search: request.Search{
				Lang: "de",
				Id:   []int{9002},
			},
			Options: request.Options{
				HotelDetails: request.HotelDetails.BASIC_INFO |
					request.HotelDetails.SHORT_DESCRIPTION,
			},
		}

		return requestRoot
	})

	fmt.Printf("%+v\n", responseRoot)
}
