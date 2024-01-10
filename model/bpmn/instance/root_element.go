package instance

type RootElementInterface interface {
	BaseElementInterface
}

type RootElement struct {
	BaseElement
}

type RootElemnts struct {
	Collaborations []Collaboration `xml:"collaboration"`
	Processes      []Process       `xml:"process"`
	DataStores     []DataStore     `xml:"dataStore"`
	Categories     []Category      `xml:"category"`
	GlobalTasks    []GlobalTask    `xml:"globalTask"`
	Messages       []Message       `xml:"message"`
}
