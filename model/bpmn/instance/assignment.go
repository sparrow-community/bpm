package instance

type Assignment struct {
	BaseElement
	From Expression `xml:"from,attr"`
	To   Expression `xml:"to,attr"`
}
