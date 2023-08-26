package instance

type TProcessType string

const (
	TProcessTypeNone    TProcessType = "None"
	TProcessTypePublic  TProcessType = "Public"
	TProcessTypePrivate TProcessType = "Private"
)

type Process struct {
	CallableElement
	ProcessType  TProcessType `xml:"processType,attr,omitempty"`
	IsClosed     bool         `xml:"isClosed,attr,omitempty"`
	IsExecutable bool         `xml:"isExecutable,attr,omitempty"`
	Auditing     *Auditing    `xml:"auditing,omitempty"`
	Monitoring   *Monitoring  `xml:"monitoring,omitempty"`
	Properties   []Property   `xml:"property,omitempty"`
	LaneSets     []LaneSet    `xml:"laneSet,omitempty"`
	FlowElements
	Artifacts
	ResourceRoles            []ResourceRole            `xml:"resourceRole,omitempty"`
	CorrelationSubscriptions []CorrelationSubscription `xml:"correlationSubscription,omitempty"`
	Supports                 []string                  `xml:"supports,omitempty"`
}
