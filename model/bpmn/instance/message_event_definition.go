package instance

type MessageEventDefinition struct {
	EventDefinition
	MessageRef   string `xml:"messageRef,attr"`
	OperationRef string `xml:"operationRef"`
}
