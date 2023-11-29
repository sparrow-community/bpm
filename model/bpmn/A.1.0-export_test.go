package bpmn

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sparrow-community/bpm/model/bpmn/instance"
)

func TestA_1_0_export(t *testing.T) {
	// create test use ./test/A.1.0-export.bpmn
	path := "./test/A.1.0-export.bpmn"
	modelInstance, err := BpmnModelInstanceFromFile(path)
	if err != nil {
		t.Fatalf("BpmnModelInstanceFromFile error %s: %s", path, err)
	}
	if modelInstance == nil {
		t.Fatalf("modelInstance is nil, %s", path)
	}

	excepted := &BpmnModelInstance{
		Definitions: &instance.Definitions{
			XMLName: xml.Name{
				Space: "http://www.omg.org/spec/BPMN/20100524/MODEL",
				Local: "definitions",
			},
			ID:              "sid-38422fae-e03e-43a3-bef4-bd33b32041b2",
			Name:            "",
			TargetNamespace: "http://bpmn.io/bpmn",
			Exporter:        "Camunda Modeler",
			ExporterVersion: "5.0.0",
			Rootlemnts: instance.Rootlemnts{
				Processes: []instance.Process{
					{
						CallableElement: instance.CallableElement{
							RootElement: instance.RootElement{
								BaseElement: instance.BaseElement{
									ID: "Process_1",
								},
							},
						},
						IsExecutable: false,
						FlowElements: instance.FlowElements{
							StartEvents: []instance.StartEvent{
								{
									CatchEvent: instance.CatchEvent{
										Event: instance.Event{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Event_1pmxsnn",
													},
													Name: "Start Event",
												},
												Outgoing: []string{"Flow_0iyzbi9"},
											},
										},
									},
								},
							},
							EndEvents: []instance.EndEvent{
								{
									ThrowEvent: instance.ThrowEvent{
										Event: instance.Event{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Event_0ki4ik8",
													},
													Name: "End Event",
												},
												Incoming: []string{"Flow_01pjh7d"},
											},
										},
									},
								},
							},
							Tasks: []instance.Task{
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Activity_10i3hk7",
												},
												Name: "Task 1",
											},
											Incoming: []string{"Flow_0iyzbi9"},
											Outgoing: []string{"Flow_0ll5ug1"},
										},
									},
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Activity_1eb0bmc",
												},
												Name: "Task 2",
											},
											Incoming: []string{"Flow_0ll5ug1"},
											Outgoing: []string{"Flow_0ec6s1g"},
										},
									},
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Activity_1m3q7qr",
												},
												Name: "Task 3",
											},
											Incoming: []string{"Flow_0ec6s1g"},
											Outgoing: []string{"Flow_01pjh7d"},
										},
									},
								},
							},
							SequenceFlows: []instance.SequenceFlow{
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0iyzbi9",
										},
									},
									SourceRef: "Event_1pmxsnn",
									TargetRef: "Activity_10i3hk7",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0ll5ug1",
										},
									},
									SourceRef: "Activity_10i3hk7",
									TargetRef: "Activity_1eb0bmc",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0ec6s1g",
										},
									},
									SourceRef: "Activity_1eb0bmc",
									TargetRef: "Activity_1m3q7qr",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_01pjh7d",
										},
									},
									SourceRef: "Activity_1m3q7qr",
									TargetRef: "Event_0ki4ik8",
								},
							},
						},
					},
				},
			},
		},
	}

	if diff := cmp.Diff(excepted, modelInstance); diff != "" {
		t.Errorf("Mismatch in %s (-want +got):\n%s", path, diff)
	}
}
