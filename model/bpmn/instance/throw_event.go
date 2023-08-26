package instance

type ThrowEvent struct {
	Event
	DataInputs            []DataInput            `xml:"dataInput"`
	DataInputAssociations []DataInputAssociation `xml:"dataInputAssociation"`
	InputSet              *InputSet              `xml:"inputSet"`
	EventDefinitions      []EventDefinition      `xml:"eventDefinition"`
	EventDefinitionRefs   []string               `xml:"eventDefinitionRef"`
}
