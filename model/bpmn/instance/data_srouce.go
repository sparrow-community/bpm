package instance

type DataStore struct {
	RootElement
	Name           string    `xml:"name,attr"`
	Capacity       int       `xml:"capacity,attr"`
	IsUnlimited    bool      `xml:"isUnlimited,attr"`
	ItemSubjectRef string    `xml:"itemSubjectRef,attr"`
	DataState      DataState `xml:"dataState"`
}
