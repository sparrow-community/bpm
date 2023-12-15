package instance

type Artifact struct {
	BaseElement
}

type Artifacts struct {
	TextAnnotations []TextAnnotation `xml:"textAnnotation"`
	Associations    []Association    `xml:"association"`
	Group           []Group          `xml:"group"`
}
