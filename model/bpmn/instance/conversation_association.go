package instance

type ConversationAssociation struct {
	BaseElement
	InnerConversationNodeRef string `xml:"innerConversationNodeRef,attr"`
	OuterConversationNodeRef string `xml:"outerConversationNodeRef,attr"`
}
