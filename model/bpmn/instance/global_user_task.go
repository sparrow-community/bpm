package instance

type GlobalUserTask struct {
	GlobalTask
	Rendering      []Rendering `xml:"rendering"`
	Implementation string      `xml:"implementation,attr"`
}
