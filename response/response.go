package response

import "encoding/xml"

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
	Id          int         `xml:"id"`
	Name        string      `xml:"name"`
	Stars       float64     `xml:"stars"`
	Geolocation Geolocation `xml:"geolocation"`
	IdLts       string      `xml:"id_lts"`
	Bookable    bool        `xml:"bookable"`
	Headline    string      `xml:"headline"`
}

type Geolocation struct {
	Latitude  float64 `xml:"latitude"`
	Longitude float64 `xml:"longitude"`
	Altitude  int     `xml:"altitude"`
	Distance  float64 `xml:"distance"`
}
