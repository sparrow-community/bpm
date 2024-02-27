package instance

type Resource struct {
	RootElement
	Name              string              `xml:"name,attr"`
	ResourceParameter []ResourceParameter `xml:"resourceParameter"`
}
