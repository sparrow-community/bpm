package instance

type ConditionalEventDefinition struct {
	EventDefinition
	Condition ConditionExpression `xml:"condition"`
}
