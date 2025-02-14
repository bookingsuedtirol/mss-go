package response

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/bookingsuedtirol/mss-go/shared"
	"github.com/microcosm-cc/bluemonday"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func (t *DateTime) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var str string
	err := decoder.DecodeElement(&str, &start)
	if err != nil {
		return err
	}

	value, err := shared.ParseLocalDateTime(time.DateTime, str)
	if err != nil {
		return err
	}

	if value != nil {
		*t = DateTime{*value}
	}

	return nil
}

func (t *DateTimeWithTimeZone) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var str string
	err := decoder.DecodeElement(&str, &start)
	if err != nil {
		return err
	}

	value, err := shared.ParseDateTime("2006-01-02T15:04:05-07:00", str)
	if err != nil {
		return err
	}

	if value != nil {
		*t = DateTimeWithTimeZone{*value}
	}

	return nil
}

func (t *Time) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var str string
	err := decoder.DecodeElement(&str, &start)
	if err != nil {
		return err
	}

	if str == "" {
		return nil
	}

	value, err := shared.ParseDateTime("15:04", str)
	if err != nil {
		return err
	}

	if value != nil {
		*t = Time{
			Time:  *value,
			Valid: true,
		}
	}

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

func (input *Availabilities) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var value string
	err := decoder.DecodeElement(&value, &start)

	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	runes := []rune(value)
	out := make([]bool, 0, len(runes))

	for _, r := range runes {
		status, err := strconv.Atoi(string(r))

		if err != nil {
			return err
		}

		var v bool

		switch status {
		case 1:
			// 1 means available
			v = true
		case 2:
			// 2 means unavailable.
			v = false
		default:
			return errors.New("failed to parse availability")
		}

		out = append(out, v)
	}

	*input = out

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
	p.AllowElements("b", "br", "em", "i", "li", "ol", "p", "strong", "u", "ul")
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

func (input *LimitPerSeconds) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var str string
	err := decoder.DecodeElement(&str, &start)
	if err != nil {
		return err
	}

	if str == "" {
		*input = LimitPerSeconds{}
		return nil
	}

	splitStr := strings.Split(str, "/")
	if len(splitStr) != 2 {
		return errors.New("failed to parse rate_limit.limit")
	}

	limit, err := strconv.Atoi(splitStr[0])
	if err != nil {
		return err
	}

	secs, err := strconv.Atoi(splitStr[1])
	if err != nil {
		return err
	}

	*input = LimitPerSeconds{
		Requests: limit,
		Duration: time.Duration(secs) * time.Second,
	}

	return nil
}

func (input *Duration) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var str string
	err := decoder.DecodeElement(&str, &start)
	if err != nil {
		return err
	}

	if str == "" {
		return nil
	}

	secs, err := strconv.Atoi(str)
	if err != nil {
		return err
	}

	*input = Duration{time.Duration(secs) * time.Second}

	return nil
}

func (u *pictureURL) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var s string

	if err := decoder.DecodeElement(&s, &start); err != nil {
		return err
	}

	parsed, err := url.Parse(s)
	if err != nil {
		return err
	}

	parsed.Host = "cdn.easychannel.it"
	*u = pictureURL{parsed}

	return nil
}
