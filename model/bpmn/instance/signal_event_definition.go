package instance

type SignalEventDefinition struct {
	EventDefinition
	SignalRef string `xml:"signalRef,attr"`
}
