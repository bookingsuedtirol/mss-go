package response

import (
	"encoding/xml"
	"strings"
	"time"

	"github.com/hgv/mss-go/shared"
)

type Address struct {
	Street        string `xml:"street"`
	Zip           string `xml:"zip"`
	Zipcode       string `xml:"zipcode"`
	City          string `xml:"city"`
	Country       string `xml:"country"`
	UrlStreetview string `xml:"url_streetview"`
}

type Bank struct {
	Name  string `xml:"name"`
	Iban  string `xml:"iban"`
	Swift string `xml:"swift"`
}

type Booking struct {
	BookingId     int          `xml:"booking_id"`
	StornoId      int          `xml:"storno_id"`
	BookingDate   DateWithTime `xml:"booking_date"`
	Source        string       `xml:"source"`
	HotelId       int          `xml:"hotel_id"`
	Arrival       shared.Date  `xml:"arrival"`
	Departure     shared.Date  `xml:"departure"`
	Service       int          `xml:"service"`
	BookingStatus int          `xml:"booking_status"`
	Cancelled     int          `xml:"cancelled"`
	Note          string       `xml:"note"`
	Hotel         Hotel        `xml:"hotel"`
	Guest         Guest        `xml:"guest"`
	Company       Company      `xml:"company"`
	Payment       Payment      `xml:"payment"`
	Room          []Room       `xml:"room"`
	ExtraPrice    []Price      `xml:"extra_price"`
	Offer         []Offer      `xml:"offer"`
	Insurance     Insurance    `xml:"insurance"`
	Coupon        Coupon       `xml:"coupon"`
}

type CancelPolicy struct {
	Id              int          `xml:"id"`
	Refundable      bool         `xml:"refundable"`
	RefundableUntil DateWithTime `xml:"refundable_until"`
	Penalties       []Penalty    `xml:"penalties>penalty"`
	Description     Nl2brString  `xml:"description"`
	Priority        string       `xml:"priority"`
}

type Channel struct {
	ChannelId        string           `xml:"channel_id"`
	OfferId          int              `xml:"offer_id"`
	OfferDescription []Offer          `xml:"offer_description>offer"`
	RoomPrice        []RoomPrice      `xml:"room_price>price"`
	RoomDescription  []Room           `xml:"room_description>room"`
	ServicePrice     []Price          `xml:"service_price>price"`
	FromPrice        int              `xml:"from_price"`
	BasePrice        []RoomPrice      `xml:"base_price>price"`
	CancelPolicies   []CancelPolicy   `xml:"cancel_policies>cancel_policy"`
	PaymentTerms     []PaymentTerm    `xml:"payment_terms>payment_term"`
	PriceList        ChannelPriceList `xml:"pricelist"`
}

type ChannelPriceList struct {
	OfferId   int     `xml:"offer_id"`
	Inclusive []Price `xml:"inclusive>price"`
}

// TODO: Perhaps use time.Time instead of string here?
type CheckInOut struct {
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
	Email TrimmedString `xml:"email"`
	Phone string        `xml:"phone"`
	Fax   string        `xml:"fax"`
	Web   string        `xml:"web"`
}

type Coupon struct {
	CouponType   string `xml:"coupon_type"`
	CouponCode   string `xml:"coupon_code"`
	CouponStatus string `xml:"coupon_status"`
	CouponValid  int    `xml:"coupon_valid"`
	// TODO: use custom time: DateTime<'Y-m-d\TH:i:sO'>
	CouponValidFrom string `xml:"coupon_valid_from"`
	// TODO: use custom time: DateTime<'Y-m-d\TH:i:sO'>
	CouponValidTo string  `xml:"coupon_valid_to"`
	CouponTitle   string  `xml:"coupon_title"`
	CouponPercent string  `xml:"coupon_percent"`
	CouponAmount  float64 `xml:"coupon_amount"`
}

type CouponService struct {
	Provider int `xml:"provider"`
}

type Day struct {
	Date         shared.Date   `xml:"date"`
	Free         int           `xml:"free"`
	Restrictions []Restriction `xml:"restrictions>restriction"`
}

