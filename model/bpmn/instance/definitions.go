package instance

import "encoding/xml"

type Definitions struct {
	XMLName            xml.Name `xml:"definitions"`
	ID                 string   `xml:"id,attr"`
	Name               string   `xml:"name,attr"`
	TargetNamespace    string   `xml:"targetNamespace,attr"`
	ExpressionLanguage string   `xml:"expressionLanguage,attr"`
	TypeLanguage       string   `xml:"typeLanguage,attr"`
	Exporter           string   `xml:"exporter,attr"`
	ExporterVersion    string   `xml:"exporterVersion,attr"`
	Rootlemnts
}
