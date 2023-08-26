package instance

type GatewayDirection string

const (
	GatewayDirectionUnspecified GatewayDirection = "Unspecified"
	GatewayDirectionConverging  GatewayDirection = "Converging"
	GatewayDirectionDiverging   GatewayDirection = "Diverging"
	GatewayDirectionMixed       GatewayDirection = "Mixed"
)

type Gateway struct {
	FlowNode
	GatewayDirection GatewayDirection `xml:"gatewayDirection,attr"`
}
