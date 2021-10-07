package mss

import (
	"os"
	"testing"

	"github.com/HGV/mss-go/request"
	"github.com/HGV/mss-go/response"
	"github.com/HGV/mss-go/types/hotel_details"
	"github.com/HGV/mss-go/types/method"
)

var client = Client{
	User:     os.Getenv("MSS_USER"),
	Password: os.Getenv("MSS_PASSWORD"),
	Source:   os.Getenv("MSS_SOURCE"),
}

func TestEnvVariablesAreDefined(t *testing.T) {
	if client.User == "" || client.Password == "" || client.Source == "" {
		t.Error("The env variables MSS_USER etc. are not set.")
	}
}

func TestSimpleMssCall(t *testing.T) {
	responseRoot, err := client.Request(func(requestRoot request.Root) request.Root {
		requestRoot.Header.Method = method.GetHotelList
		requestRoot.Request = request.Request{
			Search: &request.Search{
				Id: []int{9002},
			},
			Options: &request.Options{
				HotelDetails: hotel_details.BasicInfo |
					hotel_details.Coordinates,
			},
		}

		return requestRoot
	})

	if err != nil {
		panic(err)
	}

	want := response.TrimmedString("Testhotel Webseitentool HGV (hotelhgv.it)")
	got := responseRoot.Result.Hotel[0].Name

	if got != want {
		t.Errorf("failed, want \"%v\", got \"%v\"", want, got)
	}
}
