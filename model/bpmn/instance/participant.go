package instance

type Participant struct {
	BaseElement
	Name                    string                   `xml:"name,attr"`
	ProcessRef              string                   `xml:"processRef,attr"`
	InterfaceRefs           []string                 `xml:"interfaceRef"`
	EndPointRefs            []string                 `xml:"endPointRef"`
	ParticipantMultiplicity *ParticipantMultiplicity `xml:"participantMultiplicity"`
}
