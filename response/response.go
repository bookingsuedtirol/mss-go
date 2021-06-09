package response

type Root struct {
	Header Header `xml:"header,omitempty"`
	Result Result `xml:"result,omitempty"`
}

type Header struct {
	Error    Error  `xml:"error,omitempty"`
	ResultId string `xml:"result_id,omitempty"`
	Source   string `xml:"source,omitempty"`
	Time     string `xml:"time,omitempty"`
}

type Error struct {
	Code    int    `xml:"code,omitempty"`
	Message string `xml:"message,omitempty"`
}

type Result struct {
	Hotel    []Hotel    `xml:"hotel,omitempty"`
	Location []Location `xml:"location,omitempty"`
}

type Location struct {
	Id        int     `xml:"id,omitempty"`
	RootId    int     `xml:"root_id,omitempty"`
	ParentId  int     `xml:"parent_id,omitempty"`
	VirtualId string  `xml:"virtual_id,omitempty"`
	Typ       string  `xml:"typ,omitempty"`
	Visible   int     `xml:"visible,omitempty"`
	Latitude  float64 `xml:"latitude,omitempty"`
	Longitude float64 `xml:"longitude,omitempty"`
	NameDeu   string  `xml:"name_deu,omitempty"`
	NameIta   string  `xml:"name_ita,omitempty"`
	NameEng   string  `xml:"name_eng,omitempty"`
}

type Hotel struct {
	Id       int    `xml:"id,omitempty"`
	IdLts    string `xml:"id_lts,omitempty"`
	Bookable bool   `xml:"bookable,omitempty"`
	Headline string `xml:"headline,omitempty"`
}
