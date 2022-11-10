package request

import (
	"encoding/xml"

	"github.com/HGV/mss-go/shared"
)

// The boolean type for request data. This must be used instead of the
// standard "bool". Maps to 0/1 (which MSS expects) instead of "true"/"false"
type Bool bool

type Address struct {
	Street  string `xml:"street"`
	ZIPCode string `xml:"zipcode"`
	City    string `xml:"city"`
	// The ISO 3166 alpha-3 code
	Country string `xml:"country"`
}

type Company struct {
	Name          string   `xml:"name"`
	TaxNumber     string   `xml:"taxnumber"`
	RecipientCode string   `xml:"recipient_code"`
	Address       *Address `xml:"address"`
}

type Coupon struct {
	Code string            `xml:"coupon_code"`
	Type shared.CouponType `xml:"coupon_type"`
}

type Credentials struct {
	User     string `xml:"user"`
	Password string `xml:"password"`
	Source   string `xml:"source"`
}

type StornoReason int

const (
	StornoReasonUnknown StornoReason = iota
	StornoReasonGuestUnavailable
	StornoReasonPropertyRequestedCancellation
	StornoReasonGuestChoseAnotherDestination
	StornoReasonGuestChoseAnotherProperty
	StornoReasonOther StornoReason = 99
)

type Data struct {
	Guest            *Guest       `xml:"guest"`
	Company          *Company     `xml:"company"`
	Payment          *Payment     `xml:"payment"`
	Note             string       `xml:"note"`
	Details          *Details     `xml:"details"`
	Form             *Form        `xml:"form"`
	Tracking         *Tracking    `xml:"tracking"`
	Insurance        Bool         `xml:"insurance"`
	StornoReason     StornoReason `xml:"storno_reason"`
	StornoReasonText string       `xml:"storno_reason_text"`
}

type Details struct {
	ExtraPrices []ExtraPrice `xml:"extra_price"`
	Coupon      *Coupon      `xml:"coupon"`
}

type ExtraPrice struct {
	ID     int     `xml:"price_id"`
	Amount float64 `xml:"price_amount"`
}

type Form struct {
	URLSuccess string `xml:"url_success"`
	URLFailure string `xml:"url_failure"`
}

type Guest struct {
	Gender     shared.Gender `xml:"gender"`
	Prefix     string        `xml:"prefix"`
	FirstName  string        `xml:"firstname"`
	LastName   string        `xml:"lastname"`
	Email      string        `xml:"email"`
	Phone      string        `xml:"phone"`
	Address    *Address      `xml:"address"`
	Newsletter Bool          `xml:"newsletter"`
}

type Method string

const (
	MethodGetHotelList          Method = "getHotelList"
	MethodGetSpecialList        Method = "getSpecialList"
	MethodGetRoomList           Method = "getRoomList"
	MethodGetPriceList          Method = "getPriceList"
	MethodGetRoomAvailability   Method = "getRoomAvailability"
	MethodPrepareBooking        Method = "prepareBooking"
	MethodGetBooking            Method = "getBooking"
	MethodCancelBooking         Method = "cancelBooking"
	MethodCreateInquiry         Method = "createInquiry"
	MethodGetUserSources        Method = "getUserSources"
	MethodGetLocationList       Method = "getLocationList"
	MethodGetMasterpackagesList Method = "getMasterpackagesList"
	MethodGetThemeList          Method = "getThemeList"
	MethodValidateCoupon        Method = "validateCoupon"
)

type Header struct {
	Credentials Credentials `xml:"credentials"`
	Method      Method      `xml:"method"`
	Paging      *Paging     `xml:"paging"`
	ResultID    string      `xml:"result_id"`
}

type Logging struct {
	Step string `xml:"step"`
}

type HotelDetails int

const (
	HotelDetailsBasicInfo HotelDetails = 1 << iota
	HotelDetailsThemes
	HotelDetailsHotelFacilities
	HotelDetailsShortDescription
	HotelDetailsFullDescription
	HotelDetailsGeographicInformation
	HotelDetailsCoordinates
	HotelDetailsAddress
	HotelDetailsContacts
	HotelDetailsPaymentOptionsForOnlineBooking
	HotelDetailsPaymentOptionsAtHotel
	HotelDetailsLogo
	HotelDetailsHeaderImages
	HotelDetailsGallery
	HotelDetailsHotelMatching
	HotelDetailsGeographicalInformationAsText
	HotelDetailsHotelNavigatorData
	HotelDetailsDetailedHotelFacilities
	HotelDetailsLTSSpecificParameters
	HotelDetailsSalesPoint
	HotelDetailsCheckInOut
	HotelDetailsSourceData
	HotelDetailsBoardData HotelDetails = 2 << iota
	HotelDetailsCouponServiceData
	HotelDetailsRoomTypes HotelDetails = 8 << iota
)

type OfferDetails int

