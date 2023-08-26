package instance

type DataOutput struct {
	BaseElement
	Name           string     `xml:"name,attr"`
	ItemSubjectRef string     `xml:"itemSubjectRef,attr"`
	IsCollection   bool       `xml:"isCollection,attr"`
	DataState      *DataState `xml:"dataState,attr"`
}
