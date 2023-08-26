package instance

type CorrelationKey struct {
	BaseElement
	Name                    string   `xml:"name,attr"`
	CorrelationPropertyRefs []string `xml:"correlationPropertyRef"`
}
