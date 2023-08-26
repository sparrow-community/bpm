package instance

type MessageFlowAssociation struct {
	BaseElement
	InnerMessageFlowRef string `xml:"innerMessageFlowRef,attr"`
	OuterMessageFlowRef string `xml:"outerMessageFlowRef,attr"`
}
