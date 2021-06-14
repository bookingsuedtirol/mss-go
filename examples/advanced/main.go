package main

import (
	"fmt"
	"os"
	"time"

	"github.com/hgv/mss-go"
	"github.com/hgv/mss-go/request"
)

func main() {
	client := mss.Client{
		User:     os.Getenv("MSS_USER"),
		Password: os.Getenv("MSS_PASSWORD"),
		Source:   os.Getenv("MSS_SOURCE"),
	}

	responseRoot := client.Request(func(requestRoot request.Root) request.Root {
		requestRoot.Header.Method = request.Method.GetHotelList
		requestRoot.Request = request.Request{
			Search: request.Search{
				Id: []int{9002},
				SearchOffer: request.SearchOffer{
					Arrival:   request.Date(time.Now()),
					Departure: request.Date(time.Now().AddDate(0, 0, 7)),
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
			Options: request.Options{
				OfferDetails: request.OfferDetails.BASIC_INFO |
					request.OfferDetails.ROOM_TITLE |
					request.OfferDetails.CANCEL_POLICIES |
					request.OfferDetails.PAYMENT_TERMS,
			},
		}

		return requestRoot
	})

	hotel := responseRoot.Result.Hotel[0]

	fmt.Printf("%+v\n", hotel)
}
