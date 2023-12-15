package instance

// TimerEventDefinition, choice TimeDate, TimeDuration, TimeCycle
type TimerEventDefinition struct {
	EventDefinition
	TimeDate     Expression `xml:"timeDate"`
	TimeDuration Expression `xml:"timeDuration"`
	TimeCycle    Expression `xml:"timeCycle"`
}
