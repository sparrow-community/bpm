package instance

type CallableElement struct {
	RootElement
	Name            string          `xml:"name,attr"`
	IoSpecification IoSpecification `xml:"ioSpecification,omitempty"`
	IoBinding       []IoBinding     `xml:"ioBinding,omitempty"`
}
