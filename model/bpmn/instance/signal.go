package instance

type Signal struct {
	RootElement
	Name         string `xml:"name,attr"`
	StructureRef string `xml:"structureRef,attr"`
}
