package instance

type EscalationEventDefinition struct {
	EventDefinition
	EscalationRef string `xml:"escalationRef,attr"`
}
