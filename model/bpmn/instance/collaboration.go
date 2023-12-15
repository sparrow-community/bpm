package instance

type Collaboration struct {
	RootElement
	Name         string        `xml:"name,attr"`
	IsClosed     bool          `xml:"isClosed,attr"`
	Participants []Participant `xml:"participant"`
	MessageFlows []MessageFlow `xml:"messageFlow"`
	Artifacts
	ConversationNodes
	ConversationAssociations []ConversationAssociation `xml:"conversationAssociation"`
	ParticipantAssociations  []ParticipantAssociation  `xml:"participantAssociation"`
	MessageFlowAssociations  []MessageFlowAssociation  `xml:"messageFlowAssociation"`
	CorrelationKeys          []CorrelationKey          `xml:"correlationKey"`
	ChoreographyRef          []string                  `xml:"choreographyRef"`
	ConversationLinks        []ConversationLink        `xml:"conversationLink"`
}
