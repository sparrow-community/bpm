package instance

import "encoding/xml"

type Documentation struct {
	Id         string      `xml:"id,attr,omitempty"`
	TextFormat string      `xml:"textFormat,attr,omitempty"`
	Value      string      `xml:",chardata"`
	Any        []xml.Token `xml:",any"`
}
