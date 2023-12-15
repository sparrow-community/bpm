package instance

type ThrowEvent struct {
	Event
	DataInputs            []DataInput            `xml:"dataInput"`
	DataInputAssociations []DataInputAssociation `xml:"dataInputAssociation"`
	InputSet              InputSet               `xml:"inputSet"`
	EventDefinitions
	EventDefinitionRefs []string `xml:"eventDefinitionRef"`
}
