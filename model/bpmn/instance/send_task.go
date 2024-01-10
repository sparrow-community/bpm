package instance

type SendTask struct {
	Task
	Implementation string `xml:"implementation,attr"`
	MessageRef     string `xml:"messageRef,attr"`
	OperationRef   string `xml:"operationRef,attr"`
}
