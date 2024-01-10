package instance

type Artifact struct {
	BaseElement
}

type Artifacts struct {
	TextAnnotations []TextAnnotation `xml:"textAnnotation"`
	Associations    []Association    `xml:"association"`
	Groups          []Group          `xml:"group"`
}
