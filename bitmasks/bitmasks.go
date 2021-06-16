package bitmasks

var ErrorCodes = struct {
	GENERIC_ERROR             int
	AUTHENTICATION_ERROR      int
	INVALID_XML               int
	INVALID_METHOD            int
	RESULT_ID_NOT_IN_CACHE    int
	INVALID_MISSING_PARAMETER int
	BOOKING_VALIDATION_FAILED int
	PERMISSIONS_DENIED        int
}{
	GENERIC_ERROR:             1,
	AUTHENTICATION_ERROR:      2,
	INVALID_XML:               4,
	INVALID_METHOD:            8,
	RESULT_ID_NOT_IN_CACHE:    16,
	INVALID_MISSING_PARAMETER: 32,
	BOOKING_VALIDATION_FAILED: 64,
	PERMISSIONS_DENIED:        128,
}

var HotelDetails = struct {
	BASIC_INFO                         int
	THEMES                             int
	HOTEL_FACILITIES                   int
	SHORT_DESCRIPTION                  int
	FULL_DESCRIPTION                   int
	GEOGRAPHIC_INFORMATION             int
	COORDINATES                        int
	ADDRESS                            int
	CONTACTS                           int
	PAYMENT_OPTIONS_FOR_ONLINE_BOOKING int
	PAYMENT_OPTIONS_AT_HOTEL           int
	LOGO                               int
	HEADER_IMAGES                      int
	GALLERY                            int
	HOTEL_MATCHING                     int
	GEOGRAPHICAL_INFORMATION_AS_TEXT   int
	HOTEL_NAVIGATOR_DATA               int
	DETAILED_HOTEL_FACILITIES          int
	SALES_POINT                        int
	LTS_SPECIFIC_PARAMETERS            int
	CHECK_IN_OUT                       int
	SOURCE_DATA                        int
	BOARD_DATA                         int
	COUPON_SERVICE_DATA                int
}{
	BASIC_INFO:                         1,
	THEMES:                             2,
	HOTEL_FACILITIES:                   4,
	SHORT_DESCRIPTION:                  8,
	FULL_DESCRIPTION:                   16,
	GEOGRAPHIC_INFORMATION:             32,
	COORDINATES:                        64,
	ADDRESS:                            128,
	CONTACTS:                           256,
	PAYMENT_OPTIONS_FOR_ONLINE_BOOKING: 512,
	PAYMENT_OPTIONS_AT_HOTEL:           1024,
	LOGO:                               2048,
	HEADER_IMAGES:                      4096,
	GALLERY:                            8192,
	HOTEL_MATCHING:                     16384,
	GEOGRAPHICAL_INFORMATION_AS_TEXT:   32768,
	HOTEL_NAVIGATOR_DATA:               65536,
	DETAILED_HOTEL_FACILITIES:          131072,
	SALES_POINT:                        524288,
	LTS_SPECIFIC_PARAMETERS:            262144,
	CHECK_IN_OUT:                       1048576,
	SOURCE_DATA:                        2097152,
	BOARD_DATA:                         8388608,
	COUPON_SERVICE_DATA:                16777216,
}

var Method = struct {
	GetLocationList string
	GetHotelList    string
}{
	GetLocationList: "getLocationList",
	GetHotelList:    "getHotelList",
}

var OfferDetails = struct {
	BASIC_INFO              int
	ROOM_CODE               int
	ROOM_TITLE              int
	PRICE_DETAILS           int
	ROOM_IMAGES             int
	ROOM_FACILITIES_FILTER  int
	ROOM_DESCRIPTION        int
	INCLUDED_SERVICES       int
	ADDITIONAL_SERVICES     int
	ROOM_FACILITIES_DETAILS int
	PRICE_IMAGES            int
	THEMES                  int
	ROOM_FEATURES           int
	CANCEL_POLICIES         int
	PAYMENT_TERMS           int
}{
	BASIC_INFO:              1,
	ROOM_CODE:               4,
	ROOM_TITLE:              8,
	PRICE_DETAILS:           16,
	ROOM_IMAGES:             32,
	ROOM_FACILITIES_FILTER:  64,
	ROOM_DESCRIPTION:        256,
	INCLUDED_SERVICES:       1024,
	ADDITIONAL_SERVICES:     2048,
	ROOM_FACILITIES_DETAILS: 4096,
	PRICE_IMAGES:            8192,
	THEMES:                  16384,
	ROOM_FEATURES:           32768,
	CANCEL_POLICIES:         262144,
	PAYMENT_TERMS:           1048576,
}

