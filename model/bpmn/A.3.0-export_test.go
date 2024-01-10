package bpmn

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sparrow-community/bpm/model/bpmn/instance"
)

func TestA_3_0_export(t *testing.T) {
	path := "./test/A.3.0-export.bpmn"
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
			ID:              "Definitions_06decf0",
			TargetNamespace: "http://bpmn.io/schema/bpmn",
			Exporter:        "Camunda Modeler",
			ExporterVersion: "5.2.0",
			RootElemnts: instance.RootElemnts{
				Processes: []instance.Process{
					{
						CallableElement: instance.CallableElement{
							RootElement: instance.RootElement{
								BaseElement: instance.BaseElement{
									ID: "Process_1qh1mjw",
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
											Outgoing: []string{"Flow_0su0msi"},
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
													ID: "Activity_164hjmz",
												},
												Name: "Task 1",
											},
											Incoming: []string{"Flow_0su0msi"},
											Outgoing: []string{"Flow_0iagd7s"},
										},
									},
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Activity_0mu4iin",
												},
												Name: "Task 2",
											},
											Incoming: []string{"Flow_0dukxt7"},
											Outgoing: []string{"Flow_0fg1rbu"},
										},
									},
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Activity_04ib995",
												},
												Name: "Task 4",
											},
											Incoming: []string{"Flow_0hvhb63"},
											Outgoing: []string{"Flow_12zl0ro"},
										},
									},
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Activity_12jlku6",
												},
												Name: "Task 3",
											},
											Incoming: []string{"Flow_0g8ko8k"},
											Outgoing: []string{"Flow_02rz41z"},
										},
									},
								},
							},
							SequenceFlows: []instance.SequenceFlow{
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0su0msi",
										},
									},
									SourceRef: "StartEvent_1",
									TargetRef: "Activity_164hjmz",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0iagd7s",
										},
									},
									SourceRef: "Activity_164hjmz",
									TargetRef: "Activity_1j4b29j",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0dukxt7",
										},
									},
									SourceRef: "Activity_1j4b29j",
									TargetRef: "Activity_0mu4iin",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0hvhb63",
										},
									},
									SourceRef: "Event_1bgdnfg",
									TargetRef: "Activity_04ib995",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_12zl0ro",
										},
									},
									SourceRef: "Activity_04ib995",
									TargetRef: "Event_1179od7",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0fg1rbu",
										},
									},
									SourceRef: "Activity_0mu4iin",
									TargetRef: "Event_1jn5oqm",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0g8ko8k",
										},
									},
									SourceRef: "Event_1uez1gc",
									TargetRef: "Activity_12jlku6",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_02rz41z",
										},
									},
									SourceRef: "Activity_12jlku6",
									TargetRef: "Event_1jn5oqm",
								},
							},
							SubProcess: []instance.SubProcess{
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Activity_1j4b29j",
												},
												Name: "Collapsed Sub-Process",
											},
											Incoming: []string{"Flow_0iagd7s"},
											Outgoing: []string{"Flow_0dukxt7"},
										},
									},
								},
							},
							BoundaryEvents: []instance.BoundaryEvent{
								{
									CatchEvent: instance.CatchEvent{
										Event: instance.Event{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Event_1uez1gc",
													},
													Name: "Boundary Intermediate Event Non-Interrupting Message",
												},
												Outgoing: []string{"Flow_0g8ko8k"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											MessageEventDefinitions: []instance.MessageEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "MessageEventDefinition_1eelr04",
															},
														},
													},
												},
											},
										},
									},
									CancelActivity: false,
									AttachedToRef:  "Activity_1j4b29j",
								},
								{
									CatchEvent: instance.CatchEvent{
										Event: instance.Event{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Event_1bgdnfg",
													},
													Name: "Boundary Intermediate Event Interrupting Escalation",
												},
												Outgoing: []string{"Flow_0hvhb63"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											EscalationEventDefinitions: []instance.EscalationEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "EscalationEventDefinition_1u8izyv",
															},
														},
													},
												},
											},
										},
									},
									AttachedToRef: "Activity_1j4b29j",
								},
							},
							EndEvents: []instance.EndEvent{
								{
									ThrowEvent: instance.ThrowEvent{
										Event: instance.Event{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Event_1179od7",
													},
													Name: "End Event 2",
												},
												Incoming: []string{"Flow_12zl0ro"},
											},
										},
									},
								},
								{
									ThrowEvent: instance.ThrowEvent{
										Event: instance.Event{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Event_1jn5oqm",
													},
													Name: "End Event 1",
												},
												Incoming: []string{"Flow_0fg1rbu", "Flow_02rz41z"},
											},
										},
									},
								},
							},
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
