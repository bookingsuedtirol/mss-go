package mss

import (
	"os"
	"testing"

	"github.com/HGV/mss-go/request"
	"github.com/HGV/mss-go/types/hoteldetails"
	"github.com/HGV/mss-go/types/method"
)

var client = NewClient(Credentials{
	User:     os.Getenv("MSS_USER"),
	Password: os.Getenv("MSS_PASSWORD"),
	Source:   os.Getenv("MSS_SOURCE"),
})

func TestEnvVariablesAreDefined(t *testing.T) {
	if c := client.credentials; c.User == "" || c.Password == "" || c.Source == "" {
		t.Error("env variables MSS_USER etc. not set")
	}
}

func TestSimpleMssCall(t *testing.T) {
	responseRoot, err := client.Request(func(requestRoot request.Root) request.Root {
		requestRoot.Header.Method = method.GetHotelList
		requestRoot.Request = request.Request{
			Search: &request.Search{
				IDs: []int{9002},
			},
			Options: &request.Options{
				HotelDetails: hoteldetails.BasicInfo |
					hoteldetails.Coordinates,
			},
		}

		return requestRoot
	})

	if err != nil {
		panic(err)
	}

	want := "Testhotel Webseitentool HGV (hotelhgv.it)"
	got := responseRoot.Result.Hotels[0].Name

	if got != want {
		t.Errorf("failed, want \"%v\", got \"%v\"", want, got)
	}
}
