package instance

type DataStoreReference struct {
	FlowElement
	ItemSubjectRef string    `xml:"itemSubjectRef,attr"`
	DataStoreRef   string    `xml:"dataStoreRef,attr"`
	DataState      DataState `xml:"dataState,attr"`
}
