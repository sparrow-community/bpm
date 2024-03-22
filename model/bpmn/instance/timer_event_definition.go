package instance

// TimerEventDefinition, choice TimeDate, TimeDuration, TimeCycle
type TimerEventDefinition struct {
	EventDefinition
	TimeDate     ExpressionUnMarshal `xml:"timeDate"`
	TimeDuration ExpressionUnMarshal `xml:"timeDuration"`
	TimeCycle    ExpressionUnMarshal `xml:"timeCycle"`
}
