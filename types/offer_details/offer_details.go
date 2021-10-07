package offer_details

import "github.com/HGV/mss-go/request"

const (
	BasicInfo             request.OfferDetails = 1
	RoomCode              request.OfferDetails = 4
	RoomTitle             request.OfferDetails = 8
	PriceDetails          request.OfferDetails = 16
	RoomImages            request.OfferDetails = 32
	RoomFacilitiesFilter  request.OfferDetails = 64
	RoomDescription       request.OfferDetails = 256
	IncludedServices      request.OfferDetails = 1024
	AdditionalServices    request.OfferDetails = 2048
	RoomFacilitiesDetails request.OfferDetails = 4096
	PriceImages           request.OfferDetails = 8192
	Themes                request.OfferDetails = 16384
	RoomFeatures          request.OfferDetails = 32768
	CancelPolicies        request.OfferDetails = 262144
	PaymentTerms          request.OfferDetails = 1048576
)
