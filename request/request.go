package request

import (
	"encoding/xml"

	"github.com/HGV/mss-go/shared"
)

type Address struct {
	Street  string `xml:"street"`
	Zipcode string `xml:"zipcode"`
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

type Data struct {
	Guest     *Guest    `xml:"guest"`
	Company   *Company  `xml:"company"`
	Payment   *Payment  `xml:"payment"`
	Note      string    `xml:"note"`
	Details   *Details  `xml:"details"`
	Form      *Form     `xml:"form"`
	Tracking  *Tracking `xml:"tracking"`
	Insurance int       `xml:"insurance"`
}

type Details struct {
	ExtraPrices []ExtraPrice `xml:"extra_price"`
	Coupon      *Coupon      `xml:"coupon"`
}

type ExtraPrice struct {
	PriceId     int     `xml:"price_id"`
	PriceAmount float64 `xml:"price_amount"`
}

type Form struct {
	UrlSuccess string `xml:"url_success"`
	UrlFailure string `xml:"url_failure"`
}

type Guest struct {
	Gender    string   `xml:"gender"`
	Prefix    string   `xml:"prefix"`
	Firstname string   `xml:"firstname"`
	Lastname  string   `xml:"lastname"`
	Email     string   `xml:"email"`
	Phone     string   `xml:"phone"`
	Address   *Address `xml:"address"`
	// TODO: map to bool?
	Newsletter int `xml:"newsletter"`
}

type Method string

type Header struct {
	Credentials Credentials `xml:"credentials"`
	Method      Method      `xml:"method"`
	Paging      *Paging     `xml:"paging"`
	ResultId    string      `xml:"result_id"`
}

type Logging struct {
	Step string `xml:"step"`
}

type HotelDetails int

type OfferDetails int

type Options struct {
	HotelDetails         HotelDetails `xml:"hotel_details"`
	OfferDetails         OfferDetails `xml:"offer_details"`
	RoomDetails          int          `xml:"room_details"`
	SpecialDetails       int          `xml:"special_details"`
	SeoDetails           int          `xml:"seo_details"`
	PictureDate          *shared.Date `xml:"picture_date"`
	LtsBookable          int          `xml:"lts_bookable"`
	GetAvailability      int          `xml:"get_availability"`
	GetRestrictions      int          `xml:"get_restrictions"`
	GetRoomdetails       int          `xml:"get_roomdetails"`
	GetMasterpackages    int          `xml:"get_masterpackages"`
	BasePrice            int          `xml:"base_price"`
	PricelistDetails     int          `xml:"pricelist_details"`
	OnlySubscribedHotels int          `xml:"only_subscribed_hotels"`
	OnlyAvailable        int          `xml:"only_available"`
}

type Order struct {
	Field string `xml:"field"`
	Dir   string `xml:"dir"`
}

type Paging struct {
	Start int `xml:"start"`
	Limit int `xml:"limit"`
}

type Payment struct {
	Method  int `xml:"method"`
	Invoice int `xml:"invoice"`
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
	OfferId  int             `xml:"offer_id"`
	RoomId   int             `xml:"room_id"`
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
	ResultId           string              `xml:"result_id"`
	Agent              string              `xml:"agent"`
	Ids                []int               `xml:"id"`
	SearchHotel        *SearchHotel        `xml:"search_hotel"`
	SearchLocation     *SearchLocation     `xml:"search_location"`
	SearchDistance     *SearchDistance     `xml:"search_distance"`
	SearchOffer        *SearchOffer        `xml:"search_offer"`
	SearchLts          *shared.LtsData     `xml:"search_lts"`
	SearchSpecial      *SearchSpecial      `xml:"search_special"`
	SearchAvailability *SearchAvailability `xml:"search_availability"`
	SearchPricelist    *SearchPriceList    `xml:"search_pricelist"`
	In                 []int               `xml:"in"`
	IdOfchannel        string              `xml:"id_ofchannel,omitempty"`
	TransactionId      string              `xml:"transaction_id"`
	BookingId          string              `xml:"booking_id"`
	GuestEmail         string              `xml:"guest_email"`
	RootId             int                 `xml:"root_id"`
	ExternalId         int                 `xml:"external_id"`
	Type               string              `xml:"typ"`
	SeoType            string              `xml:"seo_typ"`
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
	OfferIds []int        `xml:"offer_id"`
	RoomIds  []int        `xml:"room_id"`
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
	LocationsLts []string `xml:"location_lts"`
}

type SearchOffer struct {
	Arrival    *shared.Date     `xml:"arrival"`
	Departure  *shared.Date     `xml:"departure"`
	Service    shared.Board     `xml:"service"`
	Feature    int              `xml:"feature,omitempty"`
	ChannelIds []string         `xml:"channel_id"`
	Rooms      []Room           `xml:"room"`
	Type       shared.OfferType `xml:"typ"`
	Rateplan   *Rateplan        `xml:"rateplan"`
}

type SearchPriceList struct {
	DateFrom *shared.Date `xml:"date_from"`
	DateTo   *shared.Date `xml:"date_to"`
	Service  shared.Board `xml:"service"`
	RoomIds  []int        `xml:"room_id"`
	Type     int          `xml:"typ"`
}

type SearchSpecial struct {
	OfferIds []int        `xml:"offer_id"`
	DateFrom *shared.Date `xml:"date_from"`
	DateTo   *shared.Date `xml:"date_to"`
	Themes   []int        `xml:"theme"`
	PoiIds   []int        `xml:"poi_id"`
	PoiCats  []int        `xml:"poi_cat"`
	Validity *Validity    `xml:"validity"`
	Type     int          `xml:"typ"`
	Premium  int          `xml:"premium"`
	Status   int          `xml:"status"`
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
	Valid     int          `xml:"valid"`
	Offers    int          `xml:"offers"`
	Arrival   *shared.Date `xml:"arrival"`
	Departure *shared.Date `xml:"departure"`
	Service   shared.Board `xml:"service"`
	Rooms     []Room       `xml:"room"`
}
