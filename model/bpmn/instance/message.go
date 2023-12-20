package instance

type Message struct {
	RootElement
	Name    string `xml:"name,attr"`
	ItemRef string `xml:"itemRef,attr"`
}
