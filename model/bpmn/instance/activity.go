package instance

type Activity struct {
	FlowNode
	IsForCompensation      bool                    `xml:"isForCompensation,attr"`
	StartQuantity          int                     `xml:"startQuantity,attr"`
	CompletionQuantity     int                     `xml:"completionQuantity,attr"`
	Default                string                  `xml:"default,attr"`
	IoSpecification        IoSpecification         `xml:"ioSpecification"`
	Properties             []Property              `xml:"property"`
	DataInputAssociations  []DataInputAssociation  `xml:"dataInputAssociation"`
	DataOutputAssociations []DataOutputAssociation `xml:"dataOutputAssociation"`
	ResourceRoles          []ResourceRole          `xml:"resourceRole"`
	LoopCharacteristicsElements
}
