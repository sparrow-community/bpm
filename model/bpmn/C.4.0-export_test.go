package bpmn

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sparrow-community/bpm/model/bpmn/instance"
)

func TestC_4_0_export(t *testing.T) {
	// create test use ./test/C.4.0-export.bpmn
	path := "./test/C.4.0-export.bpmn"
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
			ID:              "Definitions_008og7z",
			TargetNamespace: "http://bpmn.io/schema/bpmn",
			Exporter:        "Camunda Modeler",
			ExporterVersion: "5.2.0",
			RootElemnts: instance.RootElemnts{
				Collaborations: []instance.Collaboration{
					{
						RootElement: instance.RootElement{
							BaseElement: instance.BaseElement{
								ID: "Collaboration_08tzwy8",
							},
						},
						Participants: []instance.Participant{
							{
								BaseElement: instance.BaseElement{
									ID: "Participant_0d33sio",
								},
								Name:       "Money Bank",
								ProcessRef: "Process_07wr932",
							},
						},
					},
				},
				Processes: []instance.Process{
					{
						CallableElement: instance.CallableElement{
							RootElement: instance.RootElement{
								BaseElement: instance.BaseElement{
									ID: "Process_07wr932",
								},
							},
						},
						IsExecutable: false,
						LaneSets: []instance.LaneSet{
							{
								BaseElement: instance.BaseElement{
									ID: "LaneSet_0ujafgq",
								},
								Lanes: []instance.Lane{
									{
										BaseElement: instance.BaseElement{
											ID: "Lane_0w7h48m",
										},
										Name: "HR Department",
										FlowNodeRefs: []string{
											"Event_0aqeoee",
											"Activity_1vif48i",
											"Gateway_1nitinv",
											"Activity_0gis8ag",
											"Activity_1iegcae",
											"Gateway_13pi9cy",
											"Activity_1fydq08",
											"Activity_1ptc2l6",
											"Activity_0fxxj1x",
											"Gateway_0tmub1o",
											"Activity_0pb64l8",
										},
										ChildLaneSet: instance.LaneSet{
											BaseElement: instance.BaseElement{
												ID: "LaneSet_0gxwbwu",
											},
										},
									},
									{
										BaseElement: instance.BaseElement{
											ID: "Lane_1x43fzg",
										},
										Name: "Responsible Department",
										FlowNodeRefs: []string{
											"Activity_1567dil",
											"Event_1uj58zc",
											"Activity_1uixx0o",
											"Activity_05503d3",
											"Event_0qacxgu",
											"Event_0axou2i",
											"Gateway_1s7by1w",
											"Gateway_0vtzq53",
											"Event_0bjpe63",
											"Activity_0nzk36a",
											"Activity_0yq2h1t",
											"Event_1dg5pp3",
										},
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
														ID: "Event_0aqeoee",
													},
													Name: "Candidate accepted offer",
												},
												Outgoing: []string{"Flow_161isb1"},
											},
										},
									},
								},
							},
							UserTasks: []instance.UserTask{
								{
									Task: instance.Task{
										Activity: instance.Activity{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Activity_1vif48i",
													},
													Name: "Send candidate Contract",
												},
												Incoming: []string{"Flow_161isb1", "Flow_0d72bsp"},
												Outgoing: []string{"Flow_1kkgeei"},
											},
										},
									},
								},
								{
									Task: instance.Task{
										Activity: instance.Activity{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Activity_0gis8ag",
													},
													Name: "Review terms of contract",
												},
												Incoming: []string{"Flow_1k91huu"},
												Outgoing: []string{"Flow_0d72bsp"},
											},
										},
									},
								},
								{
									Task: instance.Task{
										Activity: instance.Activity{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Activity_1iegcae",
													},
													Name: "Get signature on contract and notify responsible department",
												},
												Incoming: []string{"Flow_02hcma9"},
												Outgoing: []string{"Flow_1xymj9u"},
											},
											DataOutputAssociations: []instance.DataOutputAssociation{
												{
													DataAssociation: instance.DataAssociation{
														BaseElement: instance.BaseElement{
															ID: "DataOutputAssociation_0q6chhu",
														},
														TargetRef: "DataStoreReference_1el2zpd",
													},
												},
											},
										},
									},
								},
								{
									Task: instance.Task{
										Activity: instance.Activity{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Activity_1fydq08",
													},
													Name: "Inform employee of company policies",
												},
												Incoming: []string{"Flow_0jlumov"},
												Outgoing: []string{"Flow_0gs48ue"},
											},
										},
									},
								},
								{
									Task: instance.Task{
										Activity: instance.Activity{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Activity_1ptc2l6",
													},
													Name: "Introduce employee to company Mission, Vision and Values",
												},
												Incoming: []string{"Flow_0gs48ue"},
												Outgoing: []string{"Flow_0f14mpt"},
											},
										},
									},
								},
								{
									Task: instance.Task{
										Activity: instance.Activity{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Activity_1567dil",
													},
													Name: "Request preparations for a new employee",
												},
												Incoming: []string{"Flow_0u0mecw"},
												Outgoing: []string{"Flow_0qufezv"},
											},
										},
									},
								},
								{
									Task: instance.Task{
										Activity: instance.Activity{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Activity_0fxxj1x",
													},
													Name: "Register for medical insurance",
												},
												Incoming: []string{"Flow_13ztu1l"},
												Outgoing: []string{"Flow_0jag15j"},
											},
										},
									},
								},
								{
									Task: instance.Task{
										Activity: instance.Activity{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Activity_1uixx0o",
													},
													Name: "Introduce new employee to the team",
												},
												Incoming: []string{"Flow_0pzvylb"},
												Outgoing: []string{"Flow_1fuxj6b"},
											},
										},
									},
								},
								{
									Task: instance.Task{
										Activity: instance.Activity{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Activity_05503d3",
													},
													Name: "Perform training for position",
												},
												Incoming: []string{"Flow_1fuxj6b"},
												Outgoing: []string{"Flow_1ikbbnd"},
											},
										},
									},
								},
								{
									Task: instance.Task{
										Activity: instance.Activity{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Activity_0nzk36a",
													},
													Name: "Compile welcome package",
												},
												Incoming: []string{"Flow_1367fco"},
												Outgoing: []string{"Flow_0n1kpzs"},
											},
										},
									},
								},
								{
									Task: instance.Task{
										Activity: instance.Activity{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Activity_0pb64l8",
													},
													Name: "Perform training for time reports sick leave and holidays",
												},
												Incoming: []string{"Flow_0f14mpt"},
												Outgoing: []string{"Flow_13ztu1l"},
											},
										},
									},
								},
								{
									Task: instance.Task{
										Activity: instance.Activity{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Activity_0yq2h1t",
													},
													Name: "Give employee welcome package",
												},
												Incoming: []string{"Flow_0n1kpzs"},
												Outgoing: []string{"Flow_1acnq20"},
											},
										},
									},
								},
							},
							ExclusiveGatewaies: []instance.ExclusiveGateway{
								{
									Gateway: instance.Gateway{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Gateway_1nitinv",
												},
												Name: "Contract terms accepted ?",
											},
											Incoming: []string{"Flow_1kkgeei"},
											Outgoing: []string{"Flow_1k91huu", "Flow_02hcma9"},
										},
									},
								},
							},
							ParallelGatewaies: []instance.ParallelGateway{
								{
									Gateway: instance.Gateway{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Gateway_13pi9cy",
												},
												Name: "Non exclusive Gateway",
											},
											Incoming: []string{"Flow_1xymj9u"},
											Outgoing: []string{"Flow_0jlumov", "Flow_0u0mecw"},
										},
									},
								},
								{
									Gateway: instance.Gateway{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Gateway_0tmub1o",
												},
												Name: "Non exclusive Gateway",
											},
											Incoming: []string{"Flow_0jag15j", "Flow_1jp11xh"},
											Outgoing: []string{"Flow_0pzvylb"},
										},
									},
								},
								{
									Gateway: instance.Gateway{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Gateway_1s7by1w",
												},
											},
											Incoming: []string{"Flow_01sh6l9", "Flow_0utj2cp", "Flow_1dqygu2"},
											Outgoing: []string{"Flow_1367fco"},
										},
									},
								},
								{
									Gateway: instance.Gateway{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Gateway_0vtzq53",
												},
												Name: "Non exclusive Gateway",
											},
											Incoming: []string{"Flow_1ikbbnd"},
											Outgoing: []string{"Flow_1ysfo4f", "Flow_07uts0i", "Flow_11nkk44"},
										},
									},
								},
							},
							DataStoreReferences: []instance.DataStoreReference{
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "DataStoreReference_1el2zpd",
										},
										Name: "Employee Details",
									},
								},
							},
							IntermediateThrowEvents: []instance.IntermediateThrowEvent{
								{
									ThrowEvent: instance.ThrowEvent{
										Event: instance.Event{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Event_1uj58zc",
													},
													Name: "New employee in department X",
												},
												Incoming: []string{"Flow_0qufezv"},
												Outgoing: []string{"Flow_1jp11xh"},
											},
											Properties: []instance.Property{
												{
													BaseElement: instance.BaseElement{
														ID: "Property_0bxqyv2",
													},
													Name: "__targetRef_placeholder",
												},
											},
										},
										DataInputAssociations: []instance.DataInputAssociation{
											{
												DataAssociation: instance.DataAssociation{
													BaseElement: instance.BaseElement{
														ID: "DataInputAssociation_0z6q52m",
													},
													SourceRef: "DataStoreReference_1el2zpd",
													TargetRef: "Property_0bxqyv2",
												},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											SignalEventDefinitions: []instance.SignalEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "SignalEventDefinition_031ap9q",
															},
														},
													},
												},
											},
										},
									},
								},
							},
							IntermediateCatchEvents: []instance.IntermediateCatchEvent{
								{
									CatchEvent: instance.CatchEvent{
										Event: instance.Event{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Event_0qacxgu",
													},
													Name: "Input from IT ready",
												},
												Incoming: []string{"Flow_1ysfo4f"},
												Outgoing: []string{"Flow_0utj2cp"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											MessageEventDefinitions: []instance.MessageEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "MessageEventDefinition_0mcgs4q",
															},
														},
													},
												},
											},
										},
									},
								},
								{
									CatchEvent: instance.CatchEvent{
										Event: instance.Event{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Event_0axou2i",
													},
													Name: "Input from Payroll ready",
												},
												Incoming: []string{"Flow_07uts0i"},
												Outgoing: []string{"Flow_01sh6l9"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											MessageEventDefinitions: []instance.MessageEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "MessageEventDefinition_11piq11",
															},
														},
													},
												},
											},
										},
									},
								},
								{
									CatchEvent: instance.CatchEvent{
										Event: instance.Event{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Event_0bjpe63",
													},
													Name: "Input from Facilities ready",
												},
												Incoming: []string{"Flow_11nkk44"},
												Outgoing: []string{"Flow_1dqygu2"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											MessageEventDefinitions: []instance.MessageEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "MessageEventDefinition_0txndb3",
															},
														},
													},
												},
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
														ID: "Event_1dg5pp3",
													},
													Name: "End Event",
												},
												Incoming: []string{"Flow_1acnq20"},
											},
										},
									},
								},
							},
							SequenceFlows: []instance.SequenceFlow{
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_161isb1",
										},
									},
									SourceRef: "Event_0aqeoee",
									TargetRef: "Activity_1vif48i",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1kkgeei",
										},
									},
									SourceRef: "Activity_1vif48i",
									TargetRef: "Gateway_1nitinv",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1k91huu",
										},
										Name: "No",
									},
									SourceRef: "Gateway_1nitinv",
									TargetRef: "Activity_0gis8ag",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0d72bsp",
										},
										Name: "",
									},
									SourceRef: "Activity_0gis8ag",
									TargetRef: "Activity_1vif48i",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_02hcma9",
										},
										Name: "Yes",
									},
									SourceRef: "Gateway_1nitinv",
									TargetRef: "Activity_1iegcae",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1xymj9u",
										},
									},
									SourceRef: "Activity_1iegcae",
									TargetRef: "Gateway_13pi9cy",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0jlumov",
										},
									},
									SourceRef: "Gateway_13pi9cy",
									TargetRef: "Activity_1fydq08",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0gs48ue",
										},
									},
									SourceRef: "Activity_1fydq08",
									TargetRef: "Activity_1ptc2l6",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0f14mpt",
										},
									},
									SourceRef: "Activity_1ptc2l6",
									TargetRef: "Activity_0pb64l8",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0jag15j",
										},
									},
									SourceRef: "Activity_0fxxj1x",
									TargetRef: "Gateway_0tmub1o",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0u0mecw",
										},
									},
									SourceRef: "Gateway_13pi9cy",
									TargetRef: "Activity_1567dil",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0qufezv",
										},
									},
									SourceRef: "Activity_1567dil",
									TargetRef: "Event_1uj58zc",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1jp11xh",
										},
									},
									SourceRef: "Event_1uj58zc",
									TargetRef: "Gateway_0tmub1o",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0pzvylb",
										},
									},
									SourceRef: "Gateway_0tmub1o",
									TargetRef: "Activity_1uixx0o",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1fuxj6b",
										},
									},
									SourceRef: "Activity_1uixx0o",
									TargetRef: "Activity_05503d3",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1ikbbnd",
										},
									},
									SourceRef: "Activity_05503d3",
									TargetRef: "Gateway_0vtzq53",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1ysfo4f",
										},
									},
									SourceRef: "Gateway_0vtzq53",
									TargetRef: "Event_0qacxgu",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_07uts0i",
										},
									},
									SourceRef: "Gateway_0vtzq53",
									TargetRef: "Event_0axou2i",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_11nkk44",
										},
									},
									SourceRef: "Gateway_0vtzq53",
									TargetRef: "Event_0bjpe63",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_01sh6l9",
										},
									},
									SourceRef: "Event_0axou2i",
									TargetRef: "Gateway_1s7by1w",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0utj2cp",
										},
									},
									SourceRef: "Event_0qacxgu",
									TargetRef: "Gateway_1s7by1w",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1dqygu2",
										},
									},
									SourceRef: "Event_0bjpe63",
									TargetRef: "Gateway_1s7by1w",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1367fco",
										},
									},
									SourceRef: "Gateway_1s7by1w",
									TargetRef: "Activity_0nzk36a",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0n1kpzs",
										},
									},
									SourceRef: "Activity_0nzk36a",
									TargetRef: "Activity_0yq2h1t",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1acnq20",
										},
									},
									SourceRef: "Activity_0yq2h1t",
									TargetRef: "Event_1dg5pp3",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_13ztu1l",
										},
									},
									SourceRef: "Activity_0pb64l8",
									TargetRef: "Activity_0fxxj1x",
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
