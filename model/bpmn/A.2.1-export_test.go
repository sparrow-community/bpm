package bpmn

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sparrow-community/bpm/model/bpmn/instance"
)

func TestA_2_1_export(t *testing.T) {
	// create test using ./test/A.2.1-export.bpmn
	path := "./test/A.2.1-export.bpmn"
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
			ID:              "Definitions_1mv2yrn",
			TargetNamespace: "http://bpmn.io/schema/bpmn",
			Exporter:        "Camunda Modeler",
			ExporterVersion: "5.0.0",
			RootElemnts: instance.RootElemnts{
				Processes: []instance.Process{
					{
						CallableElement: instance.CallableElement{
							RootElement: instance.RootElement{
								BaseElement: instance.BaseElement{
									ID: "Process_05abo3f",
								},
							},
						},
						IsExecutable: true,
						FlowElements: instance.FlowElements{
							StartEvents: []instance.StartEvent{{
								CatchEvent: instance.CatchEvent{
									Event: instance.Event{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "StartEvent_1",
												},
												Name: "Start Event",
											},
											Outgoing: []string{"Flow_0lacnkf"},
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
													ID: "Activity_0ahdk3x",
												},
												Name: "Task 1",
											},
											Incoming: []string{"Flow_0lacnkf"},
											Outgoing: []string{"Flow_187opjq"},
										},
									},
								},
								{
									Activity: instance.Activity{
										Default: "Flow_1sd2kpj",
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Activity_172ndxy",
												},
												Name: "Task 2",
											},
											Incoming: []string{"Flow_194jx6p"},
											Outgoing: []string{"Flow_1sd2kpj", "Flow_0zwjy3h"},
										},
									},
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Activity_1lz0l07",
												},
												Name: "Task 3",
											},
											Incoming: []string{"Flow_1xhc3bf", "Flow_1sd2kpj", "Flow_11m9fcl"},
											Outgoing: []string{"Flow_1a24ve1"},
										},
									},
								},
								{
									Activity: instance.Activity{
										Default: "Flow_11m9fcl",
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Activity_171qefk",
												},
												Name: "Task 4",
											},
											Incoming: []string{"Flow_19m0ydj"},
											Outgoing: []string{"Flow_11m9fcl", "Flow_01ckxme"},
										},
									},
								},
							},
							ExclusiveGatewaies: []instance.ExclusiveGateway{
								{
									Default: "Flow_194jx6p",
									Gateway: instance.Gateway{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Gateway_107rogi",
												},
												Name: "Gateway (Split Flow)",
											},
											Incoming: []string{"Flow_187opjq"},
											Outgoing: []string{"Flow_194jx6p", "Flow_1xhc3bf", "Flow_19m0ydj"},
										},
									},
								},
								{
									Gateway: instance.Gateway{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Gateway_140ec76",
												},
												Name: "Gateway (Merge Flows)",
											},
											Incoming: []string{"Flow_01ckxme", "Flow_1a24ve1"},
											Outgoing: []string{"Flow_1chzh1q"},
										},
									},
								},
							},
							SequenceFlows: []instance.SequenceFlow{
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0lacnkf",
										},
									},
									SourceRef: "StartEvent_1",
									TargetRef: "Activity_0ahdk3x",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_187opjq",
										},
									},
									SourceRef: "Activity_0ahdk3x",
									TargetRef: "Gateway_107rogi",
								},
								{
									FlowElement: instance.FlowElement{
										Name: "Default",
										BaseElement: instance.BaseElement{
											ID: "Flow_194jx6p",
										},
									},
									SourceRef: "Gateway_107rogi",
									TargetRef: "Activity_172ndxy",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1xhc3bf",
										},
									},
									SourceRef: "Gateway_107rogi",
									TargetRef: "Activity_1lz0l07",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1sd2kpj",
										},
									},
									SourceRef: "Activity_172ndxy",
									TargetRef: "Activity_1lz0l07",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_19m0ydj",
										},
									},
									SourceRef: "Gateway_107rogi",
									TargetRef: "Activity_171qefk",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_11m9fcl",
										},
									},
									SourceRef: "Activity_171qefk",
									TargetRef: "Activity_1lz0l07",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_01ckxme",
										},
										Name: "condition",
									},
									SourceRef: "Activity_171qefk",
									TargetRef: "Gateway_140ec76",
									ConditionExpression: instance.ConditionExpression{
										Type: instance.ExpressionTypeFormal,
										ExpressionSubstitution: &instance.FormalExpression{
											Expression: instance.Expression{
												BaseElementWithMixedContent: instance.BaseElementWithMixedContent{},
											},
										},
									},
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1a24ve1",
										},
									},
									SourceRef: "Activity_1lz0l07",
									TargetRef: "Gateway_140ec76",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0zwjy3h",
										},
										Name: "Condition",
									},
									SourceRef: "Activity_172ndxy",
									TargetRef: "Event_1wqqwdz",
									ConditionExpression: instance.ConditionExpression{
										Type: instance.ExpressionTypeFormal,
										ExpressionSubstitution: &instance.FormalExpression{
											Expression: instance.Expression{
												BaseElementWithMixedContent: instance.BaseElementWithMixedContent{},
											},
										},
									},
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1chzh1q",
										},
									},
									SourceRef: "Gateway_140ec76",
									TargetRef: "Event_1wqqwdz",
								},
							},
							EndEvents: []instance.EndEvent{{
								ThrowEvent: instance.ThrowEvent{
									Event: instance.Event{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Event_1wqqwdz",
												},
												Name: "End Event",
											},
											Incoming: []string{"Flow_0zwjy3h", "Flow_1chzh1q"},
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
