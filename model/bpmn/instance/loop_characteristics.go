package instance

type LoopCharacteristics struct {
	BaseElement
}

type LoopCharacteristicsElements struct {
	StandardLoopCharacteristics      []StandardLoopCharacteristics      `xml:"standardLoopCharacteristics"`
	MultiInstanceLoopCharacteristics []MultiInstanceLoopCharacteristics `xml:"multiInstanceLoopCharacteristics"`
}
