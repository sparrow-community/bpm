package instance

type RootElementInterface interface {
	BaseElementInterface
}

type RootElement struct {
	BaseElement
}

type Rootlemnts struct {
	Collaborations []Collaboration `xml:"collaboration"`
	Processes      []Process       `xml:"process"`
	DataStore      []DataStore     `xml:"dataStore"`
}