type Error struct {
	Code    int    `xml:"code"`
	Message string `xml:"message"`
}

type Feature struct {
	Id    int    `xml:"id"`
	Title string `xml:"title"`
}

type Field struct {
	Name  string `xml:"name"`
	Value string `xml:"value"`
}

type Form struct {
	FormUrl     string   `xml:"form_url"`
	FormIframe  int      `xml:"form_iframe"`
	FormMethods []string `xml:"form_methods>method"`
	FormFields  int      `xml:"form_fields"`
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
	GuestId   int           `xml:"guest_id"`
	Firstname string        `xml:"firstname"`
	Lastname  string        `xml:"lastname"`
	Prefix    string        `xml:"prefix"`
	Gender    string        `xml:"gender"`
	Email     TrimmedString `xml:"email"`
	Phone     string        `xml:"phone"`
	Address   Address       `xml:"address"`
}

type Header struct {
	Error    Error  `xml:"error"`
	ResultId string `xml:"result_id"`
	Source   string `xml:"source"`
	Paging   Paging `xml:"paging"`
	Time     string `xml:"time"`
}

type Hotel struct {
	Id                    int           `xml:"id"`
	IdLts                 string        `xml:"id_lts"`
	Bookable              bool          `xml:"bookable"`
	Name                  TrimmedString `xml:"name"`
	Type                  int           `xml:"type"`
	Stars                 float64       `xml:"stars"`
	Address               Address       `xml:"address"`
	Themes                int           `xml:"themes"`
	Features              int           `xml:"features"`
	Location              HotelLocation `xml:"location"`
	LocationName          LocationName  `xml:"location_name"`
	Geolocation           Geolocation   `xml:"geolocation"`
	Contact               Contact       `xml:"contact"`
	Headline              string        `xml:"headline"`
	Description           string        `xml:"description"`
	HotelPayment          HotelPayment  `xml:"hotel_payment"`
	Matching              Matching      `xml:"matching"`
	Logo                  []Picture     `xml:"logo>picture"`
	Pictures              []Picture     `xml:"pictures>picture"`
	AvailableFrom         shared.Date   `xml:"available_from"`
	PricesChangedAt       DateWithTime  `xml:"prices_changed_at"`
	AvailabilityChangedAt DateWithTime  `xml:"availability_changed_at"`
	// TODO: use time.Time here?
	BookableUntil string         `xml:"bookable_until"`
	Gallery       []Picture      `xml:"gallery>picture"`
	FeaturesView  []Feature      `xml:"features_view>feature"`
	Channel       Channel        `xml:"channel"`
	LtsData       shared.LtsData `xml:"lts_data"`
	Pos           []string       `xml:"pos>id_pos"`
	PriceEngine   int            `xml:"price_engine"`
	Language      string         `xml:"language"`
	CheckIn       CheckInOut     `xml:"check_in"`
	CheckOut      CheckInOut     `xml:"check_out"`
	PriceFrom     int            `xml:"price_from"`
	Board         int            `xml:"board"`
	BoardTq       int            `xml:"board_tq"`
	PersAgeMin    int            `xml:"pers_age_min"`
	ChildAgeMin   int            `xml:"child_age_min"`
	AdultAgeMin   int            `xml:"adult_age_min"`
	ChildAgeMax   int            `xml:"child_age_max"`
	AdultCntMax   int            `xml:"adult_cnt_max"`
	ChildCntMax   int            `xml:"child_cnt_max"`
	Ratings       []Rating       `xml:"ratings"`
	SourceData    SourceData     `xml:"source_data"`
	Coupon        CouponService  `xml:"coupon"`
}

type HotelLocation struct {
	IdCity      int `xml:"id_city"`
	IdCommunity int `xml:"id_community"`
	IdRegion    int `xml:"id_region"`
	IdArea      int `xml:"id_area"`
}

type HotelPayment struct {
	Methods int `xml:"methods"`
}

