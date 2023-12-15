package instance

type DataObject struct {
	FlowElement
	DataState    DataState `xml:"dataState"`
	ItemSubject  string    `xml:"itemSubject"`
	IsCollection bool      `xml:"isCollection"`
}
