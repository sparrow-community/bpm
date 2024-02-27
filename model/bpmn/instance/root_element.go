package instance

type RootElementInterface interface {
	BaseElementInterface
}

type RootElement struct {
	BaseElement
}

type RootElemnts struct {
	Collaborations  []Collaboration  `xml:"collaboration"`
	Processes       []Process        `xml:"process"`
	DataStores      []DataStore      `xml:"dataStore"`
	Categories      []Category       `xml:"category"`
	GlobalTasks     []GlobalTask     `xml:"globalTask"`
	GlobalUserTasks []GlobalUserTask `xml:"globalUserTask"`
	Messages        []Message        `xml:"message"`
	Resources       []Resource       `xml:"resource"`
}
