package instance

import (
	"encoding/xml"
	"strings"
)

type ExpressionUnMarshal struct {
	Type                   ExpressionType              `xml:"http://www.w3.org/2001/XMLSchema-instance type,attr"` // xsi:type
	ExpressionSubstitution ExpressionSubstitutionGroup `xml:",omitempty"`
}

func (e *ExpressionUnMarshal) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			parts := strings.Split(attr.Value, ":")
			if len(parts) > 1 {
				e.Type = ExpressionType(parts[1])
			} else {
				e.Type = ExpressionType(attr.Value)
			}
		}
	}
	if len(e.Type) == 0 {
		e.Type = ""
	}

	t := GetExpressionSubstitutionGroup(e.Type)
	if err := d.DecodeElement(t, &start); err != nil {
		return err
	}
	e.ExpressionSubstitution = t
	return nil
}
