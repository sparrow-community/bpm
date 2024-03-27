package instance

type SubProcess struct {
	Activity
	LaneSets []LaneSet `xml:"laneSet,omitempty"`
	FlowElements
	Artifacts
	TriggeredByEvent bool `xml:"triggeredByEvent,attr,omitempty"`
}
