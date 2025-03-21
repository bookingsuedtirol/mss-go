package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/bookingsuedtirol/mss-go"
	"github.com/bookingsuedtirol/mss-go/request"
	"github.com/bookingsuedtirol/mss-go/shared"
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

	today := shared.Date{Time: time.Now()}
	oneWeekFromNow := shared.Date{Time: time.Now().AddDate(0, 0, 7)}

	responseRoot, err := client.Request(
		context.Background(),
		func(requestRoot request.Root) request.Root {
			requestRoot.Header.Method = request.MethodGetHotelList
			requestRoot.Request = request.Request{
				Search: &request.Search{
					IDs: []int{9002},
					Offer: &request.SearchOffer{
						Arrival:   &today,
						Departure: &oneWeekFromNow,
						Rooms: []request.Room{
							{
								Seq:     1,
								Type:    0,
								Persons: []int{18, 18},
							},
						},
					},
				},
				Options: &request.Options{
					OfferDetails: request.OfferDetailsBasicInfo |
						request.OfferDetailsRoomTitle |
						request.OfferDetailsCancelPolicies |
						request.OfferDetailsPaymentTerms,
				},
			}

			return requestRoot
		},
	)

	if err != nil {
		panic(err)
	}

	hotel := responseRoot.Result.Hotels[0]

	fmt.Printf("%+v\n", hotel)
}
