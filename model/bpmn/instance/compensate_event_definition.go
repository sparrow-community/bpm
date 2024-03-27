package instance

type CompensateEventDefinition struct {
	EventDefinition
	WaitForCompletion bool `xml:"waitForCompletion,attr"`
	ActivityRef       string
}
