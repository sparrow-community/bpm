package instance

type StartEvent struct {
	CatchEvent
	Interrupting bool `xml:"interrupting,attr"`
}
