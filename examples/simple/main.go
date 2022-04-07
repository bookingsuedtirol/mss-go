package main

import (
	"fmt"
	"os"

	"github.com/HGV/mss-go"
	"github.com/HGV/mss-go/request"
	"github.com/HGV/mss-go/types/hoteldetails"
	"github.com/HGV/mss-go/types/method"
)

func main() {
	client := mss.NewClient(mss.Credentials{
		User:     os.Getenv("MSS_USER"),
		Password: os.Getenv("MSS_PASSWORD"),
		Source:   os.Getenv("MSS_SOURCE"),
	})

	responseRoot, err := client.Request(func(requestRoot request.Root) request.Root {
		requestRoot.Header.Method = method.GetHotelList
		requestRoot.Request = request.Request{
			Search: &request.Search{
				IDs: []int{11230},
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

	hotel := responseRoot.Result.Hotels[0]

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
