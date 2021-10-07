package main

import (
	"fmt"
	"os"

	"github.com/HGV/mss-go"
	"github.com/HGV/mss-go/request"
	"github.com/HGV/mss-go/types/hotel_details"
	"github.com/HGV/mss-go/types/method"
)

func main() {
	client := mss.Client{
		User:     os.Getenv("MSS_USER"),
		Password: os.Getenv("MSS_PASSWORD"),
		Source:   os.Getenv("MSS_SOURCE"),
	}

	responseRoot, err := client.Request(func(requestRoot request.Root) request.Root {
		requestRoot.Header.Method = method.GetHotelList
		requestRoot.Request = request.Request{
			Search: &request.Search{
				Id: []int{11230},
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
