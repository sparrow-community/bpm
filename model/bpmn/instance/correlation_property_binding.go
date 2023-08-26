package instance

type CorrelationPropertyBinding struct {
	BaseElement
	CorrelationPropertyRef string            `xml:"correlationPropertyRef,attr"`
	DataPath               *FormalExpression `xml:"dataPath,attr"`
}
