package hoteldetails

import "github.com/HGV/mss-go/request"

const (
	BasicInfo request.HotelDetails = 1 << iota
	Themes
	HotelFacilities
	ShortDescription
	FullDescription
	GeographicInformation
	Coordinates
	Address
	Contacts
	PaymentOptionsForOnlineBooking
	PaymentOptionsAtHotel
	Logo
	HeaderImages
	Gallery
	HotelMatching
	GeographicalInformationAsText
	HotelNavigatorData
	DetailedHotelFacilities
	LTSSpecificParameters
	SalesPoint
	CheckInOut
	SourceData
	BoardData request.HotelDetails = 2 << iota
	CouponServiceData
)
