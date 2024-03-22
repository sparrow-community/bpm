package instance

type ComplexBehaviorDefinition struct {
	BaseElement
	Condition ExpressionUnMarshal `xml:"condition"`
	Event     ImplicitThrowEvent  `xml:"event"`
}
