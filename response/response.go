package response

import (
	"encoding/xml"
	"time"

	"github.com/HGV/mss-go/shared"
)

type Address struct {
	Street  string `xml:"street"`
	ZIP     string `xml:"zip"`
	ZIPCode string `xml:"zipcode"`
	City    string `xml:"city"`
	// The ISO 3166 alpha-3 code
	Country       string `xml:"country"`
	URLStreetview string `xml:"url_streetview"`
}

type Bank struct {
	Name  string `xml:"name"`
	IBAN  string `xml:"iban"`
	SWIFT string `xml:"swift"`
}

type CancelledStatus int

// All statuses can be cancelled with cancelBooking except for Cancelled (1)
const (
	CancelledStatusNotCancelled CancelledStatus = iota
	CancelledStatusCancelled
	CancelledStatusNoShow CancelledStatus = iota + 5
	CancelledStatusUnknown
)

type Booking struct {
	BookingID     int             `xml:"booking_id"`
	StornoID      string          `xml:"storno_id"`
	BookingDate   DateTime        `xml:"booking_date"`
	Source        string          `xml:"source"`
	HotelID       int             `xml:"hotel_id"`
	Arrival       shared.Date     `xml:"arrival"`
	Departure     shared.Date     `xml:"departure"`
	Service       shared.Board    `xml:"service"`
	BookingStatus bool            `xml:"booking_status"`
	Cancelled     CancelledStatus `xml:"cancelled"`
	Note          string          `xml:"note"`
	Hotel         Hotel           `xml:"hotel"`
	Guest         Guest           `xml:"guest"`
	Company       Company         `xml:"company"`
	Payment       Payment         `xml:"payment"`
	Rooms         []Room          `xml:"room"`
	ExtraPrices   []Price         `xml:"extra_price"`
	Offers        []Offer         `xml:"offer"`
	Insurance     Insurance       `xml:"insurance"`
	Coupon        Coupon          `xml:"coupon"`
}

type CancelPolicy struct {
	ID              int       `xml:"id"`
	Refundable      bool      `xml:"refundable"`
	RefundableUntil DateTime  `xml:"refundable_until"`
	Penalties       []Penalty `xml:"penalties>penalty"`
	// Description can contain \n characters.
	Description string `xml:"description"`
	Priority    string `xml:"priority"`
}

type Channel struct {
	ChannelID         string           `xml:"channel_id"`
	OfferID           int              `xml:"offer_id"`
	OfferDescriptions []Offer          `xml:"offer_description>offer"`
	RoomPrices        []RoomPrice      `xml:"room_price>price"`
	RoomDescriptions  []Room           `xml:"room_description>room"`
	ServicePrices     []Price          `xml:"service_price>price"`
	FromPrice         int              `xml:"from_price"`
	BasePrices        []RoomPrice      `xml:"base_price>price"`
	CancelPolicies    []CancelPolicy   `xml:"cancel_policies>cancel_policy"`
	PaymentTerms      []PaymentTerm    `xml:"payment_terms>payment_term"`
	PriceList         ChannelPriceList `xml:"pricelist"`
}

type ChannelPriceList struct {
	OfferID    int     `xml:"offer_id"`
	Inclusives []Price `xml:"inclusive>price"`
}

// The From/To fields are inserted by users as
// strings without a specific format, so they can’t be parsed as time.Time.
type CheckIn struct {
	From string `xml:"from"`
	To   string `xml:"to"`
	Note string `xml:"note"`
}

type CheckOut struct {
	From string `xml:"from"`
	To   string `xml:"to"`
}

type Company struct {
	Name          string  `xml:"name"`
	Taxnumber     string  `xml:"taxnumber"`
	RecipientCode string  `xml:"recipient_code"`
	Address       Address `xml:"address"`
}

type Contact struct {
	Email string `xml:"email"`
	Phone string `xml:"phone"`
	Fax   string `xml:"fax"`
	Web   string `xml:"web"`
}

type CouponStatus string

const (
	CouponStatusRegistered CouponStatus = "registered"
	CouponStatusRedeemable CouponStatus = "redeemable"
	CouponStatusRedeemed   CouponStatus = "redeemed"
	CouponStatusExpired    CouponStatus = "expired"
	CouponStatusCancelled  CouponStatus = "cancelled"
	CouponStatusUnknown    CouponStatus = "unknown"
)

