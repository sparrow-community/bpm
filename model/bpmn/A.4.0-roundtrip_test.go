package bpmn

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sparrow-community/bpm/model/bpmn/instance"
)

func TestA_4_0_roundtrip(t *testing.T) {
	path := "./test/A.4.0-roundtrip.bpmn"
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
			ID:              "_1373649948794",
			TargetNamespace: "http://www.trisotech.com/definitions/_1373649948794",
			Name:            "A.4.0",
			Rootlemnts: instance.Rootlemnts{
				Collaborations: []instance.Collaboration{
					{
						RootElement: instance.RootElement{
							BaseElement: instance.BaseElement{
								ID: "C1373649949206",
							},
						},
						Participants: []instance.Participant{
							{
								BaseElement: instance.BaseElement{
									ID: "_046bff4f-cea3-4512-a6b1-30517fb29f2c",
								},
								Name:       "Pool",
								ProcessRef: "WFP-6-1",
							},
						},
						MessageFlows: []instance.MessageFlow{
							{
								BaseElement: instance.BaseElement{
									ID: "_b467921a-ef7b-44c5-bf78-fd624c400d17",
								},
								Name:       "Message Flow 1",
								SourceRef:  "_ab851300-b5de-4ad3-bbec-215553757fc8",
								TargetRef:  "_6fed62c8-8241-4a1d-ae67-266fda7dcead",
								MessageRef: "Message_1373649949207",
							},
							{
								BaseElement: instance.BaseElement{
									ID: "_c311cc87-677e-47a4-bdb1-8744c4ec3147",
								},
								Name:       "Message Flow 2",
								SourceRef:  "_1c347d0d-750b-4c09-980d-6877caae409b",
								TargetRef:  "_80d1f02b-f39c-45c2-b731-43df75d81779",
								MessageRef: "Message_1373649949208",
							},
						},
					},
				},
				Processes: []instance.Process{
					{
						CallableElement: instance.CallableElement{
							RootElement: instance.RootElement{
								BaseElement: instance.BaseElement{
									ID: "WFP-6-1",
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
														ID: "_c03f2b1f-32dc-41ef-b325-c9811a814fbe",
													},
													Name: "Start Event 1",
												},
												Outgoing: []string{"_44b1d373-57a1-4b8e-ba2e-3204c32519e5"},
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
													ID: "_ab851300-b5de-4ad3-bbec-215553757fc8",
												},
												Name: "Task 1",
											},
											Incoming: []string{"_44b1d373-57a1-4b8e-ba2e-3204c32519e5"},
											Outgoing: []string{"_6b7f2411-77f7-4152-be39-d8dbeb8bc460"},
										},
									},
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "_80d1f02b-f39c-45c2-b731-43df75d81779",
												},
												Name: "Task 2",
											},
											Incoming: []string{"_6b7f2411-77f7-4152-be39-d8dbeb8bc460"},
											Outgoing: []string{"_f35e3b07-7b1f-433d-9595-7fdea8996954"},
										},
									},
								},
							},
							SequenceFlows: []instance.SequenceFlow{
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_44b1d373-57a1-4b8e-ba2e-3204c32519e5",
										},
									},
									SourceRef: "_c03f2b1f-32dc-41ef-b325-c9811a814fbe",
									TargetRef: "_ab851300-b5de-4ad3-bbec-215553757fc8",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_f35e3b07-7b1f-433d-9595-7fdea8996954",
										},
									},
									SourceRef: "_80d1f02b-f39c-45c2-b731-43df75d81779",
									TargetRef: "_6e79c19f-749d-48c4-8271-d9ca028354fa",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_6b7f2411-77f7-4152-be39-d8dbeb8bc460",
										},
									},
									SourceRef: "_ab851300-b5de-4ad3-bbec-215553757fc8",
									TargetRef: "_80d1f02b-f39c-45c2-b731-43df75d81779",
								},
							},
							EndEvents: []instance.EndEvent{
								{
									ThrowEvent: instance.ThrowEvent{
										Event: instance.Event{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "_6e79c19f-749d-48c4-8271-d9ca028354fa",
													},
													Name: "End Event 1",
												},
												Incoming: []string{"_f35e3b07-7b1f-433d-9595-7fdea8996954"},
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
									ID: "WFP-6-2",
								},
							},
						},
						IsExecutable: false,
						LaneSets: []instance.LaneSet{
							{
								BaseElement: instance.BaseElement{
									ID: "ls_1373649949139",
								},
								Lanes: []instance.Lane{
									{
										BaseElement: instance.BaseElement{
											ID: "_17bebb0f-f31e-475a-b1b1-76fcc2da172b",
										},
										Name:         "Lane 1",
										FlowNodeRefs: []string{"_1c347d0d-750b-4c09-980d-6877caae409b", "_65d1bebf-e613-4317-acb2-b12b69fc67ff", "_ee35fa2c-dfea-40cf-a469-845b765a7b50", "_7c434d45-d319-457b-9fd6-853c218bc3f1", "_6fed62c8-8241-4a1d-ae67-266fda7dcead"},
									},
									{
										BaseElement: instance.BaseElement{
											ID: "_cc1845d0-ec34-44d3-8ba5-4981040d8dfe",
										},
										Name:         "Lane 2",
										FlowNodeRefs: []string{"_8e6cecb7-b247-4c43-a6b6-532fb6a89753", "_f52b6ad0-4dcc-4053-b696-b924dda01db5"},
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
														ID: "_65d1bebf-e613-4317-acb2-b12b69fc67ff",
													},
													Name: "Start Event 2",
												},
												Outgoing: []string{"_486d13e4-86ef-49b2-bba9-f03435494f0e"},
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
													ID: "_1c347d0d-750b-4c09-980d-6877caae409b",
												},
												Name: "Task 5",
											},
											Incoming: []string{"_ebbaed22-6fcb-4af7-8b7a-7ebc9cc7f150"},
											Outgoing: []string{"_0020ed6a-6dde-499f-9fda-36c8bde20ec6"},
										},
									},
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "_6fed62c8-8241-4a1d-ae67-266fda7dcead",
												},
												Name: "Task 3",
											},
											Incoming: []string{"_486d13e4-86ef-49b2-bba9-f03435494f0e"},
											Outgoing: []string{"_1873ed55-ba18-433f-8d1a-eb84d18da049", "_08d345a0-2bc3-4988-bfb2-7c2576839505"},
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
													ID: "_ee35fa2c-dfea-40cf-a469-845b765a7b50",
												},
												Name: "Expanded Sub-Process 1",
											},
											Incoming: []string{"_1873ed55-ba18-433f-8d1a-eb84d18da049"},
											Outgoing: []string{"_ebbaed22-6fcb-4af7-8b7a-7ebc9cc7f150"},
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
																	ID: "_1ffaa550-3225-4c6a-a391-3aaf224723af",
																},
																Name: "Start Event 3",
															},
															Outgoing: []string{"_3580d9ba-4f79-48c1-96da-090bd5e5172d"},
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
																ID: "_09532ad3-e571-4214-b580-7bebf4bb68b1",
															},
															Name: "Task 4",
														},
														Incoming: []string{"_3580d9ba-4f79-48c1-96da-090bd5e5172d"},
														Outgoing: []string{"_f5e6dc98-8c79-4e6b-aef9-c3f3bdc09136"},
													},
												},
											},
										},
										SequenceFlows: []instance.SequenceFlow{
											{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "_3580d9ba-4f79-48c1-96da-090bd5e5172d",
													},
												},
												SourceRef: "_1ffaa550-3225-4c6a-a391-3aaf224723af",
												TargetRef: "_09532ad3-e571-4214-b580-7bebf4bb68b1",
											},
											{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "_f5e6dc98-8c79-4e6b-aef9-c3f3bdc09136",
													},
												},
												SourceRef: "_09532ad3-e571-4214-b580-7bebf4bb68b1",
												TargetRef: "_3e5ac6ed-88d6-4f82-a647-6b253b80b004",
											},
										},
										EndEvents: []instance.EndEvent{
											{
												ThrowEvent: instance.ThrowEvent{
													Event: instance.Event{
														FlowNode: instance.FlowNode{
															FlowElement: instance.FlowElement{
																BaseElement: instance.BaseElement{
																	ID: "_3e5ac6ed-88d6-4f82-a647-6b253b80b004",
																},
																Name: "End Event 3",
															},
															Incoming: []string{"_f5e6dc98-8c79-4e6b-aef9-c3f3bdc09136"},
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
													ID: "_f52b6ad0-4dcc-4053-b696-b924dda01db5",
												},
												Name: "Expanded Sub-Process 2",
											},
											Incoming: []string{"_08d345a0-2bc3-4988-bfb2-7c2576839505"},
											Outgoing: []string{"_1f3792a7-da0f-4621-8c10-a04b67e33f5b"},
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
																	ID: "_47bef337-7915-459d-a9cd-e9c87c98f8fa",
																},
																Name: "Start Event 4",
															},
															Outgoing: []string{"_807d5f9c-e4e5-49fc-b44a-da9743a94556"},
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
																ID: "_15f8f2a4-5e55-4159-b349-403ac4cbdefb",
															},
															Name: "Task 6",
														},
														Incoming: []string{"_807d5f9c-e4e5-49fc-b44a-da9743a94556"},
														Outgoing: []string{"_237d3b22-6012-49f7-90a3-cd259426caf9"},
													},
												},
											},
										},
										SequenceFlows: []instance.SequenceFlow{
											{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "_807d5f9c-e4e5-49fc-b44a-da9743a94556",
													},
												},
												SourceRef: "_47bef337-7915-459d-a9cd-e9c87c98f8fa",
												TargetRef: "_15f8f2a4-5e55-4159-b349-403ac4cbdefb",
											},
											{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "_237d3b22-6012-49f7-90a3-cd259426caf9",
													},
												},
												SourceRef: "_15f8f2a4-5e55-4159-b349-403ac4cbdefb",
												TargetRef: "_bb8b7952-0991-4b7c-a851-97327832d7b8",
											},
										},
										EndEvents: []instance.EndEvent{
											{
												ThrowEvent: instance.ThrowEvent{
													Event: instance.Event{
														FlowNode: instance.FlowNode{
															FlowElement: instance.FlowElement{
																BaseElement: instance.BaseElement{
																	ID: "_bb8b7952-0991-4b7c-a851-97327832d7b8",
																},
																Name: "End Event 4",
															},
															Incoming: []string{"_237d3b22-6012-49f7-90a3-cd259426caf9"},
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
											ID: "_1873ed55-ba18-433f-8d1a-eb84d18da049",
										},
									},
									SourceRef: "_6fed62c8-8241-4a1d-ae67-266fda7dcead",
									TargetRef: "_ee35fa2c-dfea-40cf-a469-845b765a7b50",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_ebbaed22-6fcb-4af7-8b7a-7ebc9cc7f150",
										},
									},
									SourceRef: "_ee35fa2c-dfea-40cf-a469-845b765a7b50",
									TargetRef: "_1c347d0d-750b-4c09-980d-6877caae409b",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_0020ed6a-6dde-499f-9fda-36c8bde20ec6",
										},
									},
									SourceRef: "_1c347d0d-750b-4c09-980d-6877caae409b",
									TargetRef: "_7c434d45-d319-457b-9fd6-853c218bc3f1",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_08d345a0-2bc3-4988-bfb2-7c2576839505",
										},
									},
									SourceRef: "_6fed62c8-8241-4a1d-ae67-266fda7dcead",
									TargetRef: "_f52b6ad0-4dcc-4053-b696-b924dda01db5",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_1f3792a7-da0f-4621-8c10-a04b67e33f5b",
										},
									},
									SourceRef: "_f52b6ad0-4dcc-4053-b696-b924dda01db5",
									TargetRef: "_8e6cecb7-b247-4c43-a6b6-532fb6a89753",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_486d13e4-86ef-49b2-bba9-f03435494f0e",
										},
									},
									SourceRef: "_65d1bebf-e613-4317-acb2-b12b69fc67ff",
									TargetRef: "_6fed62c8-8241-4a1d-ae67-266fda7dcead",
								},
							},
							EndEvents: []instance.EndEvent{
								{
									ThrowEvent: instance.ThrowEvent{
										Event: instance.Event{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "_7c434d45-d319-457b-9fd6-853c218bc3f1",
													},
													Name: "End Event 2",
												},
												Incoming: []string{"_0020ed6a-6dde-499f-9fda-36c8bde20ec6"},
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
														ID: "_8e6cecb7-b247-4c43-a6b6-532fb6a89753",
													},
													Name: "End Event 5",
												},
												Incoming: []string{"_1f3792a7-da0f-4621-8c10-a04b67e33f5b"},
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
