package instance

type CatchEvent struct {
	Event
	ParallelMultiple       bool                    `xml:"parallelMultiple,attr"`
	DataOutputs            []DataOutput            `xml:"dataOutput"`
	DataOutputAssociations []DataOutputAssociation `xml:"dataOutputAssociation"`
	OutputSet              OutputSet               `xml:"outputSet"`
	EventDefinitions
}
