package main

import (
	"fmt"
	"os"
	"time"

	"github.com/HGV/mss-go"
	"github.com/HGV/mss-go/request"
	"github.com/HGV/mss-go/shared"
	"github.com/HGV/mss-go/types/method"
	"github.com/HGV/mss-go/types/offer_details"
)

func main() {
	client := mss.Client{
		User:     os.Getenv("MSS_USER"),
		Password: os.Getenv("MSS_PASSWORD"),
		Source:   os.Getenv("MSS_SOURCE"),
	}

	today := shared.Date(time.Now())
	oneWeekFromNow := shared.Date(time.Now().AddDate(0, 0, 7))

	responseRoot, err := client.Request(func(requestRoot request.Root) request.Root {
		requestRoot.Header.Method = method.GetHotelList
		requestRoot.Request = request.Request{
			Search: &request.Search{
				IDs: []int{9002},
				SearchOffer: &request.SearchOffer{
					Arrival:   &today,
					Departure: &oneWeekFromNow,
					Service:   0,
					Rooms: []request.Room{
						{
							RoomSeq:  1,
							RoomType: 0,
							Persons:  []int{18, 18},
						},
					},
				},
			},
			Options: &request.Options{
				OfferDetails: offer_details.BasicInfo |
					offer_details.RoomTitle |
					offer_details.CancelPolicies |
					offer_details.PaymentTerms,
			},
		}

		return requestRoot
	})

	if err != nil {
		panic(err)
	}

	hotel := responseRoot.Result.Hotels[0]

	fmt.Printf("%+v\n", hotel)
}
