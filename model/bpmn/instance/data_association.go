package instance

type DataAssociation struct {
	BaseElement
	SourceRef      string            `xml:"sourceRef"`
	TargetRef      string            `xml:"targetRef"`
	Transformation *FormalExpression `xml:"transformation"`
	Assignments    []*Assignment     `xml:"assignment"`
}
