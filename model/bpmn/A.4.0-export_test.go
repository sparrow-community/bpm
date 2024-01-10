package bpmn

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sparrow-community/bpm/model/bpmn/instance"
)

func TestA_4_0_export(t *testing.T) {
	path := "./test/A.4.0-export.bpmn"
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
			ID:              "Definitions_14ccggy",
			TargetNamespace: "http://bpmn.io/schema/bpmn",
			Exporter:        "Camunda Modeler",
			ExporterVersion: "5.2.0",
			RootElemnts: instance.RootElemnts{
				Collaborations: []instance.Collaboration{
					{
						RootElement: instance.RootElement{
							BaseElement: instance.BaseElement{
								ID: "Collaboration_1eeqkua",
							},
						},
						Participants: []instance.Participant{
							{
								BaseElement: instance.BaseElement{
									ID: "PoolParticipant",
								},
								Name:       "Pool",
								ProcessRef: "Process_0elb8rq",
							},
							{
								BaseElement: instance.BaseElement{
									ID: "Participant_1b8727b",
								},
								ProcessRef: "Process_0wqyt7t",
							},
						},
						MessageFlows: []instance.MessageFlow{
							{
								BaseElement: instance.BaseElement{
									ID: "MessageFlow1MessageFlow",
								},
								Name:      "Message Flow 1",
								SourceRef: "Task1Task",
								TargetRef: "Task3Task",
							},
							{
								BaseElement: instance.BaseElement{
									ID: "MessageFlow2MessageFlow",
								},
								Name:      "Message Flow 2",
								SourceRef: "Task5Task",
								TargetRef: "Task2Task",
							},
						},
					},
				},
				Processes: []instance.Process{
					{
						CallableElement: instance.CallableElement{
							RootElement: instance.RootElement{
								BaseElement: instance.BaseElement{
									ID: "Process_0elb8rq",
								},
							},
						},
						IsExecutable: true,
						FlowElements: instance.FlowElements{
							StartEvents: []instance.StartEvent{
								{
									CatchEvent: instance.CatchEvent{
										Event: instance.Event{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "StartEvent1StartEvent",
													},
													Name: "Start Event 1",
												},
												Outgoing: []string{"Flow_193i5eg"},
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
													ID: "Task1Task",
												},
												Name: "Task 1",
											},
											Incoming: []string{"Flow_193i5eg"},
											Outgoing: []string{"Flow_01btz5o"},
										},
									},
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Task2Task",
												},
												Name: "Task 2",
											},
											Incoming: []string{"Flow_01btz5o"},
											Outgoing: []string{"Flow_14uttgb"},
										},
									},
								},
							},
							SequenceFlows: []instance.SequenceFlow{
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_193i5eg",
										},
									},
									SourceRef: "StartEvent1StartEvent",
									TargetRef: "Task1Task",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_01btz5o",
										},
									},
									SourceRef: "Task1Task",
									TargetRef: "Task2Task",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_14uttgb",
										},
									},
									SourceRef: "Task2Task",
									TargetRef: "EndEvent1EndEvent",
								},
							},
							EndEvents: []instance.EndEvent{
								{
									ThrowEvent: instance.ThrowEvent{
										Event: instance.Event{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "EndEvent1EndEvent",
													},
													Name: "End Event 1",
												},
												Incoming: []string{"Flow_14uttgb"},
											},
										},
									},
								},
							},
						},
					},
					{
						CallableElement: instance.CallableElement{
							RootElement: instance.RootElement{
								BaseElement: instance.BaseElement{
									ID: "Process_0wqyt7t",
								},
							},
						},
						IsExecutable: false,
						LaneSets: []instance.LaneSet{
							{
								BaseElement: instance.BaseElement{
									ID: "LaneSet_1lf7xw1",
								},
								Lanes: []instance.Lane{
									{
										BaseElement: instance.BaseElement{
											ID: "Lane2Lane",
										},
										Name:         "Lane 2",
										FlowNodeRefs: []string{"ExpandedSubProcess2SubProcess", "EndEvent5EndEvent"},
									},
									{
										BaseElement: instance.BaseElement{
											ID: "Lane1Lane",
										},
										Name:         "Lane 1",
										FlowNodeRefs: []string{"StartEvent2StartEvent", "Task3Task", "Task5Task", "EndEvent2EndEvent", "ExpandedSubProcess1SubProcess"},
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
														ID: "StartEvent2StartEvent",
													},
													Name: "Start Event 2",
												},
												Outgoing: []string{"Flow_0g54c4e"},
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
													ID: "Task3Task",
												},
												Name: "Task 3",
											},
											Incoming: []string{"Flow_0g54c4e"},
											Outgoing: []string{"Flow_0gi2n31", "Flow_14r4yf8"},
										},
									},
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Task5Task",
												},
												Name: "Task 5",
											},
											Incoming: []string{"Flow_0fjya97"},
											Outgoing: []string{"Flow_0tx0rfn"},
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
													ID: "ExpandedSubProcess2SubProcess",
												},
												Name: "Expanded Sub-Process 2",
											},
											Incoming: []string{"Flow_0gi2n31"},
											Outgoing: []string{"Flow_04fun2b"},
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
																	ID: "StartEvent4StartEvent",
																},
																Name: "Start Event 4",
															},
															Outgoing: []string{"Flow_1ups9iv"},
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
																ID: "Task6Task",
															},
															Name: "Task 6",
														},
														Incoming: []string{"Flow_1ups9iv"},
														Outgoing: []string{"Flow_0w343e8"},
													},
												},
											},
										},
										SequenceFlows: []instance.SequenceFlow{
											{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Flow_1ups9iv",
													},
												},
												SourceRef: "StartEvent4StartEvent",
												TargetRef: "Task6Task",
											},
											{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Flow_0w343e8",
													},
												},
												SourceRef: "Task6Task",
												TargetRef: "EndEvent4EndEvent",
											},
										},
										EndEvents: []instance.EndEvent{
											{
												ThrowEvent: instance.ThrowEvent{
													Event: instance.Event{
														FlowNode: instance.FlowNode{
															FlowElement: instance.FlowElement{
																BaseElement: instance.BaseElement{
																	ID: "EndEvent4EndEvent",
																},
																Name: "End Event 4",
															},
															Incoming: []string{"Flow_0w343e8"},
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
													ID: "ExpandedSubProcess1SubProcess",
												},
												Name: "Expanded Sub-Process 1",
											},
											Incoming: []string{"Flow_14r4yf8"},
											Outgoing: []string{"Flow_0fjya97"},
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
																	ID: "StartEvent3StartEvent",
																},
																Name: "Start Event 3",
															},
															Outgoing: []string{"Flow_0iuanwa"},
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
																ID: "Task4Task",
															},
															Name: "Task 4",
														},
														Incoming: []string{"Flow_0iuanwa"},
														Outgoing: []string{"Flow_1y4ym2b"},
													},
												},
											},
										},
										SequenceFlows: []instance.SequenceFlow{
											{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Flow_0iuanwa",
													},
												},
												SourceRef: "StartEvent3StartEvent",
												TargetRef: "Task4Task",
											},
											{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Flow_1y4ym2b",
													},
												},
												SourceRef: "Task4Task",
												TargetRef: "EndEvent3EndEvent",
											},
										},
										EndEvents: []instance.EndEvent{
											{
												ThrowEvent: instance.ThrowEvent{
													Event: instance.Event{
														FlowNode: instance.FlowNode{
															FlowElement: instance.FlowElement{
																BaseElement: instance.BaseElement{
																	ID: "EndEvent3EndEvent",
																},
																Name: "End Event 3",
															},
															Incoming: []string{"Flow_1y4ym2b"},
														},
													},
												},
											},
										},
									},
								},
							},
							SequenceFlows: []instance.SequenceFlow{
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_14r4yf8",
										},
									},
									SourceRef: "Task3Task",
									TargetRef: "ExpandedSubProcess1SubProcess",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0gi2n31",
										},
									},
									SourceRef: "Task3Task",
									TargetRef: "ExpandedSubProcess2SubProcess",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_04fun2b",
										},
									},
									SourceRef: "ExpandedSubProcess2SubProcess",
									TargetRef: "EndEvent5EndEvent",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0tx0rfn",
										},
									},
									SourceRef: "Task5Task",
									TargetRef: "EndEvent2EndEvent",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0fjya97",
										},
									},
									SourceRef: "ExpandedSubProcess1SubProcess",
									TargetRef: "Task5Task",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0g54c4e",
										},
									},
									SourceRef: "StartEvent2StartEvent",
									TargetRef: "Task3Task",
								},
							},
							EndEvents: []instance.EndEvent{
								{
									ThrowEvent: instance.ThrowEvent{
										Event: instance.Event{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "EndEvent2EndEvent",
													},
													Name: "End Event 2",
												},
												Incoming: []string{"Flow_0tx0rfn"},
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
														ID: "EndEvent5EndEvent",
													},
													Name: "End Event 5",
												},
												Incoming: []string{"Flow_04fun2b"},
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
		t.Errorf("BpmnModelInstanceFromFile() mismatch (-want +got):\n%s", diff)
	}
}