const (
	OfferDetailsBasicInfo             OfferDetails = 1
	OfferDetailsRoomCode              OfferDetails = 4
	OfferDetailsRoomTitle             OfferDetails = 8
	OfferDetailsPriceDetails          OfferDetails = 16
	OfferDetailsRoomImages            OfferDetails = 32
	OfferDetailsRoomFacilitiesFilter  OfferDetails = 64
	OfferDetailsRoomDescription       OfferDetails = 256
	OfferDetailsIncludedServices      OfferDetails = 1024
	OfferDetailsAdditionalServices    OfferDetails = 2048
	OfferDetailsRoomFacilitiesDetails OfferDetails = 4096
	OfferDetailsPriceImages           OfferDetails = 8192
	OfferDetailsThemes                OfferDetails = 16384
	OfferDetailsRoomFeatures          OfferDetails = 32768
	OfferDetailsCancelPolicies        OfferDetails = 262144
	OfferDetailsPaymentTerms          OfferDetails = 1048576
)

type RoomDetails int

const (
	RoomDetailsBasicInfo             RoomDetails = 4
	RoomDetailsTitle                 RoomDetails = 8
	RoomDetailsRoomImages            RoomDetails = 32
	RoomDetailsRoomFacilitiesFilter  RoomDetails = 64
	RoomDetailsRoomDescription       RoomDetails = 256
	RoomDetailsRoomFacilitiesDetails RoomDetails = 4096
	RoomDetailsRoomFeatures          RoomDetails = 32768
	RoomDetailsRoomNumbers           RoomDetails = 65536
)

type SpecialDetails int

const (
	SpecialDetailsBasicInfo SpecialDetails = 1 << iota
	SpecialDetailsTitle
	SpecialDetailsDescriptions
	SpecialDetailsSeasons
	SpecialDetailsImages
	SpecialDetailsThemes
	SpecialDetailsIncludedServices
	SpecialDetailsHotelIncludedServices
	SpecialDetailsHotelMandatoryServices
)

type LTSBookable int

const (
	LTSBookableOverridenByIgnoreBookability LTSBookable = iota
	LTSBookableOnlyBookableLTSOffers
	LTSBookableOnlyNonBookableLTSOffers
)

type PriceListDetails int

const (
	PriceListDetailsBaseData  PriceListDetails = 1
	PriceListDetailsHeadlines PriceListDetails = 8
	PriceListDetailsSeasons   PriceListDetails = 4194304
)

type Options struct {
	HotelDetails         HotelDetails     `xml:"hotel_details"`
	OfferDetails         OfferDetails     `xml:"offer_details"`
	RoomDetails          RoomDetails      `xml:"room_details"`
	SpecialDetails       SpecialDetails   `xml:"special_details"`
	PictureDate          *shared.Date     `xml:"picture_date"`
	LTSBookable          LTSBookable      `xml:"lts_bookable"`
	GetAvailability      Bool             `xml:"get_availability"`
	GetRestrictions      Bool             `xml:"get_restrictions"`
	GetRoomdetails       Bool             `xml:"get_roomdetails"`
	GetMasterpackages    Bool             `xml:"get_masterpackages"`
	BasePrice            Bool             `xml:"base_price"`
	PriceListDetails     PriceListDetails `xml:"pricelist_details"`
	OnlySubscribedHotels Bool             `xml:"only_subscribed_hotels"`
	OnlyAvailable        Bool             `xml:"only_available"`
	CheckRestrictions    Bool             `xml:"check_restrictions"`
}

type Order struct {
	Field OrderField `xml:"field"`
	Dir   Direction  `xml:"dir"`
}

type OrderField string

const (
	OrderFieldDate  OrderField = "date"
	OrderFieldRand  OrderField = "rand"
	OrderFieldStars OrderField = "stars"
	OrderFieldName  OrderField = "name"
)

type Direction string

const (
	DirectionAsc  Direction = "asc"
	DirectionDesc Direction = "desc"
)

type Paging struct {
	Start int `xml:"start"`
	Limit int `xml:"limit"`
}

type Payment struct {
	Method  shared.PaymentMethod `xml:"method"`
	Invoice Bool                 `xml:"invoice"`
}

type Rateplan struct {
	Code   string `xml:"code"`
	Source string `xml:"source"`
}

type Request struct {
	Search  *Search  `xml:"search"`
	Options *Options `xml:"options"`
	Order   *Order   `xml:"order"`
	Data    *Data    `xml:"data"`
	Logging *Logging `xml:"logging"`
}

type Room struct {
	OfferID int             `xml:"offer_id"`
	ID      int             `xml:"room_id"`
	Service shared.Board    `xml:"service,omitempty"`
	Type    shared.RoomType `xml:"room_type"`
	Seq     int             `xml:"room_seq"`
	Persons []int           `xml:"person"`
}

type Root struct {
	XMLName xml.Name `xml:"root"`
	Version string   `xml:"version"`
	Header  Header   `xml:"header"`
	Request Request  `xml:"request"`
}

type IDOfChannel string

