package instance

type ConditionalEventDefinition struct {
	EventDefinition
	Condition ExpressionUnMarshal `xml:"condition"`
}
