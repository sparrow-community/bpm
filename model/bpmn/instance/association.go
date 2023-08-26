package instance

type AssociationDirection string

const (
	AssociationDirectionNone AssociationDirection = "None"
	AssociationDirectionOne  AssociationDirection = "One"
	AssociationDirectionBoth AssociationDirection = "Both"
)

type Association struct {
	Artifact
	SourceRef            string               `xml:"sourceRef,attr"`
	TargetRef            string               `xml:"targetRef,attr"`
	AssociationDirection AssociationDirection `xml:"associationDirection,attr"`
}
