package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/bookingsuedtirol/mss-go"
	"github.com/bookingsuedtirol/mss-go/request"
)

func main() {
	client := mss.NewClient(
		&http.Client{
			Timeout: 20 * time.Second,
		},
		mss.Credentials{
			User:     os.Getenv("MSS_USER"),
			Password: os.Getenv("MSS_PASSWORD"),
			Source:   os.Getenv("MSS_SOURCE"),
		},
	)

	responseRoot, err := client.Request(context.Background(),
		func(requestRoot request.Root) request.Root {
			requestRoot.Header.Method = request.MethodGetHotelList
			requestRoot.Request = request.Request{
				Search: &request.Search{
					IDs: []int{11230},
				},
				Options: &request.Options{
					HotelDetails: request.HotelDetailsBasicInfo |
						request.HotelDetailsCoordinates,
				},
			}

			return requestRoot
		},
	)

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
