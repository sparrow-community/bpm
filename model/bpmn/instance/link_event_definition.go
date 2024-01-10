package instance

type LinkEventDefinition struct {
	EventDefinition
	Name    string   `xml:"name,attr"`
	Srouces []string `xml:"source,attr"`
	Target  string   `xml:"target,attr"`
}
