package instance

type UserTask struct {
	Task
	Rendering      []Rendering `xml:"rendering"`
	Implementation string      `xml:"implementation,attr"`
}