type Insurance struct {
	InsuranceType int    `xml:"insurance_type"`
	InsuranceUrl  string `xml:"insurance_url"`
	AgencyNr      string `xml:"agency_nr"`
}

type Location struct {
	Id        int     `xml:"id"`
	RootId    int     `xml:"root_id"`
	ParentId  int     `xml:"parent_id"`
	VirtualId string  `xml:"virtual_id"`
	Typ       string  `xml:"typ"`
	Visible   int     `xml:"visible"`
	Latitude  float64 `xml:"latitude"`
	Longitude float64 `xml:"longitude"`
	NameDeu   string  `xml:"name_deu"`
	NameIta   string  `xml:"name_ita"`
	NameSpa   string  `xml:"name_spa"`
	NameFra   string  `xml:"name_fra"`
	NameRus   string  `xml:"name_rus"`
	NameDan   string  `xml:"name_dan"`
	NameEng   string  `xml:"name_eng"`
}

type LocationName struct {
	NameCity      string `xml:"name_city"`
	NameCommunity string `xml:"name_community"`
	NameRegion    string `xml:"name_region"`
	NameArea      string `xml:"name_area"`
}

type Matching struct {
	IdBok int `xml:"id_bok"`
	IdExp int `xml:"id_exp"`
	IdHtl int `xml:"id_htl"`
	IdHrs int `xml:"id_hrs"`
}

type Occupancy struct {
	Min int `xml:"min"`
	Max int `xml:"max"`
	Std int `xml:"std"`
	Mfp int `xml:"mfp"`
}

type Offer struct {
	OfferId          int       `xml:"offer_id"`
	OfferGid         int       `xml:"offer_gid"`
	OfferBaseId      int       `xml:"offer_base_id"`
	OfferTyp         int       `xml:"offer_typ"`
	OfferShow        int       `xml:"offer_show"`
	OfferTitle       string    `xml:"offer_title"`
	Title            string    `xml:"title"`
	OfferDescription string    `xml:"offer_description"`
	Description      string    `xml:"description"`
	Pictures         []Picture `xml:"pictures>picture"`
	Themes           []Theme   `xml:"themes>theme"`
}

type Paging struct {
	Count int `xml:"count"`
	Total int `xml:"total"`
}

type Payment struct {
	Method     int     `xml:"method"`
	Price      float64 `xml:"price"`
	Prepayment float64 `xml:"prepayment"`
	Invoice    int     `xml:"invoice"`
}

type PaymentTerm struct {
	Id          int       `xml:"id"`
	OwnerId     int       `xml:"owner_id"`
	Methods     int       `xml:"methods"`
	Ccards      int       `xml:"ccards"`
	Prepayment  int       `xml:"prepayment"`
	Priority    int       `xml:"priority"`
	Bank        Bank      `xml:"bank"`
	Description string    `xml:"description"`
	Insurance   Insurance `xml:"insurance"`
}

type Penalty struct {
	Percent     int          `xml:"percent"`
	Datefrom    DateWithTime `xml:"datefrom"`
	Daysarrival int          `xml:"daysarrival"`
}

type Picture struct {
	Url       string `xml:"url"`
	Time      int    `xml:"time"`
	Title     string `xml:"title"`
	Copyright string `xml:"copyright"`
	Width     int    `xml:"width"`
	Height    int    `xml:"height"`
}

type Price struct {
	PriceId          int       `xml:"price_id"`
	PriceTyp         int       `xml:"price_typ"`
	Title            string    `xml:"title"`
	PriceTitle       string    `xml:"price_title"`
	Description      string    `xml:"description"`
	PriceDescription string    `xml:"price_description"`
	Supplement       int       `xml:"supplement"`
	PriceSupplement  int       `xml:"price_supplement"`
	PriceWs          float64   `xml:"price_ws"`
	PriceBb          float64   `xml:"price_bb"`
	PriceHb          float64   `xml:"price_hb"`
	PriceFb          float64   `xml:"price_fb"`
	PriceAi          float64   `xml:"price_ai"`
	PriceAmount      float64   `xml:"price_amount"`
	PriceValue       float64   `xml:"price_value"`
	PriceTotal       float64   `xml:"price_total"`
	Unit             int       `xml:"unit"`
	Pictures         []Picture `xml:"pictures>picture"`
}

