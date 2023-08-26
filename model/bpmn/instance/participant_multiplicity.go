package instance

type ParticipantMultiplicity struct {
	BaseElement
	Minimum int `xml:"minimum,attr"`
	Maximum int `xml:"maximum,attr"`
}
