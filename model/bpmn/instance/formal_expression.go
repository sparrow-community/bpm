package instance

type FormalExpression struct {
	Expression
	Language           string `xml:"language,attr"`
	EvaluatesToTypeRef string `xml:"evaluatesToTypeRef,attr"`
}
