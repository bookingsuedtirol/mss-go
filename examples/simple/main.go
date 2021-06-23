package main

import (
	"fmt"
	"os"

	"github.com/HGV/mss-go"
	"github.com/HGV/mss-go/bitmasks"
	"github.com/HGV/mss-go/request"
)

func main() {
	client := mss.Client{
		User:     os.Getenv("MSS_USER"),
		Password: os.Getenv("MSS_PASSWORD"),
		Source:   os.Getenv("MSS_SOURCE"),
	}

	responseRoot, err := client.Request(func(requestRoot request.Root) request.Root {
		requestRoot.Header.Method = bitmasks.Method.GetHotelList
		requestRoot.Request = request.Request{
			Search: &request.Search{
				Id: []int{11230},
			},
			Options: &request.Options{
				HotelDetails: bitmasks.HotelDetails.BASIC_INFO |
					bitmasks.HotelDetails.COORDINATES,
			},
		}

		return requestRoot
	})

	if err != nil {
		panic(err)
	}

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
