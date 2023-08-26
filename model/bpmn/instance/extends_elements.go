package instance

import "encoding/xml"

type ExtensionElements struct {
	Any []xml.Token `xml:",any"`
}