type PriceList struct {
	OfferId        int      `xml:"offer_id"`
	OfferTyp       int      `xml:"offer_typ"`
	OfferBaseId    int      `xml:"offer_base_id"`
	SpecialTyp     int      `xml:"special_typ"`
	PrlMode        int      `xml:"prl_mode"`
	PrlUnit        int      `xml:"prl_unit"`
	DaysArrival    int      `xml:"days_arrival"`
	DaysDeparture  int      `xml:"days_departure"`
	DaysDurMin     int      `xml:"days_dur_min"`
	DaysDurMax     int      `xml:"days_dur_max"`
	DaysArrivalMin int      `xml:"days_arrival_min"`
	DaysArrivalMax int      `xml:"days_arrival_max"`
	ChildrenMin    int      `xml:"children_min"`
	ChildrenMax    int      `xml:"children_max"`
	AdultsMin      int      `xml:"adults_min"`
	AdultsMax      int      `xml:"adults_max"`
	Title          string   `xml:"title"`
	Season         []Season `xml:"season"`
}

type Properties struct {
	Area        int `xml:"area"`
	BedRooms    int `xml:"bed_rooms"`
	LivingRooms int `xml:"living_rooms"`
	DiningRooms int `xml:"dining_rooms"`
	BathRooms   int `xml:"bath_rooms"`
	WcRooms     int `xml:"wc_rooms"`
	Min         int `xml:"min"`
	Max         int `xml:"max"`
	Std         int `xml:"std"`
	Mfp         int `xml:"mfp"`
}

type Rating struct {
	Id       string      `xml:"id"`
	Provider string      `xml:"provider"`
	Value    float64     `xml:"value"`
	Count    int         `xml:"count"`
	Date     shared.Date `xml:"date"`
}

type Restriction struct {
	ObjId          int `xml:"obj_id"`
	ObjSubId       int `xml:"obj_sub_id"`
	ObjSubOnly     int `xml:"obj_sub_only"`
	Service        int `xml:"service"`
	Arrival        int `xml:"arrival"`
	Departure      int `xml:"departure"`
	Min            int `xml:"min"`
	MinArrival     int `xml:"min_arrival"`
	Max            int `xml:"max"`
	MaxArrival     int `xml:"max_arrival"`
	Close          int `xml:"close"`
	ChildrenMin    int `xml:"children_min"`
	ChildrenMax    int `xml:"children_max"`
	Holes          int `xml:"holes"`
	DaysArrivalMin int `xml:"days_arrival_min"`
	DaysArrivalMax int `xml:"days_arrival_max"`
	PersAgeMin     int `xml:"pers_age_min"`
}

type Result struct {
	Hotel    []Hotel         `xml:"hotel"`
	Special  []Special       `xml:"special"`
	Tracking Tracking        `xml:"tracking"`
	Source   []Source        `xml:"source"`
	SeoText  []SeoText       `xml:"seo_text"`
	Location []Location      `xml:"location"`
	Theme    []ThemeListItem `xml:"theme"`
}

type Room struct {
	RoomId          int          `xml:"room_id"`
	RoomLtsId       string       `xml:"room_lts_id"`
	OfferId         int          `xml:"offer_id"`
	Service         int          `xml:"service"`
	RoomType        int          `xml:"room_type"`
	RoomCode        string       `xml:"room_code"`
	RoomTitle       string       `xml:"room_title"`
	RoomDescription string       `xml:"room_description"`
	Title           string       `xml:"title"`
	Description     Nl2brString  `xml:"description"`
	RoomPersons     string       `xml:"room_persons"`
	RoomFree        int          `xml:"room_free"`
	Features        int          `xml:"features"`
	FeaturesView    []Feature    `xml:"features_view>feature"`
	RoomTotal       float64      `xml:"room_total"`
	Pictures        []Picture    `xml:"pictures>picture"`
	RoomPrice       []Price      `xml:"room_price"`
	CancelPolicy    CancelPolicy `xml:"cancel_policy"`
	PaymentTerm     PaymentTerm  `xml:"payment_term"`
	Properties      Properties   `xml:"properties"`
	Occupancy       Occupancy    `xml:"occupancy"`
	RoomNumbers     []string     `xml:"room_numbers>number"`
	RoomDetails     []RoomDetail `xml:"room_details>room_detail"`
	Days            []Day        `xml:"days>day"`
	PriceFrom       int          `xml:"price_from"`
	PriceList       []PriceList  `xml:"pricelist"`
}

