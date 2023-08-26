package instance

type ResourceRole struct {
	BaseElement
	Name                         string                       `xml:"name,attr"`
	ResourceRef                  string                       `xml:"resourceRef,attr"`
	ResourceParameterBindings    []ResourceParameterBinding   `xml:"resrouceParameterBinding"`
	ResourceAssignmentExpression ResourceAssignmentExpression `xml:"resourceAssignmentExpression"`
}
