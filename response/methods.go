package response

import (
	"bytes"
	"encoding/xml"
	"io"
	"strconv"
	"strings"

	"github.com/HGV/mss-go/shared"
	"github.com/microcosm-cc/bluemonday"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func (input *DateTime) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var str string
	err := decoder.DecodeElement(&str, &start)
	if err != nil {
		return err
	}

	value, err := shared.ParseDateTime("2006-01-02 15:04:05", str)
	if err != nil {
		return err
	}

	*input = DateTime(*value)
	return nil
}

func (input *Time) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var str string
	err := decoder.DecodeElement(&str, &start)
	if err != nil {
		return err
	}

	if str == "" {
		input.Valid = false
		return nil
	}

	value, err := shared.ParseDateTime("15:04", str)
	if err != nil {
		return err
	}

	*input = Time{
		Time:  *value,
		Valid: true,
	}
	return nil
}

func (input *Nl2brString) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var value string
	err := decoder.DecodeElement(&value, &start)

	if err != nil {
		return err
	}

	nl2brString := strings.ReplaceAll(value, "\n", "<br />\n")
	*input = Nl2brString(nl2brString)

	return nil
}

func (input *NormalizedHTMLString) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var value string
	err := decoder.DecodeElement(&value, &start)

	if err != nil {
		return err
	}

	if value == "" {
		*input = ""
		return nil
	}

	fixed, err := fixHTML(strings.NewReader(value))

	if err != nil {
		return err
	}

	sanitized := sanitizePolicy.SanitizeReader(fixed)

	*input = NormalizedHTMLString(sanitized.String())

	return nil
}

// Decode from a comma-separated list of ints
func (input *Ints) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var value string
	err := decoder.DecodeElement(&value, &start)

	if err != nil {
		return err
	}

	// Don’t assign anything (is a nil-slice by default).
	if value == "" {
		return nil
	}

	idStrings := strings.Split(value, ",")
	out := make([]int, 0, len(idStrings))

	for _, idStr := range idStrings {
		id, err := strconv.Atoi(idStr)

		if err != nil {
			return err
		}

		out = append(out, id)
	}

	*input = out

	return nil
}

var sanitizePolicy = getSanitizePolicy()

func getSanitizePolicy() *bluemonday.Policy {
	p := bluemonday.StrictPolicy()
	// TODO: Check if this is the right list of allowed elements
	p.AllowElements("p", "ul", "ol", "li", "b", "strong", "br", "em", "u")
	return p
}

// Fix ill-formed HTML (which MSS sometimes outputs, e.g.
// missing closing elements) and output well-formed HTML.
func fixHTML(input io.Reader) (io.Reader, error) {
	nodes, err := html.ParseFragment(input, &html.Node{
		Type:     html.ElementNode,
		Data:     "body",
		DataAtom: atom.Body,
	})

	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)

	for _, node := range nodes {
		err = html.Render(buf, node)
		if err != nil {
			return nil, err
		}
	}

	return buf, nil
}
