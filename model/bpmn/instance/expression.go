package instance

type ExpressionType string

const (
	ExpressionTypeDefault ExpressionType = ""
	ExpressionTypeFormal  ExpressionType = "tFormalExpression"
)

type ExpressionSubstitutionGroup interface {
	Type() ExpressionType
}

var ExpressionTypeRegister = map[ExpressionType]ExpressionSubstitutionGroup{
	ExpressionTypeDefault: &Expression{},
	ExpressionTypeFormal:  &FormalExpression{},
}

type Expression struct {
	BaseElementWithMixedContent
}

func (e *Expression) Type() ExpressionType {
	return ExpressionTypeDefault
}
