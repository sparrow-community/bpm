package instance

type FlowElement struct {
	BaseElement
	Name              string      `xml:"name,attr"`
	Auditing          *Auditing   `xml:"auditing"`
	Monitoring        *Monitoring `xml:"monitoring"`
	CategoryValueRefs []string    `xml:"categoryValueRef"`
}

type FlowElements struct {
	StartEvents         []StartEvent         `xml:"startEvent"`
	EndEvents           []EndEvent           `xml:"endEvent"`
	Tasks               []Task               `xml:"task"`
	ServiceTasks        []ServiceTask        `xml:"serviceTask"`
	SequenceFlows       []SequenceFlow       `xml:"sequenceFlow"`
	DataStoreReferences []DataStoreReference `xml:"dataStoreReference"`
	ExclusiveGatewaies  []ExclusiveGateway   `xml:"exclusiveGateway"`
}
