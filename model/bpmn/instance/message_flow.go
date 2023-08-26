package instance

type MessageFlow struct {
	BaseElement
	Name       string `xml:"name,attr"`
	SourceRef  string `xml:"sourceRef,attr"`
	TargetRef  string `xml:"targetRef,attr"`
	MessageRef string `xml:"messageRef,attr"`
}
