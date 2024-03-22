package instance

type ExpressionType string

const (
	ExpressionTypeDefault ExpressionType = ""
	ExpressionTypeFormal  ExpressionType = "tFormalExpression"
)

type ExpressionSubstitutionGroup interface {
	Type() ExpressionType
}

var ExpressionTypeRegister = map[ExpressionType]func() ExpressionSubstitutionGroup{
	ExpressionTypeDefault: func() ExpressionSubstitutionGroup {
		return &Expression{}
	},
	ExpressionTypeFormal: func() ExpressionSubstitutionGroup {
		return &FormalExpression{}
	},
}

func GetExpressionSubstitutionGroup(t ExpressionType) ExpressionSubstitutionGroup {
	if factory, ok := ExpressionTypeRegister[t]; ok {
		return factory()
	}
	return &Expression{}
}

type Expression struct {
	BaseElementWithMixedContent
}

func (e *Expression) Type() ExpressionType {
	return ExpressionTypeDefault
}
