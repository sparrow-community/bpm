package bpmn

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sparrow-community/bpm/model/bpmn/instance"
)

func TestA_4_1_export(t *testing.T) {
	path := "./test/A.4.1-export.bpmn"
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
			ID:              "Definitions_18qzd9b",
			TargetNamespace: "http://bpmn.io/schema/bpmn",
			Exporter:        "Camunda Modeler",
			ExporterVersion: "5.4.2",
			Rootlemnts: instance.Rootlemnts{
				Collaborations: []instance.Collaboration{
					{
						RootElement: instance.RootElement{
							BaseElement: instance.BaseElement{
								ID: "Collaboration_0weizb5",
							},
						},
						Participants: []instance.Participant{
							{
								BaseElement: instance.BaseElement{
									ID: "Participant_08dproj",
								},
								Name:       "Pool 1",
								ProcessRef: "Process_0h42ymn",
							},
							{
								BaseElement: instance.BaseElement{
									ID: "Participant_1cs40k3",
								},
								Name:       "Pool 1",
								ProcessRef: "Process_18nmg48",
							},
						},
						MessageFlows: []instance.MessageFlow{
							{
								BaseElement: instance.BaseElement{
									ID: "Flow_0xbmmt1",
								},
								Name:      "Message Flow 2",
								SourceRef: "Activity_0xw7dvh",
								TargetRef: "Activity_1n3qt39",
							},
							{
								BaseElement: instance.BaseElement{
									ID: "Flow_0xqa8km",
								},
								Name:      "Message Flow 1",
								SourceRef: "Activity_1n97zgn",
								TargetRef: "Activity_0mmksew",
							},
						},
					},
				},
				Processes: []instance.Process{
					{
						CallableElement: instance.CallableElement{
							RootElement: instance.RootElement{
								BaseElement: instance.BaseElement{
									ID: "Process_0h42ymn",
								},
							},
						},
						IsExecutable: true,
						LaneSets: []instance.LaneSet{
							{
								BaseElement: instance.BaseElement{
									ID: "LaneSet_0k0qmv1",
								},
								Lanes: []instance.Lane{
									{
										BaseElement: instance.BaseElement{
											ID: "Lane_198nahk",
										},
										Name:         "Lane 1",
										FlowNodeRefs: []string{"StartEvent_1", "Activity_1n97zgn", "Activity_1n3qt39", "Event_1a4shvj"},
									},
								},
							},
						},
						FlowElements: instance.FlowElements{
							StartEvents: []instance.StartEvent{
								{
									CatchEvent: instance.CatchEvent{
										Event: instance.Event{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "StartEvent_1",
													},
													Name: "Start Event 1",
												},
												Outgoing: []string{"Flow_1hlsvck"},
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
													ID: "Activity_1n97zgn",
												},
												Name: "Task 1",
											},
											Incoming: []string{"Flow_1hlsvck"},
											Outgoing: []string{"Flow_1gcj53o"},
										},
									},
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Activity_1n3qt39",
												},
												Name: "Task 2",
											},
											Incoming: []string{"Flow_1gcj53o"},
											Outgoing: []string{"Flow_1dzj0ew"},
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
														ID: "Event_1a4shvj",
													},
													Name: "End Event 1",
												},
												Incoming: []string{"Flow_1dzj0ew"},
											},
										},
									},
								},
							},
							SequenceFlows: []instance.SequenceFlow{
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1hlsvck",
										},
									},
									SourceRef: "StartEvent_1",
									TargetRef: "Activity_1n97zgn",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1gcj53o",
										},
									},
									SourceRef: "Activity_1n97zgn",
									TargetRef: "Activity_1n3qt39",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1dzj0ew",
										},
									},
									SourceRef: "Activity_1n3qt39",
									TargetRef: "Event_1a4shvj",
								},
							},
						},
					},
					{
						CallableElement: instance.CallableElement{
							RootElement: instance.RootElement{
								BaseElement: instance.BaseElement{
									ID: "Process_18nmg48",
								},
							},
						},
						IsExecutable: false,
						LaneSets: []instance.LaneSet{
							{
								BaseElement: instance.BaseElement{
									ID: "LaneSet_11j5v67",
								},
								Lanes: []instance.Lane{
									{
										BaseElement: instance.BaseElement{
											ID: "Lane_1pmrgfe",
										},
										Name:         "Lane 4",
										FlowNodeRefs: []string{"Event_099731c", "Activity_0mmksew"},
									},
									{
										BaseElement: instance.BaseElement{
											ID: "Lane_0da6uqu",
										},
										Name:         "Lane 3",
										FlowNodeRefs: []string{"Activity_0xw7dvh", "Event_0zje79t", "Event_0wacqbg", "Activity_0hqqmqa", "Activity_1svfico"},
									},
								},
							},
						},
						FlowElements: instance.FlowElements{
							StartEvents: []instance.StartEvent{
								{
									CatchEvent: instance.CatchEvent{
										Event: instance.Event{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Event_099731c",
													},
													Name: "Start Event 2",
												},
												Outgoing: []string{"Flow_1dburd6"},
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
													ID: "Activity_0mmksew",
												},
												Name: "Task 3",
											},
											Incoming: []string{"Flow_1dburd6"},
											Outgoing: []string{"Flow_1jziro2", "Flow_0bzve1s"},
										},
									},
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Activity_0xw7dvh",
												},
												Name: "Task 5",
											},
											Incoming: []string{"Flow_04a6y1c"},
											Outgoing: []string{"Flow_1u6nbqs"},
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
														ID: "Event_0zje79t",
													},
													Name: "End Event 2",
												},
												Incoming: []string{"Flow_1u6nbqs"},
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
														ID: "Event_0wacqbg",
													},
													Name: "End Event 5",
												},
												Incoming: []string{"Flow_0yduxnw"},
											},
										},
									},
								},
							},
							SubProcess: []instance.SubProcess{
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Activity_0hqqmqa",
												},
												Name: "Expanded Sub-Process 2",
											},
											Incoming: []string{"Flow_1jziro2"},
											Outgoing: []string{"Flow_0yduxnw"},
										},
									},
									FlowElements: instance.FlowElements{
										StartEvents: []instance.StartEvent{
											{
												CatchEvent: instance.CatchEvent{
													Event: instance.Event{
														FlowNode: instance.FlowNode{
															FlowElement: instance.FlowElement{
																BaseElement: instance.BaseElement{
																	ID: "Event_1v26oan",
																},
																Name: "Start Event 4",
															},
															Outgoing: []string{"Flow_079eynm"},
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
																ID: "Activity_0sqf0ju",
															},
															Name: "Task 6",
														},
														Incoming: []string{"Flow_079eynm"},
														Outgoing: []string{"Flow_1palkux"},
													},
												},
											},
										},
										SequenceFlows: []instance.SequenceFlow{
											{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Flow_079eynm",
													},
												},
												SourceRef: "Event_1v26oan",
												TargetRef: "Activity_0sqf0ju",
											},
											{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Flow_1palkux",
													},
												},
												SourceRef: "Activity_0sqf0ju",
												TargetRef: "Event_1r1pop3",
											},
										},
										EndEvents: []instance.EndEvent{
											{
												ThrowEvent: instance.ThrowEvent{
													Event: instance.Event{
														FlowNode: instance.FlowNode{
															FlowElement: instance.FlowElement{
																BaseElement: instance.BaseElement{
																	ID: "Event_1r1pop3",
																},
																Name: "End Event 4",
															},
															Incoming: []string{"Flow_1palkux"},
														},
													},
												},
											},
										},
									},
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Activity_1svfico",
												},
												Name: "Expanded Sub-Process 1",
											},
											Incoming: []string{"Flow_0bzve1s"},
											Outgoing: []string{"Flow_04a6y1c"},
										},
									},
									FlowElements: instance.FlowElements{
										StartEvents: []instance.StartEvent{
											{
												CatchEvent: instance.CatchEvent{
													Event: instance.Event{
														FlowNode: instance.FlowNode{
															FlowElement: instance.FlowElement{
																BaseElement: instance.BaseElement{
																	ID: "Event_0xx2ic8",
																},
																Name: "Start Event 3",
															},
															Outgoing: []string{"Flow_0a4cvz3"},
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
																	ID: "Event_0gpo0qc",
																},
																Name: "End Event 3",
															},
															Incoming: []string{"Flow_031u0ap"},
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
																ID: "Activity_01cj9um",
															},
															Name: "Task 4",
														},
														Incoming: []string{"Flow_0a4cvz3"},
														Outgoing: []string{"Flow_031u0ap"},
													},
												},
											},
										},
										SequenceFlows: []instance.SequenceFlow{
											{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Flow_0a4cvz3",
													},
												},
												SourceRef: "Event_0xx2ic8",
												TargetRef: "Activity_01cj9um",
											},
											{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Flow_031u0ap",
													},
												},
												SourceRef: "Activity_01cj9um",
												TargetRef: "Event_0gpo0qc",
											},
										},
									},
								},
							},
							SequenceFlows: []instance.SequenceFlow{
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1dburd6",
										},
									},
									SourceRef: "Event_099731c",
									TargetRef: "Activity_0mmksew",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1jziro2",
										},
									},
									SourceRef: "Activity_0mmksew",
									TargetRef: "Activity_0hqqmqa",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0bzve1s",
										},
									},
									SourceRef: "Activity_0mmksew",
									TargetRef: "Activity_1svfico",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_04a6y1c",
										},
									},
									SourceRef: "Activity_1svfico",
									TargetRef: "Activity_0xw7dvh",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1u6nbqs",
										},
									},
									SourceRef: "Activity_0xw7dvh",
									TargetRef: "Event_0zje79t",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0yduxnw",
										},
									},
									SourceRef: "Activity_0hqqmqa",
									TargetRef: "Event_0wacqbg",
								},
							},
						},
					},
				},
			},
		},
	}

	if diff := cmp.Diff(expected, modelInstance); diff != "" {
		t.Errorf("BpmnModelInstanceFromFile() mismatch (-want +got):\n%s", diff)
	}
}
