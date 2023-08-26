package instance

type ConversationLink struct {
	BaseElement
	Name      string `xml:"name,attr"`
	SourceRef string `xml:"sourceRef,attr"`
	TargetRef string `xml:"targetRef,attr"`
}
