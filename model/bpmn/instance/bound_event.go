package instance

type BoundaryEvent struct {
	CatchEvent
	AttachedToRef  string `xml:"attachedToRef,attr"`
	CancelActivity bool   `xml:"cancelActivity,attr"`
}
