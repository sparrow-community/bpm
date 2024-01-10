package instance

type InclusiveGateway struct {
	Gateway
	Default string `xml:"default,attr"`
}
