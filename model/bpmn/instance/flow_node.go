package instance

type FlowNode struct {
	FlowElement
	Incoming []string `xml:"incoming"`
	Outgoing []string `xml:"outgoing"`
}
