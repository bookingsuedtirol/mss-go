package mss

import (
	"os"
	"testing"

	"github.com/hgv/mss-go/bitmasks"
	"github.com/hgv/mss-go/request"
	"github.com/hgv/mss-go/response"
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
	responseRoot := client.Request(func(requestRoot request.Root) request.Root {
		requestRoot.Header.Method = bitmasks.Method.GetHotelList
		requestRoot.Request = request.Request{
			Search: &request.Search{
				Id: []int{9002},
			},
			Options: &request.Options{
				HotelDetails: bitmasks.HotelDetails.BASIC_INFO |
					bitmasks.HotelDetails.COORDINATES,
			},
		}

		return requestRoot
	})

	want := response.TrimmedString("Testhotel Webseitentool HGV (hotelhgv.it)")
	got := responseRoot.Result.Hotel[0].Name

	if got != want {
		t.Errorf("failed, want \"%v\", got \"%v\"", want, got)
	}
}
