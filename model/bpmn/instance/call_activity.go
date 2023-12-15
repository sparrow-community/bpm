package instance

type CallActivity struct {
	Activity
	CalledElement string `xml:"calledElement,attr"`
}
