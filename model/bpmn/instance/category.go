package instance

type Category struct {
	RootElement
	CategoryValues []CategoryValue `xml:"categoryValue"`
	Name           string          `xml:"name,attr"`
}
