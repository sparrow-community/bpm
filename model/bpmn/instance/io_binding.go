package instance

type IoBinding struct {
	InputOutputBinding
}

type InputOutputBinding struct {
	BaseElement
	DataInputs  []DataInput  `xml:"dataInput"`
	DataOutputs []DataOutput `xml:"dataOutput"`
	InputSets   []InputSet   `xml:"inputSet"`
	OutputSets  []OutputSet  `xml:"outputSet"`
}