type RoomDetail struct {
	Number       string `xml:"number"`
	Availability string `xml:"availability"`
}

type RoomPrice struct {
	RoomId         int     `xml:"room_id"`
	RoomSeq        int     `xml:"room_seq"`
	OfferId        int     `xml:"offer_id"`
	PriceDetails   []Price `xml:"price_details>price"`
	PriceTotal     Price   `xml:"price_total"`
	PriceInclusive Price   `xml:"price_inclusive"`
	CancelPolicyId int     `xml:"cancel_policy_id"`
	PaymentTermId  int     `xml:"payment_term_id"`
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

type SeoText struct {
	Id             int       `xml:"id"`
	Typ            string    `xml:"typ"`
	ExternalId     int       `xml:"external_id"`
	HeadlineDeu    string    `xml:"headline_deu"`
	DescriptionDeu string    `xml:"description_deu"`
	VideoDeu       string    `xml:"video_deu"`
	UrlDeu         string    `xml:"url_deu"`
	PagetitleDeu   string    `xml:"pagetitle_deu"`
	MetadescDeu    string    `xml:"metadesc_deu"`
	HeadlineIta    string    `xml:"headline_ita"`
	DescriptionIta string    `xml:"description_ita"`
	VideoIta       string    `xml:"video_ita"`
	UrlIta         string    `xml:"url_ita"`
	PagetitleIta   string    `xml:"pagetitle_ita"`
	MetadescIta    string    `xml:"metadesc_ita"`
	HeadlineEng    string    `xml:"headline_eng"`
	DescriptionEng string    `xml:"description_eng"`
	VideoEng       string    `xml:"video_eng"`
	UrlEng         string    `xml:"url_eng"`
	PagetitleEng   string    `xml:"pagetitle_eng"`
	MetadescEng    string    `xml:"metadesc_eng"`
	HeadlineSpa    string    `xml:"headline_spa"`
	DescriptionSpa string    `xml:"description_spa"`
	VideoSpa       string    `xml:"video_spa"`
	UrlSpa         string    `xml:"url_spa"`
	PagetitleSpa   string    `xml:"pagetitle_spa"`
	MetadescSpa    string    `xml:"metadesc_spa"`
	HeadlineFra    string    `xml:"headline_fra"`
	DescriptionFra string    `xml:"description_fra"`
	VideoFra       string    `xml:"video_fra"`
	UrlFra         string    `xml:"url_fra"`
	PagetitleFra   string    `xml:"pagetitle_fra"`
	MetadescFra    string    `xml:"metadesc_fra"`
	HeadlineRus    string    `xml:"headline_rus"`
	DescriptionRus string    `xml:"description_rus"`
	VideoRus       string    `xml:"video_rus"`
	UrlRus         string    `xml:"url_rus"`
	PagetitleRus   string    `xml:"pagetitle_rus"`
	MetadescRus    string    `xml:"metadesc_rus"`
	HeadlineDan    string    `xml:"headline_dan"`
	DescriptionDan string    `xml:"description_dan"`
	VideoDan       string    `xml:"video_dan"`
	UrlDan         string    `xml:"url_dan"`
	PagetitleDan   string    `xml:"pagetitle_dan"`
	MetadescDan    string    `xml:"metadesc_dan"`
	Pictures       []Picture `xml:"pictures>picture"`
}

type Source struct {
	SrcId      int    `xml:"src_id"`
	Sourcename string `xml:"sourcename"`
	Logkey     string `xml:"logkey"`
	De         string `xml:"de"`
	En         string `xml:"en"`
	It         string `xml:"it"`
}

type SourceData struct {
	Description string `xml:"description"`
	Headline    string `xml:"headline"`
	Url         string `xml:"url"`
}

type Special struct {
	OfferId        int         `xml:"offer_id"`
	Status         int         `xml:"status"`
	Valid          int         `xml:"valid"`
	OfferTyp       int         `xml:"offer_typ"`
	SpecialTyp     int         `xml:"special_typ"`
	SpecialPremium int         `xml:"special_premium"`
	DaysArrival    int         `xml:"days_arrival"`
	DaysDeparture  int         `xml:"days_departure"`
	DaysDurMin     int         `xml:"days_dur_min"`
	DaysDurMax     int         `xml:"days_dur_max"`
	DaysArrivalMin int         `xml:"days_arrival_min"`
	DaysArrivalMax int         `xml:"days_arrival_max"`
	ChildrenMin    int         `xml:"children_min"`
	ChildrenMax    int         `xml:"children_max"`
	AdultsMin      int         `xml:"adults_min"`
	AdultsMax      int         `xml:"adults_max"`
	PersAgeMin     int         `xml:"pers_age_min"`
	ChildAgeMin    int         `xml:"child_age_min"`
	ChildAgeMax    int         `xml:"child_age_max"`
	AdultAgeMin    int         `xml:"adult_age_min"`
	ValidStart     shared.Date `xml:"valid_start"`
	ValidEnd       shared.Date `xml:"valid_end"`
	Title          string      `xml:"title"`
	Description    string      `xml:"description"`
	Hotels         []Hotel     `xml:"hotels>hotel"`
	Seasons        []Season    `xml:"seasons>season"`
	Services       []int       `xml:"services>service"`
	Inclusive      []Price     `xml:"inclusive>price"`
	Pictures       []Picture   `xml:"pictures>picture"`
	Themes         []Theme     `xml:"themes>theme"`
}

type Theme struct {
	Id    int    `xml:"id"`
	Title string `xml:"title"`
}

type ThemeListItem struct {
	Id        int    `xml:"id"`
	FilterId  int    `xml:"filter_id"`
	VirtualId int    `xml:"virtual_id"`
	Sequence  int    `xml:"sequence"`
	TitleDeu  string `xml:"title_deu"`
	TitleIta  string `xml:"title_ita"`
	TitleEng  string `xml:"title_eng"`
	TitleSpa  string `xml:"title_spa"`
	TitleFra  string `xml:"title_fra"`
	TitleRus  string `xml:"title_rus"`
	TitleDan  string `xml:"title_dan"`
}

type Tracking struct {
	Pixel string `xml:"pixel"`
}

type DateWithTime time.Time

// TODO: implement string representation method like in time.Time
func (input *DateWithTime) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	layout := "2006-01-02 15:04:05"
	var value string
	err := decoder.DecodeElement(&value, &start)

	if err != nil {
		return err
	}

	// use default time if empty
	if value == "" {
		*input = DateWithTime(time.Time{})
		return nil
	}

	// TODO: parse local time to UTC?
	parsedTime, err := time.Parse(layout, value)

	if err != nil {
		return err
	}

	*input = DateWithTime(parsedTime)

	return nil
}

type TrimmedString string

func (input *TrimmedString) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var value string
	err := decoder.DecodeElement(&value, &start)

	if err != nil {
		return err
	}

	*input = TrimmedString(strings.TrimSpace(value))

	return nil
}

type Nl2brString string

func (input *Nl2brString) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var value string
	err := decoder.DecodeElement(&value, &start)

	if err != nil {
		return err
	}

	trimmedString := strings.TrimSpace(value)
	nl2brString := strings.ReplaceAll(trimmedString, "\n", "<br />\n")
	*input = Nl2brString(nl2brString)

	return nil
}