type Coupon struct {
	CouponType      shared.CouponType    `xml:"coupon_type"`
	CouponCode      string               `xml:"coupon_code"`
	CouponStatus    CouponStatus         `xml:"coupon_status"`
	CouponValid     bool                 `xml:"coupon_valid"`
	CouponValidFrom DateTimeWithTimeZone `xml:"coupon_valid_from"`
	CouponValidTo   DateTimeWithTimeZone `xml:"coupon_valid_to"`
	CouponTitle     string               `xml:"coupon_title"`
	CouponPercent   float64              `xml:"coupon_percent"`
	CouponAmount    float64              `xml:"coupon_amount"`
}

type CouponProvider int

const (
	CouponProviderEasiCoupon CouponProvider = iota + 1
	CouponProviderGetavo
)

type CouponService struct {
	Provider CouponProvider `xml:"provider"`
}

type Day struct {
	Date         shared.Date   `xml:"date"`
	Free         int           `xml:"free"`
	Restrictions []Restriction `xml:"restrictions>restriction"`
}

type ErrorCode int

const (
	ErrorCodeGenericError ErrorCode = 1 << iota
	ErrorCodeAuthenticationError
	ErrorCodeInvalidXML
	ErrorCodeInvalidMethod
	ErrorCodeResultIDNotInCache
	ErrorCodeInvalidMissingParameter
	ErrorCodeBookingValidationFailed
	ErrorCodePermissionsDenied
)

type Error struct {
	Code    ErrorCode `xml:"code"`
	Message string    `xml:"message"`
}

type Feature struct {
	ID    int    `xml:"id"`
	Title string `xml:"title"`
}

type Field struct {
	Name  string `xml:"name"`
	Value string `xml:"value"`
}

type FormIframe int

const (
	FormIframeNotAllowed FormIframe = iota
	FormIframeAllowed
	FormIframeAllowedIfHTTPS
)

type FormMethod string

const (
	FormMethodPOST FormMethod = "POST"
	FormMethodGET  FormMethod = "GET"
)

type FormField struct {
	Name  string `xml:"name"`
	Value string `xml:"value"`
}

type Form struct {
	FormURL     string       `xml:"form_url"`
	FormIframe  FormIframe   `xml:"form_iframe"`
	FormMethods []FormMethod `xml:"form_methods>method"`
	FormFields  []FormField  `xml:"form_fields>field"`
}

type Geolocation struct {
	Latitude  float64 `xml:"latitude"`
	Longitude float64 `xml:"longitude"`

	// TODO: use int if value float64 from accommodation 10148
	// returns an int by MSS
	Altitude float64 `xml:"altitude"`

	Distance float64 `xml:"distance"`
}

type Guest struct {
	GuestID   int           `xml:"guest_id"`
	FirstName string        `xml:"firstname"`
	LastName  string        `xml:"lastname"`
	Prefix    string        `xml:"prefix"`
	Gender    shared.Gender `xml:"gender"`
	Email     string        `xml:"email"`
	Phone     string        `xml:"phone"`
	Address   Address       `xml:"address"`
}

type Header struct {
	Error     Error     `xml:"error"`
	ResultID  string    `xml:"result_id"`
	Source    string    `xml:"source"`
	Paging    Paging    `xml:"paging"`
	RateLimit RateLimit `xml:"rate_limit"`
	Time      string    `xml:"time"`
}

type PriceEngine int

const (
	PriceEngineEasiSuite PriceEngine = iota
	PriceEngineLTS
	PriceEngineDisabled
)

type BoardBit int

const (
	BoardBitUndefined    BoardBit = 0
	BoardBitWithoutBoard BoardBit = 1 << (iota - 1)
	BoardBitWithBreakfast
	BoardBitHalfBoard
	BoardBitFullBoard
	BoardBitAllInclusive
)

