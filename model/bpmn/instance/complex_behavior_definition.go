package instance

type ComplexBehaviorDefinition struct {
	BaseElement
	Condition ConditionExpression `xml:"condition"`
	Event     ImplicitThrowEvent  `xml:"event"`
}
