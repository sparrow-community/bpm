package instance

type ResourceRoles struct {
	ResourceRoles   []ResourceRole   `xml:"resourceRole"`
	Performers      []Performer      `xml:"performer"`
	PotentialOwners []PotentialOwner `xml:"potentialOwner"`
	HumanPerformers []HumanPerformer `xml:"humanPerformer"`
}
