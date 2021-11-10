package request

import "encoding/xml"

func (input Bool) MarshalXML(element *xml.Encoder, start xml.StartElement) error {
	out := 0
	if input {
		out = 1
	}
	return element.EncodeElement(out, start)
}