type Hotel struct {
	ID                    int                  `xml:"id"`
	IDLTS                 string               `xml:"id_lts"`
	Bookable              bool                 `xml:"bookable"`
	Name                  string               `xml:"name"`
	Type                  shared.HotelType     `xml:"type"`
	Stars                 float64              `xml:"stars"`
	Address               Address              `xml:"address"`
	Themes                shared.Theme         `xml:"themes"`
	Features              shared.HotelFeature  `xml:"features"`
	Location              HotelLocation        `xml:"location"`
	LocationName          LocationName         `xml:"location_name"`
	Geolocation           Geolocation          `xml:"geolocation"`
	Contact               Contact              `xml:"contact"`
	Headline              string               `xml:"headline"`
	Description           NormalizedHTMLString `xml:"description"`
	HotelPayment          HotelPayment         `xml:"hotel_payment"`
	Matching              Matching             `xml:"matching"`
	Logo                  Picture              `xml:"logo>picture"`
	Pictures              []Picture            `xml:"pictures>picture"`
	AvailableFrom         shared.Date          `xml:"available_from"`
	PricesChangedAt       DateTime             `xml:"prices_changed_at"`
	AvailabilityChangedAt DateTime             `xml:"availability_changed_at"`
	BookableUntil         Time                 `xml:"bookable_until"`
	Gallery               []Picture            `xml:"gallery>picture"`
	FeaturesView          []Feature            `xml:"features_view>feature"`
	Channel               Channel              `xml:"channel"`
	LTSData               shared.LTSData       `xml:"lts_data"`
	POS                   []string             `xml:"pos>id_pos"`
	PriceEngine           PriceEngine          `xml:"price_engine"`
	Language              string               `xml:"language"`
	CheckIn               CheckIn              `xml:"check_in"`
	CheckOut              CheckOut             `xml:"check_out"`
	PriceFrom             int                  `xml:"price_from"`
	Board                 BoardBit             `xml:"board"`
	BoardThreeQuarters    bool                 `xml:"board_tq"`
	PersAgeMin            int                  `xml:"pers_age_min"`
	ChildAgeMin           int                  `xml:"child_age_min"`
	AdultAgeMin           int                  `xml:"adult_age_min"`
	ChildAgeMax           int                  `xml:"child_age_max"`
	AdultCntMax           int                  `xml:"adult_cnt_max"`
	ChildCntMax           int                  `xml:"child_cnt_max"`
	Ratings               []Rating             `xml:"ratings>rating"`
	SourceData            SourceData           `xml:"source_data"`
	Coupon                CouponService        `xml:"coupon"`
	RoomTypes             shared.RoomType      `xml:"room_types"`
	InformalMail          bool                 `xml:"informal_mail"`
}

type HotelLocation struct {
	IDCity      int `xml:"id_city"`
	IDCommunity int `xml:"id_community"`
	IDRegion    int `xml:"id_region"`
	IDArea      int `xml:"id_area"`
}

type PaymentMethods int

const (
	PaymentMethodsCreditCard PaymentMethods = 8 << iota
	PaymentMethodsATM        PaymentMethods = 32 << iota
	PaymentMethodsMastercard
	PaymentMethodsVisa
	PaymentMethodsDinersClub
	PaymentMethodsAmericanExpress
)

type HotelPayment struct {
	Methods PaymentMethods `xml:"methods"`
}

type InsuranceType int

const (
	InsuranceTypeHGV InsuranceType = iota + 1
	InsuranceTypeHogast
)

type Insurance struct {
	InsuranceType InsuranceType `xml:"insurance_type"`
	InsuranceURL  string        `xml:"insurance_url"`
	AgencyNr      string        `xml:"agency_nr"`
}

type Visibility int

// The difference between VisibilityVisible1 and VisibilityVisible3 is unclear.
const (
	VisibilityHidden   Visibility = 0
	VisibilityVisible1 Visibility = 1
	VisibilityVisible3 Visibility = 3
)

type Location struct {
	ID         int                 `xml:"id"`
	RootID     int                 `xml:"root_id"`
	ParentID   int                 `xml:"parent_id"`
	VirtualIDs Ints                `xml:"virtual_id"`
	Type       shared.LocationType `xml:"typ"`
	Visible    Visibility          `xml:"visible"`
	Latitude   float64             `xml:"latitude"`
	Longitude  float64             `xml:"longitude"`
	NameDeu    string              `xml:"name_deu"`
	NameIta    string              `xml:"name_ita"`
	NameSpa    string              `xml:"name_spa"`
	NameFra    string              `xml:"name_fra"`
	NameRus    string              `xml:"name_rus"`
	NameDan    string              `xml:"name_dan"`
	NameEng    string              `xml:"name_eng"`
}

type LocationName struct {
	NameCity      string `xml:"name_city"`
	NameCommunity string `xml:"name_community"`
	NameRegion    string `xml:"name_region"`
	NameArea      string `xml:"name_area"`
}

