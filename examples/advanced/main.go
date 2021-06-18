package main

import (
	"fmt"
	"os"
	"time"

	"github.com/HGV/mss-go"
	"github.com/HGV/mss-go/bitmasks"
	"github.com/HGV/mss-go/request"
	"github.com/HGV/mss-go/shared"
)

func main() {
	client := mss.Client{
		User:     os.Getenv("MSS_USER"),
		Password: os.Getenv("MSS_PASSWORD"),
		Source:   os.Getenv("MSS_SOURCE"),
	}

	today := shared.Date(time.Now())
	oneWeekFromNow := shared.Date(time.Now().AddDate(0, 0, 7))

	responseRoot := client.Request(func(requestRoot request.Root) request.Root {
		requestRoot.Header.Method = bitmasks.Method.GetHotelList
		requestRoot.Request = request.Request{
			Search: &request.Search{
				Id: []int{9002},
				SearchOffer: &request.SearchOffer{
					Arrival:   &today,
					Departure: &oneWeekFromNow,
					Service:   0,
					Room: []request.Room{
						{
							RoomSeq:  1,
							RoomType: 0,
							Person:   []int{18, 18},
						},
					},
				},
			},
			Options: &request.Options{
				OfferDetails: bitmasks.OfferDetails.BASIC_INFO |
					bitmasks.OfferDetails.ROOM_TITLE |
					bitmasks.OfferDetails.CANCEL_POLICIES |
					bitmasks.OfferDetails.PAYMENT_TERMS,
			},
		}

		return requestRoot
	})

	hotel := responseRoot.Result.Hotel[0]

	fmt.Printf("%+v\n", hotel)
}
