package instance

type Group struct {
	Artifact
	CategoryValueRef string `xml:"categoryValueRef,attr"`
}
