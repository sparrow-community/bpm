package instance

type OutputSet struct {
	BaseElement
	Name                     string   `xml:"name,attr"`
	DataOutputRefs           []string `xml:"dataOutputRefs"`
	OptionalOutputRefs       []string `xml:"optionalOutputRefs"`
	WhileExecutingOutputRefs []string `xml:"whileExecutingOutputRefs"`
	InputSetRefs             []string `xml:"inputSetRefs"`
}
