package instance

type MultiInstanceFlowCondition string

const (
	MultiInstanceFlowConditionNone    MultiInstanceFlowCondition = "None"
	MultiInstanceFlowConditionOne     MultiInstanceFlowCondition = "One"
	MultiInstanceFlowConditionAll     MultiInstanceFlowCondition = "All"
	MultiInstanceFlowConditionComplex MultiInstanceFlowCondition = "Complex"
)

type MultiInstanceLoopCharacteristics struct {
	LoopCharacteristics
	IsSequential               bool                        `xml:"isSequential,attr"`
	Behavior                   MultiInstanceFlowCondition  `xml:"behavior,attr"`
	OneBehaviorEventRef        string                      `xml:"oneBehaviorEventRef,attr"`
	NoneBehaviorEventRef       string                      `xml:"noneBehaviorEventRef,attr"`
	LoopCardinality            ExpressionUnMarshal         `xml:"loopCardinality"`
	LoopDataInputRef           string                      `xml:"loopDataInputRef,attr"`
	LoopDataOutputRef          string                      `xml:"loopDataOutputRef,attr"`
	InputDataItem              DataInput                   `xml:"inputDataItem"`
	OutputDataItem             DataOutput                  `xml:"outputDataItem"`
	ComplexBehaviorDefinitions []ComplexBehaviorDefinition `xml:"complexBehaviorDefinition"`
	CompletionCondition        ExpressionUnMarshal         `xml:"completionCondition"`
}
