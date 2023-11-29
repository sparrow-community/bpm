package instance

type EventDefinition struct {
	RootElement
}

type EventDefinitions struct {
	MessageEventDefinitions    []MessageEventDefinition    `xml:"messageEventDefinition"`
	EscalationEventDefinitions []EscalationEventDefinition `xml:"escalationEventDefinition"`
}
