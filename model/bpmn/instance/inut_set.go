package instance

type InputSet struct {
	BaseElement
	Name                    string   `xml:"name,attr"`
	DataInputRefs           []string `xml:"dataInputRefs"`
	OptionalInputRefs       []string `xml:"optionalInputRefs"`
	WhileExecutingInputRefs []string `xml:"whileExecutingInputRefs"`
	OutputSetRefs           []string `xml:"outputSetRefs"`
}
