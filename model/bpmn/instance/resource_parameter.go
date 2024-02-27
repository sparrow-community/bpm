package instance

type ResourceParameter struct {
	BaseElement
	Name       string `xml:"name,attr"`
	Type       string `xml:"type,attr"`
	IsRequired bool   `xml:"isRequired,attr"`
}
