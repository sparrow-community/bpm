package instance

type ResourceParameterBinding struct {
	BaseElement
	ParameterRef string `xml:"parameterRef,attr"`
	Expression   string `xml:"expression,attr"`
}
