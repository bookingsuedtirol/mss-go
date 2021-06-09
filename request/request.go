package request

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

type Credentials struct {
	User     string `xml:"user,omitempty"`
	Password string `xml:"password,omitempty"`
	Source   string `xml:"source,omitempty"`
}

type Header struct {
	Credentials Credentials `xml:"credentials,omitempty"`
	// TODO: restring values, e.g. getHotelList
	Method string `xml:"method,omitempty"`
}

type Root struct {
	Version string  `xml:"version,omitempty"`
	Header  Header  `xml:"header,omitempty"`
	Request Request `xml:"request,omitempty"`
}

type Request struct {
	Search  Search  `xml:"search,omitempty"`
	Options Options `xml:"options,omitempty"`
}

type Search struct {
	Lang        string `xml:"lang,omitempty"`
	Id          []int  `xml:"id,omitempty"`
	IdOfchannel string `xml:"id_ofchannel,omitempty"`
	RootId      int    `xml:"root_id,omitempty"`
}

type Options struct {
	HotelDetails int `xml:"hotel_details,omitempty"`
}