var PriceListDetails = struct {
	BASE_DATA int
	HEADLINES int
	SEASONS   int
}{
	BASE_DATA: 1,
	HEADLINES: 8,
	SEASONS:   4194304,
}

var RoomDetails = struct {
	BASIC_INFO              int
	TITLE                   int
	ROOM_IMAGES             int
	ROOM_FACILITIES_FILTER  int
	ROOM_DESCRIPTION        int
	ROOM_FACILITIES_DETAILS int
	ROOM_FEATURES           int
	ROOM_NUMBERS            int
}{
	BASIC_INFO:              4,
	TITLE:                   8,
	ROOM_IMAGES:             32,
	ROOM_FACILITIES_FILTER:  64,
	ROOM_DESCRIPTION:        256,
	ROOM_FACILITIES_DETAILS: 4096,
	ROOM_FEATURES:           32768,
	ROOM_NUMBERS:            65536,
}

var SeoDetails = struct {
	PICTURES int
}{
	PICTURES: 32,
}

var SpecialDetails = struct {
	BASIC_INFO               int
	TITLE                    int
	DESCRIPTIONS             int
	SEASONS                  int
	IMAGES                   int
	THEMES                   int
	INCLUDED_SERVICES        int
	HOTEL_INCLUDED_SERVICES  int
	HOTEL_MANDATORY_SERVICES int
}{
	BASIC_INFO:               1,
	TITLE:                    2,
	DESCRIPTIONS:             4,
	SEASONS:                  8,
	IMAGES:                   16,
	THEMES:                   32,
	INCLUDED_SERVICES:        64,
	HOTEL_INCLUDED_SERVICES:  128,
	HOTEL_MANDATORY_SERVICES: 256,
}

var SpecialPremium = struct {
	VITALPINA                         int
	FAMILY_HOTELS_PREMIUM             int
	VINUM_HOTELS_PREMIUM              int
	SÜDTIROL_BALANCE_PREMIUM          int
	VITALPINA_DURCHATMEN              int
	VITALPINA_WOHLFÜHLEN              int
	VITAPLINA_ERNÄHRUNG               int
	VITAPLINA_AKTIV                   int
	VITALPINA_PREMIUM                 int
	BIKEHOTELS_MOUNTAINBIKE           int
	BIKEHOTELS_BIKE_TOURING_AND_EBIKE int
	BIKEHOTELS_ROADBIKE               int
	BIKEHOTELS_PREMIUM                int
}{
	VITALPINA:                         1,
	FAMILY_HOTELS_PREMIUM:             2,
	VINUM_HOTELS_PREMIUM:              4,
	SÜDTIROL_BALANCE_PREMIUM:          8,
	VITALPINA_DURCHATMEN:              16,
	VITALPINA_WOHLFÜHLEN:              32,
	VITAPLINA_ERNÄHRUNG:               64,
	VITAPLINA_AKTIV:                   128,
	VITALPINA_PREMIUM:                 256,
	BIKEHOTELS_MOUNTAINBIKE:           512,
	BIKEHOTELS_BIKE_TOURING_AND_EBIKE: 1024,
	BIKEHOTELS_ROADBIKE:               2048,
	BIKEHOTELS_PREMIUM:                4096,
}

var SpecialType = struct {
	PACKAGES       int
	SPECIALS       int
	MASTERPACKAGES int
}{
	PACKAGES:       1,
	SPECIALS:       2,
	MASTERPACKAGES: 4,
}
