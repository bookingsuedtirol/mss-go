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
	Country string `xml:"country"`
}

type Company struct {
	Name          string   `xml:"name"`
	Taxnumber     string   `xml:"taxnumber"`
	RecipientCode string   `xml:"recipient_code"`
	Address       *Address `xml:"address"`
}

type Coupon struct {
	CouponCode string `xml:"coupon_code"`
	CouponType string `xml:"coupon_type"`
}

type Credentials struct {
	User     string `xml:"user"`
	Password string `xml:"password"`
	Source   string `xml:"source"`
}

type StornoReason int

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
	PriceID     int     `xml:"price_id"`
	PriceAmount float64 `xml:"price_amount"`
}

type Form struct {
	URLSuccess string `xml:"url_success"`
	URLFailure string `xml:"url_failure"`
}

type Guest struct {
	Gender     string   `xml:"gender"`
	Prefix     string   `xml:"prefix"`
	Firstname  string   `xml:"firstname"`
	Lastname   string   `xml:"lastname"`
	Email      string   `xml:"email"`
	Phone      string   `xml:"phone"`
	Address    *Address `xml:"address"`
	Newsletter Bool     `xml:"newsletter"`
}

type Method string

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

type OfferDetails int

type RoomDetails int

type SpecialDetails int

type SEODetails int

type PriceListDetails int

type Options struct {
	HotelDetails         HotelDetails     `xml:"hotel_details"`
	OfferDetails         OfferDetails     `xml:"offer_details"`
	RoomDetails          RoomDetails      `xml:"room_details"`
	SpecialDetails       SpecialDetails   `xml:"special_details"`
	SEODetails           SEODetails       `xml:"seo_details"`
	PictureDate          *shared.Date     `xml:"picture_date"`
	LTSBookable          int              `xml:"lts_bookable"`
	GetAvailability      Bool             `xml:"get_availability"`
	GetRestrictions      Bool             `xml:"get_restrictions"`
	GetRoomdetails       Bool             `xml:"get_roomdetails"`
	GetMasterpackages    Bool             `xml:"get_masterpackages"`
	BasePrice            Bool             `xml:"base_price"`
	PriceListDetails     PriceListDetails `xml:"pricelist_details"`
	OnlySubscribedHotels Bool             `xml:"only_subscribed_hotels"`
	OnlyAvailable        Bool             `xml:"only_available"`
}

type Order struct {
	Field OrderField `xml:"field"`
	Dir   Direction  `xml:"dir"`
}

type OrderField string

type Direction string

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
	OfferID  int             `xml:"offer_id"`
	RoomID   int             `xml:"room_id"`
	Service  shared.Board    `xml:"service,omitempty"`
	RoomType shared.RoomType `xml:"room_type"`
	RoomSeq  int             `xml:"room_seq"`
	Persons  []int           `xml:"person"`
}

type Root struct {
	XMLName xml.Name `xml:"root"`
	Version string   `xml:"version"`
	Header  Header   `xml:"header"`
	Request Request  `xml:"request"`
}

type Search struct {
	Lang               string              `xml:"lang"`
	ResultID           string              `xml:"result_id"`
	Agent              string              `xml:"agent"`
	IDs                []int               `xml:"id"`
	SearchHotel        *SearchHotel        `xml:"search_hotel"`
	SearchLocation     *SearchLocation     `xml:"search_location"`
	SearchDistance     *SearchDistance     `xml:"search_distance"`
	SearchOffer        *SearchOffer        `xml:"search_offer"`
	SearchLTS          *shared.LTSData     `xml:"search_lts"`
	SearchSpecial      *SearchSpecial      `xml:"search_special"`
	SearchAvailability *SearchAvailability `xml:"search_availability"`
	SearchPricelist    *SearchPriceList    `xml:"search_pricelist"`
	In                 []int               `xml:"in"`
	IDOfchannel        string              `xml:"id_ofchannel,omitempty"`
	TransactionID      string              `xml:"transaction_id"`
	BookingID          int                 `xml:"booking_id"`
	GuestEmail         string              `xml:"guest_email"`
	RootID             int                 `xml:"root_id"`
	ExternalID         int                 `xml:"external_id"`
	Type               shared.LocationType `xml:"typ"`
	SEOType            string              `xml:"seo_typ"`
	LocationDetails    int                 `xml:"location_details"`
	CouponCode         string              `xml:"coupon_code"`
	CouponType         string              `xml:"coupon_type"`
	TotalPrice         float64             `xml:"total_price"`
	Arrival            *shared.Date        `xml:"arrival"`
	Departure          *shared.Date        `xml:"departure"`
}

type SearchAvailability struct {
	DateFrom *shared.Date `xml:"date_from"`
	DateTo   *shared.Date `xml:"date_to"`
	OfferIDs []int        `xml:"offer_id"`
	RoomIDs  []int        `xml:"room_id"`
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

type SearchPriceList struct {
	DateFrom *shared.Date `xml:"date_from"`
	DateTo   *shared.Date `xml:"date_to"`
	Service  shared.Board `xml:"service"`
	RoomIDs  []int        `xml:"room_id"`
	Type     int          `xml:"typ"`
}

type SearchSpecial struct {
	OfferIDs []int                 `xml:"offer_id"`
	DateFrom *shared.Date          `xml:"date_from"`
	DateTo   *shared.Date          `xml:"date_to"`
	Themes   []shared.ThemeID      `xml:"theme"`
	PoiIDs   []int                 `xml:"poi_id"`
	PoiCats  []int                 `xml:"poi_cat"`
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
	Companyinfo string `xml:"companyinfo"`
}

type Validity struct {
	Valid     Bool         `xml:"valid"`
	Offers    Bool         `xml:"offers"`
	Arrival   *shared.Date `xml:"arrival"`
	Departure *shared.Date `xml:"departure"`
	Service   shared.Board `xml:"service"`
	Rooms     []Room       `xml:"room"`
}
