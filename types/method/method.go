package method

import "github.com/HGV/mss-go/request"

const (
	GetHotelList          request.Method = "getHotelList"
	GetSpecialList        request.Method = "getSpecialList"
	GetRoomList           request.Method = "getRoomList"
	GetPriceList          request.Method = "getPriceList"
	GetRoomAvailability   request.Method = "getRoomAvailability"
	PrepareBooking        request.Method = "prepareBooking"
	GetBooking            request.Method = "getBooking"
	CreateInquiry         request.Method = "createInquiry"
	GetUserSources        request.Method = "getUserSources"
	GetLocationList       request.Method = "getLocationList"
	GetMasterpackagesList request.Method = "getMasterpackagesList"
	GetThemeList          request.Method = "getThemeList"
	GetSeoTexts           request.Method = "getSeoTexts"
	ValidateCoupon        request.Method = "validateCoupon"
)
