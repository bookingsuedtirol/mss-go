package request

import "encoding/xml"

var Method = struct {
	GetLocationList string
	GetHotelList    string
}{
	GetLocationList: "getLocationList",
	GetHotelList:    "getHotelList",
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

type Credentials struct {
	User     string `xml:"user"`
	Password string `xml:"password"`
	Source   string `xml:"source"`
}

type Header struct {
	Credentials Credentials `xml:"credentials"`
	// TODO: restring values, e.g. getHotelList
	Method string `xml:"method"`
}

type Root struct {
	XMLName xml.Name `xml:"root"`
	Version string   `xml:"version"`
	Header  Header   `xml:"header"`
	Request Request  `xml:"request"`
}

type Request struct {
	Search  Search  `xml:"search"`
	Options Options `xml:"options"`
}

type Search struct {
	Lang        string `xml:"lang"`
	Id          []int  `xml:"id,omitempty"`
	IdOfchannel string `xml:"id_ofchannel,omitempty"`
	RootId      int    `xml:"root_id,omitempty"`
}

type Options struct {
	HotelDetails int `xml:"hotel_details,omitempty"`
	OfferDetails int `xml:"offer_details,omitempty"`
}
