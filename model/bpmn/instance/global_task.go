package instance

type GlobalTask struct {
	CallableElement
	ResourceRoles []ResourceRole `xml:"resourceRole"`
}