const (
	IDOfChannelHGV        IDOfChannel = "hgv"
	IDOfChannelLTS        IDOfChannel = "lts"
	IDOfChannelBookingCom IDOfChannel = "bok"
	IDOfChannelHotelDe    IDOfChannel = "htl"
	IDOfChannelExpediaCom IDOfChannel = "exp"
	IDOfChannelHrsCom     IDOfChannel = "hrs"
	IDOfChannelLTSPOS     IDOfChannel = "pos"
	IDOfChannelGoogle     IDOfChannel = "gog"
)

type Search struct {
	Lang          string              `xml:"lang"`
	ResultID      string              `xml:"result_id"`
	Agent         string              `xml:"agent"`
	IDs           []int               `xml:"id"`
	Hotel         *SearchHotel        `xml:"search_hotel"`
	Location      *SearchLocation     `xml:"search_location"`
	Distance      *SearchDistance     `xml:"search_distance"`
	Offer         *SearchOffer        `xml:"search_offer"`
	LTS           *shared.LTSData     `xml:"search_lts"`
	Special       *SearchSpecial      `xml:"search_special"`
	Availability  *SearchAvailability `xml:"search_availability"`
	PriceList     *SearchPriceList    `xml:"search_pricelist"`
	In            []int               `xml:"in"`
	IDOfChannel   IDOfChannel         `xml:"id_ofchannel,omitempty"`
	TransactionID string              `xml:"transaction_id"`
	BookingID     int                 `xml:"booking_id"`
	GuestEmail    string              `xml:"guest_email"`
	RootID        int                 `xml:"root_id"`
	ExternalID    int                 `xml:"external_id"`
	Type          shared.LocationType `xml:"typ"`
	CouponCode    string              `xml:"coupon_code"`
	CouponType    *shared.CouponType  `xml:"coupon_type"`
	TotalPrice    float64             `xml:"total_price"`
	Arrival       *shared.Date        `xml:"arrival"`
	Departure     *shared.Date        `xml:"departure"`
	StornoID      string              `xml:"storno_id"`
}

type SearchAvailability struct {
	DateFrom *shared.Date     `xml:"date_from"`
	DateTo   *shared.Date     `xml:"date_to"`
	OfferIDs []int            `xml:"offer_id"`
	RoomIDs  []int            `xml:"room_id"`
	Type     shared.OfferType `xml:"typ"`
}

type SearchDistance struct {
	Latitude  float64 `xml:"latitude"`
	Longitude float64 `xml:"longitude"`
	Radius    int     `xml:"radius"`
}

type SearchHotel struct {
	Name     string              `xml:"name"`
	Types    []shared.HotelType  `xml:"type"`
	Stars    *Stars              `xml:"stars"`
	Feature  shared.HotelFeature `xml:"feature"`
	Theme    shared.Theme        `xml:"theme"`
	RoomType shared.RoomType     `xml:"room_type"`
}

type SearchLocation struct {
	Locations    []int    `xml:"location"`
	LocationsLTS []string `xml:"location_lts"`
}

type SearchOffer struct {
	Arrival    *shared.Date     `xml:"arrival"`
	Departure  *shared.Date     `xml:"departure"`
	Service    shared.Board     `xml:"service"`
	Feature    int              `xml:"feature,omitempty"`
	ChannelIDs []string         `xml:"channel_id"`
	Rooms      []Room           `xml:"room"`
	Type       shared.OfferType `xml:"typ"`
	Rateplan   *Rateplan        `xml:"rateplan"`
}

type ObjectFilter int

const (
	ObjectFilterPriceList ObjectFilter = iota
	ObjectFilterPackage
	ObjectFilterSpecial
	ObjectFilterDiscountOrSurcharge ObjectFilter = 4
)

type SearchPriceList struct {
	DateFrom *shared.Date `xml:"date_from"`
	DateTo   *shared.Date `xml:"date_to"`
	Service  shared.Board `xml:"service"`
	RoomIDs  []int        `xml:"room_id"`
	Type     ObjectFilter `xml:"typ"`
}

type SearchSpecial struct {
	OfferIDs []int                 `xml:"offer_id"`
	DateFrom *shared.Date          `xml:"date_from"`
	DateTo   *shared.Date          `xml:"date_to"`
	Themes   []shared.ThemeID      `xml:"theme"`
	POIIDs   []int                 `xml:"poi_id"`
	POICats  []int                 `xml:"poi_cat"`
	Validity *Validity             `xml:"validity"`
	Type     shared.SpecialType    `xml:"typ"`
	Premium  shared.SpecialPremium `xml:"premium"`
	Status   int                   `xml:"status"`
}

type Stars struct {
	Min float64 `xml:"min"`
	Max float64 `xml:"max"`
}

type Tracking struct {
	Partner     string `xml:"partner"`
	Media       string `xml:"media"`
	Campain     string `xml:"campain"`
	Campaign    string `xml:"campaign"`
	CompanyInfo string `xml:"companyinfo"`
}

type Validity struct {
	Valid     Bool         `xml:"valid"`
	Offers    Bool         `xml:"offers"`
	Arrival   *shared.Date `xml:"arrival"`
	Departure *shared.Date `xml:"departure"`
	Service   shared.Board `xml:"service"`
	Rooms     []Room       `xml:"room"`
}
