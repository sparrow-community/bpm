package instance

type ConversationNode struct {
	BaseElement
	ParticipantRefs []string         `xml:"participantRef,attr"`
	MessageFlowRef  []string         `xml:"messageFlowRef,attr"`
	CorrelationKeys []CorrelationKey `xml:"correlationKey"`
}

type ConversationNodes struct {
}
