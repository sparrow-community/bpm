package instance

type DataInput struct {
	BaseElement
	Name           string `xml:"name,attr"`
	ItemSubjectRef string `xml:"itemSubjectRef,attr"`
	IsCollection   bool   `xml:"isCollection,attr"`
	DataState      string `xml:"dataState"`
}
