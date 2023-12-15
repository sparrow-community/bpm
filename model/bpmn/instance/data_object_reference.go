package instance

type DataObjectReference struct {
	FlowElement
	DataState      DataState `xml:"dataState"`
	ItemSubjectRef string    `xml:"itemSubjectRef,attr"`
	DataObjectRef  string    `xml:"dataObjectRef,attr"`
}