type Matching struct {
	IDBok int `xml:"id_bok"`
	IDExp int `xml:"id_exp"`
	IDHtl int `xml:"id_htl"`
	IDHrs int `xml:"id_hrs"`
}

type Occupancy struct {
	Min int `xml:"min"`
	Max int `xml:"max"`
	Std int `xml:"std"`
	Mfp int `xml:"mfp"`
}

type Offer struct {
	OfferID          int                  `xml:"offer_id"`
	OfferGID         int                  `xml:"offer_gid"`
	OfferBaseID      int                  `xml:"offer_base_id"`
	OfferType        shared.OfferType     `xml:"offer_typ"`
	OfferTitle       string               `xml:"offer_title"`
	Title            string               `xml:"title"`
	OfferDescription string               `xml:"offer_description"`
	Description      NormalizedHTMLString `xml:"description"`
	Pictures         []Picture            `xml:"pictures>picture"`
	Themes           []Theme              `xml:"themes>theme"`
}

type Paging struct {
	Count int `xml:"count"`
	Total int `xml:"total"`
}

type RateLimit struct {
	Limit     LimitPerSeconds `xml:"limit"`
	Remaining int             `xml:"remaining"`
	Reset     Duration        `xml:"reset"`
}

type LimitPerSeconds struct {
	Requests int
	Duration time.Duration
}

type Duration struct {
	time.Duration
}

type Payment struct {
	Method     shared.PaymentMethod `xml:"method"`
	Price      float64              `xml:"price"`
	Prepayment float64              `xml:"prepayment"`
	Invoice    bool                 `xml:"invoice"`
	Bank       Bank                 `xml:"bank"`
}

type PaymentTerm struct {
	ID          int                  `xml:"id"`
	OwnerID     int                  `xml:"owner_id"`
	Methods     shared.PaymentMethod `xml:"methods"`
	CCards      int                  `xml:"ccards"`
	Prepayment  int                  `xml:"prepayment"`
	Priority    int                  `xml:"priority"`
	Bank        Bank                 `xml:"bank"`
	Description string               `xml:"description"`
	Insurance   Insurance            `xml:"insurance"`
}

type Penalty struct {
	Percent     int      `xml:"percent"`
	Datefrom    DateTime `xml:"datefrom"`
	Daysarrival int      `xml:"daysarrival"`
}

type Picture struct {
	URL       string `xml:"url"`
	Time      int    `xml:"time"`
	Title     string `xml:"title"`
	Copyright string `xml:"copyright"`
	Width     int    `xml:"width"`
	Height    int    `xml:"height"`
}

type Supplement int

const (
	SupplementRoomPrice Supplement = iota
	SupplementSurchargesOrDiscounts
	SupplementIncludedServices
)

type PriceUnit int

const (
	PriceUnitEuro PriceUnit = iota
	PriceUnitPercent
)

type Price struct {
	PriceID          int                  `xml:"price_id"`
	PriceType        shared.OfferType     `xml:"price_typ"`
	Title            string               `xml:"title"`
	PriceTitle       string               `xml:"price_title"`
	Description      NormalizedHTMLString `xml:"description"`
	PriceDescription string               `xml:"price_description"`
	Supplement       Supplement           `xml:"supplement"`
	PriceSupplement  int                  `xml:"price_supplement"`
	PriceWS          float64              `xml:"price_ws"`
	PriceBB          float64              `xml:"price_bb"`
	PriceHB          float64              `xml:"price_hb"`
	PriceFB          float64              `xml:"price_fb"`
	PriceAI          float64              `xml:"price_ai"`
	PriceAmount      float64              `xml:"price_amount"`
	PriceValue       float64              `xml:"price_value"`
	PriceTotal       float64              `xml:"price_total"`
	Unit             PriceUnit            `xml:"unit"`
	Pictures         []Picture            `xml:"pictures>picture"`
}

