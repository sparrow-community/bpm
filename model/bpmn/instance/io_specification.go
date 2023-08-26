package instance

type IoSpecification struct {
	InputOutputSpecification
}

type InputOutputSpecification struct {
	BaseElement
	DataInputs  []DataInput  `xml:"dataInput"`
	DataOutputs []DataOutput `xml:"dataOutput"`
	InputSets   []InputSet   `xml:"inputSet"`
	OutputSets  []OutputSet  `xml:"outputSet"`
}
