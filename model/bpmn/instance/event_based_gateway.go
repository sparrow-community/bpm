package instance

type EventGatewayType string

const (
	EventGatewayTypeExclusive EventGatewayType = "Exclusive"
	EventGatewayTypeParallel  EventGatewayType = "Parallel"
)

type EventBasedGateway struct {
	Gateway
	Instantiate      bool             `xml:"instantiate,attr"`
	EventGatewayType EventGatewayType `xml:"eventGatewayType,attr"`
}
