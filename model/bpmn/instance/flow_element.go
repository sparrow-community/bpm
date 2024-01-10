package instance

type FlowElement struct {
	BaseElement
	Name              string     `xml:"name,attr"`
	Auditing          Auditing   `xml:"auditing"`
	Monitoring        Monitoring `xml:"monitoring"`
	CategoryValueRefs []string   `xml:"categoryValueRef"`
}

type FlowElements struct {
	StartEvents             []StartEvent             `xml:"startEvent"`
	EndEvents               []EndEvent               `xml:"endEvent"`
	Tasks                   []Task                   `xml:"task"`
	UserTasks               []UserTask               `xml:"userTask"`
	ServiceTasks            []ServiceTask            `xml:"serviceTask"`
	SendTasks               []SendTask               `xml:"sendTask"`
	ReceiveTasks            []ReceiveTask            `xml:"receiveTask"`
	SequenceFlows           []SequenceFlow           `xml:"sequenceFlow"`
	DataStoreReferences     []DataStoreReference     `xml:"dataStoreReference"`
	ParallelGatewaies       []ParallelGateway        `xml:"parallelGateway"`
	ExclusiveGatewaies      []ExclusiveGateway       `xml:"exclusiveGateway"`
	InclusiveGatewaies      []InclusiveGateway       `xml:"inclusiveGateway"`
	EventBasedGatewaies     []EventBasedGateway      `xml:"eventBasedGateway"`
	SubProcess              []SubProcess             `xml:"subProcess"`
	BoundaryEvents          []BoundaryEvent          `xml:"boundaryEvent"`
	CallActivities          []CallActivity           `xml:"callActivity"`
	DataObjectReferenes     []DataObjectReference    `xml:"dataObjectReference"`
	DataObjects             []DataObject             `xml:"dataObject"`
	IntermediateThrowEvents []IntermediateThrowEvent `xml:"intermediateThrowEvent"`
	IntermediateCatchEvents []IntermediateCatchEvent `xml:"intermediateCatchEvent"`
	ImplicitThrowEvents     []ImplicitThrowEvent     `xml:"implicitThrowEvent"`
}
