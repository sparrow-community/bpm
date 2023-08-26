package instance

type Property struct {
	BaseElement
	Name           string `xml:"name,attr"`
	ItemSubjectRef string `xml:"itemSubjectRef,attr"`
	// child elements
	DataState *DataState `xml:"dataState"`
}
