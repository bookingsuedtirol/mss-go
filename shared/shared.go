package shared

import (
	"encoding/xml"
	"time"
)

type Date time.Time

func (input Date) MarshalXML(element *xml.Encoder, start xml.StartElement) error {
	timeValue := time.Time(input)
	timeString := timeValue.Format("2006-01-02")
	return element.EncodeElement(timeString, start)
}

type LtsData struct {
	A0Ene int `xml:"A0Ene"`
	A0MTV int `xml:"A0MTV"`
	A0Rep int `xml:"A0Rep"`
}
