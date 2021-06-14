package response

import (
	"encoding/xml"
	"strings"
	"time"
)

type Root struct {
	XMLName xml.Name `xml:"root"`
	Header  Header   `xml:"header"`
	Result  Result   `xml:"result"`
}

type Header struct {
	Error    Error  `xml:"error"`
	ResultId string `xml:"result_id"`
	Source   string `xml:"source"`
	Time     string `xml:"time"`
}

type Error struct {
	Code    int    `xml:"code"`
	Message string `xml:"message"`
}

type Result struct {
	Hotel    []Hotel    `xml:"hotel"`
	Location []Location `xml:"location"`
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
	NameEng   string  `xml:"name_eng"`
}

type Hotel struct {
	Id          int           `xml:"id"`
	Name        TrimmedString `xml:"name"`
	Stars       float64       `xml:"stars"`
	Geolocation Geolocation   `xml:"geolocation"`
	IdLts       string        `xml:"id_lts"`
	Bookable    bool          `xml:"bookable"`
	Headline    string        `xml:"headline"`
	Channel     Channel       `xml:"channel"`
}

type Geolocation struct {
	Latitude  float64 `xml:"latitude"`
	Longitude float64 `xml:"longitude"`
	Altitude  int     `xml:"altitude"`
	Distance  float64 `xml:"distance"`
}

type Channel struct {
	ChannelId        string         `xml:"channel_id"`
	OfferId          int            `xml:"offer_id"`
	RoomPrice        []RoomPrice    `xml:"room_price>price"`
	RoomDescription  []Room         `xml:"room_description>room"`
	OfferDescription []Offer        `xml:"offer_description>offer"`
	CancelPolicies   []CancelPolicy `xml:"cancel_policies>cancel_policy"`
}

type RoomPrice struct {
	RoomId     int   `xml:"room_id"`
	RoomSeq    int   `xml:"room_seq"`
	OfferId    int   `xml:"offer_id"`
	PriceTotal Price `xml:"price_total"`
}

type Price struct {
	PriceType int     `xml:"price_typ"`
	PriceWs   float64 `xml:"price_ws"`
	PriceBb   float64 `xml:"price_bb"`
	PriceHb   float64 `xml:"price_hb"`
	PriceFb   float64 `xml:"price_fb"`
	PriceAi   float64 `xml:"price_ai"`
}

type Room struct {
	RoomId          int         `xml:"room_id"`
	RoomType        int         `xml:"room_type"`
	Title           string      `xml:"title"`
	Description     Nl2brString `xml:"description"`
	RoomDescription string      `xml:"room_description"`
}

type Offer struct {
	OfferId int    `xml:"offer_id"`
	Title   string `xml:"title"`
}

type CancelPolicy struct {
	Id          int         `xml:"id"`
	Description Nl2brString `xml:"description"`
	Penalties   []Penalty   `xml:"penalties>penalty"`
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

type Penalty struct {
	Percent     int          `xml:"percent"`
	Datefrom    DateWithTime `xml:"datefrom"`
	Daysarrival int          `xml:"daysarrival"`
}
