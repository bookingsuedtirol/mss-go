package main

import (
	"fmt"
	"os"

	"github.com/hgv/mss-go"
	"github.com/hgv/mss-go/request"
)

func main() {
	settings := mss.ClientSettings{
		User:     os.Getenv("MSS_USER"),
		Password: os.Getenv("MSS_PASSWORD"),
		Source:   os.Getenv("MSS_SOURCE"),
	}

	sendRequest := mss.Client(settings)

	responseRoot := sendRequest(func(requestRoot request.Root) request.Root {
		requestRoot.Header.Method = request.Method.GetHotelList
		requestRoot.Request = request.Request{
			Search: request.Search{
				Id: []int{11230},
			},
			Options: request.Options{
				HotelDetails: request.HotelDetails.BASIC_INFO |
					request.HotelDetails.COORDINATES,
			},
		}

		return requestRoot
	})

	hotel := responseRoot.Result.Hotel[0]

	// (string) Hotel Lichtenstern
	fmt.Printf("(%T) %v\n", hotel.Name, hotel.Name)

	// (float64) 3.5
	fmt.Printf("(%T) %v\n", hotel.Stars, hotel.Stars)

	fmt.Printf(
		"(%T) %v\n",
		hotel.Geolocation.Latitude,
		hotel.Geolocation.Latitude,
	)

}
