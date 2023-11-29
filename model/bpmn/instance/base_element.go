package instance

import "encoding/xml"

type BaseElementInterface interface {
	GetID() string
	GetDocumentation() []Documentation
	GetExtensionElements() ExtensionElements
}

type BaseElement struct {
	ID                string            `xml:"id,attr,omitempty"`
	Documentation     []Documentation   `xml:"documentation,omitempty"`
	ExtensionElements ExtensionElements `xml:"extensionElements,omitempty"`
	Any               []xml.Token       `xml:",any"`
}

func (b BaseElement) GetID() string {
	return b.ID
}

func (b BaseElement) GetDocumentation() []Documentation {
	return b.Documentation
}

func (b BaseElement) GetExtensionElements() ExtensionElements {
	return b.ExtensionElements
}
