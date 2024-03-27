package instance

type BusinessRuleTask struct {
	Task
	Implementation string `xml:"implementation,attr"`
}