type PriceList struct {
	OfferID        int                `xml:"offer_id"`
	OfferType      shared.OfferType   `xml:"offer_typ"`
	OfferBaseID    int                `xml:"offer_base_id"`
	SpecialType    shared.SpecialType `xml:"special_typ"`
	PrlMode        int                `xml:"prl_mode"`
	PrlUnit        int                `xml:"prl_unit"`
	DaysArrival    Weekdays           `xml:"days_arrival"`
	DaysDeparture  Weekdays           `xml:"days_departure"`
	DaysDurMin     int                `xml:"days_dur_min"`
	DaysDurMax     int                `xml:"days_dur_max"`
	DaysArrivalMin int                `xml:"days_arrival_min"`
	DaysArrivalMax int                `xml:"days_arrival_max"`
	ChildrenMin    int                `xml:"children_min"`
	ChildrenMax    int                `xml:"children_max"`
	AdultsMin      int                `xml:"adults_min"`
	AdultsMax      int                `xml:"adults_max"`
	Title          string             `xml:"title"`
	Seasons        []Season           `xml:"season"`
}

type Properties struct {
	Area        int `xml:"area"`
	BedRooms    int `xml:"bed_rooms"`
	LivingRooms int `xml:"living_rooms"`
	DiningRooms int `xml:"dining_rooms"`
	BathRooms   int `xml:"bath_rooms"`
	WCRooms     int `xml:"wc_rooms"`
	Min         int `xml:"min"`
	Max         int `xml:"max"`
	Std         int `xml:"std"`
	Mfp         int `xml:"mfp"`
}

type Rating struct {
	ID       string      `xml:"id"`
	Provider string      `xml:"provider"`
	Value    float64     `xml:"value"`
	Count    int         `xml:"count"`
	Date     shared.Date `xml:"date"`
}

type Restriction struct {
	ObjID          int          `xml:"obj_id"`
	ObjSubID       int          `xml:"obj_sub_id"`
	ObjSubOnly     int          `xml:"obj_sub_only"`
	Service        shared.Board `xml:"service"`
	Arrival        bool         `xml:"arrival"`
	Departure      bool         `xml:"departure"`
	Min            int          `xml:"min"`
	MinArrival     int          `xml:"min_arrival"`
	Max            int          `xml:"max"`
	MaxArrival     int          `xml:"max_arrival"`
	Close          bool         `xml:"close"`
	ChildrenMin    int          `xml:"children_min"`
	ChildrenMax    int          `xml:"children_max"`
	Holes          bool         `xml:"holes"`
	DaysArrivalMin int          `xml:"days_arrival_min"`
	DaysArrivalMax int          `xml:"days_arrival_max"`
	PersAgeMin     int          `xml:"pers_age_min"`
}

type Result struct {
	Hotels    []Hotel         `xml:"hotel"`
	Specials  []Special       `xml:"special"`
	Tracking  Tracking        `xml:"tracking"`
	Sources   []Source        `xml:"source"`
	Locations []Location      `xml:"location"`
	Themes    []ThemeListItem `xml:"theme"`
	Booking   Booking         `xml:"booking"`
	Form      Form            `xml:"form"`
	Coupon    Coupon          `xml:"coupon"`
}

type Room struct {
	RoomID          int             `xml:"room_id"`
	RoomLTSID       string          `xml:"room_lts_id"`
	OfferID         int             `xml:"offer_id"`
	Service         shared.Board    `xml:"service"`
	RoomType        shared.RoomType `xml:"room_type"`
	RoomCode        string          `xml:"room_code"`
	RoomTitle       string          `xml:"room_title"`
	RoomDescription string          `xml:"room_description"`
	Title           string          `xml:"title"`
	// Description can contain \n characters.
	Description  string       `xml:"description"`
	RoomPersons  Ints         `xml:"room_persons"`
	RoomFree     int          `xml:"room_free"`
	Features     int          `xml:"features"`
	FeaturesView []Feature    `xml:"features_view>feature"`
	RoomTotal    float64      `xml:"room_total"`
	Pictures     []Picture    `xml:"pictures>picture"`
	RoomPrice    []Price      `xml:"room_price"`
	CancelPolicy CancelPolicy `xml:"cancel_policy"`
	PaymentTerm  PaymentTerm  `xml:"payment_term"`
	Properties   Properties   `xml:"properties"`
	Occupancy    Occupancy    `xml:"occupancy"`
	RoomNumbers  []string     `xml:"room_numbers>number"`
	RoomDetails  []RoomDetail `xml:"room_details>room_detail"`
	Days         []Day        `xml:"days>day"`
	PriceFrom    int          `xml:"price_from"`
	PriceList    []PriceList  `xml:"pricelist"`
}

type RoomDetail struct {
	Number       string         `xml:"number"`
	Availability Availabilities `xml:"availability"`
}

type Availabilities []bool

