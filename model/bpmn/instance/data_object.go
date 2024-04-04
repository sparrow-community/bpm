package instance

type DataObject struct {
	FlowElement
	DataState      DataState `xml:"dataState"`
	ItemSubjectRef string    `xml:"itemSubjectRef"`
	IsCollection   bool      `xml:"isCollection"`
}
