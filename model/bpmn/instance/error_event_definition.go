package instance

type ErrorEventDefinition struct {
	EventDefinition
	ErrorRef string `xml:"errorRef,attr"`
}