type RoomPrice struct {
	RoomID         int     `xml:"room_id"`
	RoomSeq        int     `xml:"room_seq"`
	OfferID        int     `xml:"offer_id"`
	PriceDetails   []Price `xml:"price_details>price"`
	PriceTotal     Price   `xml:"price_total"`
	PriceInclusive Price   `xml:"price_inclusive"`
	CancelPolicyID int     `xml:"cancel_policy_id"`
	PaymentTermID  int     `xml:"payment_term_id"`
}

type Root struct {
	XMLName xml.Name `xml:"root"`
	Header  Header   `xml:"header"`
	Result  Result   `xml:"result"`
}

type Season struct {
	DateStart shared.Date `xml:"date_start"`
	DateEnd   shared.Date `xml:"date_end"`
	Price     Price       `xml:"price"`
}

type Source struct {
	SrcID      int    `xml:"src_id"`
	Sourcename string `xml:"sourcename"`
	Logkey     string `xml:"logkey"`
	De         string `xml:"de"`
	En         string `xml:"en"`
	It         string `xml:"it"`
}

type SourceData struct {
	Description string `xml:"description"`
	Headline    string `xml:"headline"`
	URL         string `xml:"url"`
}

type Weekdays int

const (
	WeekdayMonday Weekdays = 1 << iota
	WeekdayTuesday
	WeekdayWednesday
	WeekdayThursday
	WeekdayFriday
	WeekdaySaturday
	WeekdaySunday
)

type Special struct {
	OfferID        int                   `xml:"offer_id"`
	Status         int                   `xml:"status"`
	Valid          bool                  `xml:"valid"`
	OfferType      shared.OfferType      `xml:"offer_typ"`
	SpecialType    shared.SpecialType    `xml:"special_typ"`
	SpecialPremium shared.SpecialPremium `xml:"special_premium"`
	DaysArrival    Weekdays              `xml:"days_arrival"`
	DaysDeparture  Weekdays              `xml:"days_departure"`
	DaysDurMin     int                   `xml:"days_dur_min"`
	DaysDurMax     int                   `xml:"days_dur_max"`
	DaysArrivalMin int                   `xml:"days_arrival_min"`
	DaysArrivalMax int                   `xml:"days_arrival_max"`
	ChildrenMin    int                   `xml:"children_min"`
	ChildrenMax    int                   `xml:"children_max"`
	AdultsMin      int                   `xml:"adults_min"`
	AdultsMax      int                   `xml:"adults_max"`
	PersAgeMin     int                   `xml:"pers_age_min"`
	ChildAgeMin    int                   `xml:"child_age_min"`
	ChildAgeMax    int                   `xml:"child_age_max"`
	AdultAgeMin    int                   `xml:"adult_age_min"`
	ValidStart     shared.Date           `xml:"valid_start"`
	ValidEnd       shared.Date           `xml:"valid_end"`
	Title          string                `xml:"title"`
	Description    NormalizedHTMLString  `xml:"description"`
	Hotels         []Hotel               `xml:"hotels>hotel"`
	Seasons        []Season              `xml:"seasons>season"`
	Services       []shared.Board        `xml:"services>service"`
	Inclusives     []Price               `xml:"inclusive>price"`
	Pictures       []Picture             `xml:"pictures>picture"`
	Themes         []Theme               `xml:"themes>theme"`
}

type Theme struct {
	ID    shared.ThemeID `xml:"id"`
	Title string         `xml:"title"`
}

type ThemeListItem struct {
	ID         int    `xml:"id"`
	FilterID   int    `xml:"filter_id"`
	VirtualIDs Ints   `xml:"virtual_id"`
	Sequence   int    `xml:"sequence"`
	TitleDeu   string `xml:"title_deu"`
	TitleIta   string `xml:"title_ita"`
	TitleEng   string `xml:"title_eng"`
	TitleSpa   string `xml:"title_spa"`
	TitleFra   string `xml:"title_fra"`
	TitleRus   string `xml:"title_rus"`
	TitleDan   string `xml:"title_dan"`
}

type Tracking struct {
	Pixel string `xml:"pixel"`
}

type DateTime struct{ time.Time }

type DateTimeWithTimeZone struct{ time.Time }

type Time struct {
	time.Time
	// Valid is true if Time (which can also be Time.isZero() == 0) is present
	Valid bool
}

type NormalizedHTMLString string

type Ints []int
