package instance

type Lane struct {
	BaseElement
	Name                string      `xml:"name,attr"`
	PartitionElementRef string      `xml:"partitionElementRef,attr"`
	PartitionElement    BaseElement `xml:"partitionElement"`
	FlowNodeRefs        []string    `xml:"flowNodeRef"`
	ChildLaneSet        LaneSet     `xml:"childLaneSet"`
}

type LaneSet struct {
	BaseElement
	Name  string `xml:"name,attr"`
	Lanes []Lane `xml:"lane"`
}
