package instance

type Import struct {
	Namespace  string `xml:"namespace,attr,omitempty"`
	Location   string `xml:"location,attr,omitempty"`
	ImportType string `xml:"importType,attr,omitempty"`
}
