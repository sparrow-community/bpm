package instance

type CorrelationSubscription struct {
	BaseElement
	CorrelationKey              string                       `xml:"correlationKey,attr"`
	CorrelationPropertyBindings []CorrelationPropertyBinding `xml:"correlationPropertyBinding"`
}
