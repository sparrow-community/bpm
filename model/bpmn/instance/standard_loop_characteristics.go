package instance

type StandardLoopCharacteristics struct {
	LoopCharacteristics
	TestBefore    bool         `xml:"testBefore,attr"`
	LoopMaximum   int          `xml:"loopMaximum,attr"`
	LoopCondition []Expression `xml:"loopCondition"`
}
