package instance

type Event struct {
	FlowNode
	Properties []Property `xml:"property"`
}
