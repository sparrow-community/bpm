package instance

type TextAnnotation struct {
	Artifact
	TextFormat string `xml:"textFormat,attr"`
	Text       string `xml:"text"`
}
