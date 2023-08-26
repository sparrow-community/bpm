package instance

type ParticipantAssociation struct {
	BaseElement
	InnerParticipantRef string `xml:"innerParticipantRef,attr"`
	OuterParticipantRef string `xml:"outerParticipantRef,attr"`
}
