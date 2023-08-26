package instance

type ExclusiveGateway struct {
	Gateway
	Default string `xml:"default,attr"`
}
