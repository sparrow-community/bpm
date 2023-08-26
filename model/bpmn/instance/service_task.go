package instance

type ServiceTask struct {
	Task
	Implementation string `xml:"implementation,attr"`
	OperationRef   string `xml:"operationRef,attr"`
}
