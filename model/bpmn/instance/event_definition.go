package instance

type EventDefinition struct {
	RootElement
}

type EventDefinitions struct {
	MessageEventDefinitions     []MessageEventDefinition     `xml:"messageEventDefinition"`
	EscalationEventDefinitions  []EscalationEventDefinition  `xml:"escalationEventDefinition"`
	TimerEventDefinitions       []TimerEventDefinition       `xml:"timerEventDefinition"`
	TerminateEventDefinitions   []TerminateEventDefinition   `xml:"terminateEventDefinition"`
	SignalEventDefinitions      []SignalEventDefinition      `xml:"signalEventDefinition"`
	ConditionalEventDefinitions []ConditionalEventDefinition `xml:"conditionalEventDefinition"`
	ErrorEventDefinitions       []ErrorEventDefinition       `xml:"errorEventDefinition"`
	LinkEventDefinitions        []LinkEventDefinition        `xml:"linkEventDefinition"`
	CompensateEventDefinitions  []CompensateEventDefinition  `xml:"compensateEventDefinition"`
}
