package instance

type SequenceFlow struct {
	FlowElement
	SourceRef           string              `xml:"sourceRef,attr"`
	TargetRef           string              `xml:"targetRef,attr"`
	IsImmediate         bool                `xml:"isImmediate,attr"`
	ConditionExpression ConditionExpression `xml:"conditionExpression"`
}
