package instance

type ReceiveTask struct {
	Task
	Implementation string `xml:"implementation,attr"`
	Instantiate    bool   `xml:"instantiate,attr"`
	MessageRef     string `xml:"messageRef,attr"`
	OperationRef   string `xml:"operationRef,attr"`
}
