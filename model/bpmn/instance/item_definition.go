package instance

type ItemKind string

const (
	ItemKindInformation ItemKind = "Information"
	ItemKindPhysical    ItemKind = "Physical"
)

type ItemDefinition struct {
	RootElement
	StructureRef string   `xml:"structureRef,attr"`
	IsCollection bool     `xml:"isCollection,attr"`
	ItemKind     ItemKind `xml:"itemKind,attr"`
}
