package bpmn

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sparrow-community/bpm/model/bpmn/instance"
)

func TestA_2_1_roundtrip(t *testing.T) {
	// create test using ./test/A.2.1-roundtrip.bpmn
	path := "./test/A.2.1-roundtrip.bpmn"
	modelInstance, err := BpmnModelInstanceFromFile(path)
	if err != nil {
		t.Fatalf("BpmnModelInstanceFromFile error %s: %s", path, err)
	}
	if modelInstance == nil {
		t.Fatalf("modelInstance is nil, %s", path)
	}

	expected := &BpmnModelInstance{
		Definitions: &instance.Definitions{
			XMLName: xml.Name{
				Space: "http://www.omg.org/spec/BPMN/20100524/MODEL",
				Local: "definitions",
			},
			ID:                 "Bpmn_Definitions_--SwsH2BEeWQ6qGdY3x14w",
			Name:               "A.2.1",
			TargetNamespace:    "http://bonitasoft.com/_To9ZoDOCEeSknpIVFCxNIQ",
			ExpressionLanguage: "http://groovy.codehaus.org/",
			Exporter:           "BonitaSoft",
			ExporterVersion:    "6.3.3",
			RootElemnts: instance.RootElemnts{
				Processes: []instance.Process{
					{
						CallableElement: instance.CallableElement{
							RootElement: instance.RootElement{
								BaseElement: instance.BaseElement{
									ID: "_To9ZoTOCEeSknpIVFCxNIQ",
									ExtensionElements: instance.ExtensionElements{
										Any: []xml.Token{nil},
									},
									Any: []xml.Token{nil},
								},
							},
							Name: "A.2.1",
						},
						IsExecutable: false,
						ProcessType:  instance.TProcessTypeNone,
						FlowElements: instance.FlowElements{
							StartEvents: []instance.StartEvent{{
								CatchEvent: instance.CatchEvent{
									Event: instance.Event{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID:                "_To9ZojOCEeSknpIVFCxNIQ",
													ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil}},
												},
												Name: "Start Event",
											},
											Outgoing: []string{"_To9Z5DOCEeSknpIVFCxNIQ"},
										},
									},
								},
							}},
							Tasks: []instance.Task{
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID:                "_To9ZpzOCEeSknpIVFCxNIQ",
													ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil}},
												},
												Name: "Task 1",
											},
											Incoming: []string{"_To9Z5DOCEeSknpIVFCxNIQ"},
											Outgoing: []string{"_To9Z5zOCEeSknpIVFCxNIQ"},
										},
									},
								},
								{
									Activity: instance.Activity{
										Default: "Bpmn_SequenceFlow_edepQQbbEealeL5I4Yl3Dw",
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID:                "_To9ZtjOCEeSknpIVFCxNIQ",
													ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil}},
												},
												Name: "Task 2",
											},
											Incoming: []string{"_To9Z6jOCEeSknpIVFCxNIQ"},
											Outgoing: []string{"_To9Z7TOCEeSknpIVFCxNIQ", "Bpmn_SequenceFlow_edepQQbbEealeL5I4Yl3Dw"},
										},
									},
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID:                "_To9ZwDOCEeSknpIVFCxNIQ",
													ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil}},
												},
												Name: "Task 3",
											},
											Incoming: []string{"_To9Z-TOCEeSknpIVFCxNIQ", "Bpmn_SequenceFlow_edepQQbbEealeL5I4Yl3Dw", "Bpmn_SequenceFlow_f9nmUQbbEealeL5I4Yl3Dw"},
											Outgoing: []string{"_To9Z8DOCEeSknpIVFCxNIQ"},
										},
									},
								},
								{
									Activity: instance.Activity{
										Default: "Bpmn_SequenceFlow_f9nmUQbbEealeL5I4Yl3Dw",
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID:                "_To9ZzzOCEeSknpIVFCxNIQ",
													ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil}},
												},
												Name: "Task 4",
											},
											Incoming: []string{"_To9Z_DOCEeSknpIVFCxNIQ"},
											Outgoing: []string{"_To9Z8zOCEeSknpIVFCxNIQ", "Bpmn_SequenceFlow_f9nmUQbbEealeL5I4Yl3Dw"},
										},
									},
								},
							},
							ExclusiveGatewaies: []instance.ExclusiveGateway{
								{
									Default: "_To9Z6jOCEeSknpIVFCxNIQ",
									Gateway: instance.Gateway{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID:                "_To9ZyjOCEeSknpIVFCxNIQ",
													ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil}},
												},
												Name: "Gateway\n\n(Split Flow)",
											},
											Incoming: []string{"_To9Z5zOCEeSknpIVFCxNIQ"},
											Outgoing: []string{"_To9Z6jOCEeSknpIVFCxNIQ", "_To9Z-TOCEeSknpIVFCxNIQ", "_To9Z_DOCEeSknpIVFCxNIQ"},
										},
									},
								},
								{
									Gateway: instance.Gateway{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID:                "_To9Z2TOCEeSknpIVFCxNIQ",
													ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil}},
												},
												Name: "Gateway\n(Merge Flows)",
											},
											Incoming: []string{"_To9Z8DOCEeSknpIVFCxNIQ", "_To9Z8zOCEeSknpIVFCxNIQ"},
											Outgoing: []string{"_To9Z9jOCEeSknpIVFCxNIQ"},
										},
									},
								},
							},
							SequenceFlows: []instance.SequenceFlow{
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID:                "_To9Z5DOCEeSknpIVFCxNIQ",
											ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil}},
										},
									},
									SourceRef: "_To9ZojOCEeSknpIVFCxNIQ",
									TargetRef: "_To9ZpzOCEeSknpIVFCxNIQ",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID:                "_To9Z5zOCEeSknpIVFCxNIQ",
											ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil}},
										},
									},
									SourceRef: "_To9ZpzOCEeSknpIVFCxNIQ",
									TargetRef: "_To9ZyjOCEeSknpIVFCxNIQ",
								},
								{
									FlowElement: instance.FlowElement{
										Name: "Default",
										BaseElement: instance.BaseElement{
											ID:                "_To9Z6jOCEeSknpIVFCxNIQ",
											ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil}},
										},
									},
									SourceRef: "_To9ZyjOCEeSknpIVFCxNIQ",
									TargetRef: "_To9ZtjOCEeSknpIVFCxNIQ",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID:                "_To9Z7TOCEeSknpIVFCxNIQ",
											ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil}},
										},
										Name: "Condition",
									},
									SourceRef: "_To9ZtjOCEeSknpIVFCxNIQ",
									TargetRef: "_To9ZsTOCEeSknpIVFCxNIQ",
									ConditionExpression: instance.ExpressionUnMarshal{
										Type: instance.ExpressionTypeFormal,
										ExpressionSubstitution: &instance.FormalExpression{
											Expression: instance.Expression{
												BaseElementWithMixedContent: instance.BaseElementWithMixedContent{
													ID: "_cVKUxjOCEeSknpIVFCxNIQ",
												},
											},
											Language: "http://www.w3.org/1999/XPath",
										},
									},
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID:                "_To9Z8DOCEeSknpIVFCxNIQ",
											ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil}},
										},
									},
									SourceRef: "_To9ZwDOCEeSknpIVFCxNIQ",
									TargetRef: "_To9Z2TOCEeSknpIVFCxNIQ",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID:                "_To9Z8zOCEeSknpIVFCxNIQ",
											ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil}},
										},
										Name: "condition",
									},
									SourceRef: "_To9ZzzOCEeSknpIVFCxNIQ",
									TargetRef: "_To9Z2TOCEeSknpIVFCxNIQ",
									ConditionExpression: instance.ExpressionUnMarshal{
										Type: instance.ExpressionTypeFormal,
										ExpressionSubstitution: &instance.FormalExpression{
											Expression: instance.Expression{
												BaseElementWithMixedContent: instance.BaseElementWithMixedContent{
													ID: "_cVKUxjOCEeSknpIVFCxNIQ",
												},
											},
											Language: "http://www.w3.org/1999/XPath",
										},
									},
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID:                "_To9Z9jOCEeSknpIVFCxNIQ",
											ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil}},
										},
									},
									SourceRef: "_To9Z2TOCEeSknpIVFCxNIQ",
									TargetRef: "_To9ZsTOCEeSknpIVFCxNIQ",
									ConditionExpression: instance.ExpressionUnMarshal{
										Type: instance.ExpressionTypeFormal,
										ExpressionSubstitution: &instance.FormalExpression{
											Expression: instance.Expression{
												BaseElementWithMixedContent: instance.BaseElementWithMixedContent{
													ID: "_cVKUxjOCEeSknpIVFCxNIQ",
												},
											},
											Language: "http://www.w3.org/1999/XPath",
										},
									},
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID:                "_To9Z-TOCEeSknpIVFCxNIQ",
											ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil}},
										},
									},
									SourceRef: "_To9ZyjOCEeSknpIVFCxNIQ",
									TargetRef: "_To9ZwDOCEeSknpIVFCxNIQ",
									ConditionExpression: instance.ExpressionUnMarshal{
										Type: instance.ExpressionTypeFormal,
										ExpressionSubstitution: &instance.FormalExpression{
											Expression: instance.Expression{
												BaseElementWithMixedContent: instance.BaseElementWithMixedContent{
													ID: "_cVKUxjOCEeSknpIVFCxNIQ",
												},
											},
											Language: "http://www.w3.org/1999/XPath",
										},
									},
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID:                "_To9Z_DOCEeSknpIVFCxNIQ",
											ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil}},
										},
									},
									SourceRef: "_To9ZyjOCEeSknpIVFCxNIQ",
									TargetRef: "_To9ZzzOCEeSknpIVFCxNIQ",
									ConditionExpression: instance.ExpressionUnMarshal{
										Type: instance.ExpressionTypeFormal,
										ExpressionSubstitution: &instance.FormalExpression{
											Expression: instance.Expression{
												BaseElementWithMixedContent: instance.BaseElementWithMixedContent{
													ID: "_cVKUxjOCEeSknpIVFCxNIQ",
												},
											},
											Language: "http://www.w3.org/1999/XPath",
										},
									},
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID:                "Bpmn_SequenceFlow_edepQQbbEealeL5I4Yl3Dw",
											ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil}},
										},
									},
									IsImmediate: true,
									SourceRef:   "_To9ZtjOCEeSknpIVFCxNIQ",
									TargetRef:   "_To9ZwDOCEeSknpIVFCxNIQ",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID:                "Bpmn_SequenceFlow_f9nmUQbbEealeL5I4Yl3Dw",
											ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil}},
										},
									},
									IsImmediate: true,
									SourceRef:   "_To9ZzzOCEeSknpIVFCxNIQ",
									TargetRef:   "_To9ZwDOCEeSknpIVFCxNIQ",
								},
							},
							EndEvents: []instance.EndEvent{{
								ThrowEvent: instance.ThrowEvent{
									Event: instance.Event{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID:                "_To9ZsTOCEeSknpIVFCxNIQ",
													ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil}},
												},
												Name: "End Event",
											},
											Incoming: []string{"_To9Z7TOCEeSknpIVFCxNIQ", "_To9Z9jOCEeSknpIVFCxNIQ"},
										},
									},
								},
							}},
						},
					},
				},
			},
		},
	}
	if diff := cmp.Diff(expected, modelInstance); diff != "" {
		t.Errorf("Unexpected result (-want, +got):\n%s", diff)
	}
}
