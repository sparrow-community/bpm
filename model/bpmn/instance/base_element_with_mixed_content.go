package instance

type BaseElementWithMixedContent struct {
	ID                string            `xml:"id,attr"`
	Documentation     []Documentation   `xml:"documentation"`
	ExtensionElements ExtensionElements `xml:"extensionElements"`
}
