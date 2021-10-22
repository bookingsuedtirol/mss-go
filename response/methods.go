package response

import (
	"encoding/xml"
	"strings"
	"time"

	"github.com/HGV/mss-go/shared"
)

func (input *DateTime) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	layout := "2006-01-02 15:04:05"
	value, err := shared.ParseDateTime(layout, decoder, start)

	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	*input = DateTime(value)

	return nil
}

func (dateTime DateTime) String() string {
	return time.Time(dateTime).String()
}

func (input *TrimmedString) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var value string
	err := decoder.DecodeElement(&value, &start)

	if err != nil {
		return err
	}

	*input = TrimmedString(strings.TrimSpace(value))

	return nil
}

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
