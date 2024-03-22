package bpmn

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sparrow-community/bpm/model/bpmn/instance"
)

func TestB_2_0_export(t *testing.T) {
	path := "./test/B.2.0-export.bpmn"
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
			ID:              "Definitions_0irsq5b",
			TargetNamespace: "http://bpmn.io/schema/bpmn",
			Exporter:        "Camunda Modeler",
			ExporterVersion: "5.2.0",
			RootElemnts: instance.RootElemnts{
				Collaborations: []instance.Collaboration{
					{
						RootElement: instance.RootElement{
							BaseElement: instance.BaseElement{
								ID: "Collaboration_1yp4w51",
							},
						},
						Participants: []instance.Participant{
							{
								BaseElement: instance.BaseElement{
									ID: "Participant_0d1io6b",
								},
								Name:       "Participant",
								ProcessRef: "Process_0nca5ry",
							},
							{
								BaseElement: instance.BaseElement{
									ID: "Participant_0s1hgls",
								},
								Name:       "Pool",
								ProcessRef: "Process_1xz7va4",
							},
						},
						MessageFlows: []instance.MessageFlow{
							{
								BaseElement: instance.BaseElement{
									ID: "Flow_1gv8gte",
								},
								Name:      "Message Flow 1",
								SourceRef: "Activity_0pn331p",
								TargetRef: "Event_0yxib88",
							},
							{
								BaseElement: instance.BaseElement{
									ID: "Flow_12bkpzd",
								},
								Name:      "Message Flow 2",
								SourceRef: "Event_0dwi53b",
								TargetRef: "Activity_0javi3i",
							},
						},
						Artifacts: instance.Artifacts{
							Groups: []instance.Group{
								{
									Artifact: instance.Artifact{
										BaseElement: instance.BaseElement{
											ID: "Group_1mitxsy",
										},
									},
									CategoryValueRef: "CategoryValue_02um9y9",
								},
							},
						},
					},
				},
				Processes: []instance.Process{
					{
						CallableElement: instance.CallableElement{
							RootElement: instance.RootElement{
								BaseElement: instance.BaseElement{
									ID: "Process_0nca5ry",
								},
							},
						},
						IsExecutable: true,
						FlowElements: instance.FlowElements{
							SequenceFlows: []instance.SequenceFlow{
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0gzdgm8",
										},
									},
									SourceRef: "StartEvent_1",
									TargetRef: "Activity_1uu0yo4",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0xky4tu",
										},
									},
									SourceRef: "Activity_1uu0yo4",
									TargetRef: "Activity_0pn331p",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1qqutfv",
										},
									},
									SourceRef: "Activity_0pn331p",
									TargetRef: "Activity_07vdczu",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0vyan0i",
										},
									},
									SourceRef: "Activity_07vdczu",
									TargetRef: "Gateway_0iz1sti",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1fmiuuz",
										},
										Name: "Conditional Sequence Flow",
									},
									SourceRef: "Gateway_0iz1sti",
									TargetRef: "Activity_00kk0w1",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0uffa0a",
										},
									},
									SourceRef: "Activity_00kk0w1",
									TargetRef: "Event_128e9tk",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_10j3phf",
										},
									},
									SourceRef: "Event_128e9tk",
									TargetRef: "Activity_03q0xwc",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_02vtnpn",
										},
									},
									SourceRef: "Activity_03q0xwc",
									TargetRef: "Activity_0thkytf",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0t9173h",
										},
									},
									SourceRef: "Activity_0thkytf",
									TargetRef: "Gateway_0y5g78s",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1exfm6h",
										},
									},
									SourceRef: "Event_1rsgo7a",
									TargetRef: "Activity_0ayd8ln",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_19165ws",
										},
									},
									SourceRef: "Gateway_0y5g78s",
									TargetRef: "Event_0403g4k",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1lz1d6w",
										},
										Name: "Default Sequence Flow 1",
									},
									SourceRef: "Gateway_0iz1sti",
									TargetRef: "Activity_0qnc8vy",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1iyrc7z",
										},
									},
									SourceRef: "Activity_0qnc8vy",
									TargetRef: "Activity_0ffbdg4",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1md83y0",
										},
									},
									SourceRef: "Activity_0ayd8ln",
									TargetRef: "Gateway_0y5g78s",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1b1ivbn",
										},
									},
									SourceRef: "Activity_0ffbdg4",
									TargetRef: "Activity_0s6eb2u",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0i7jvk9",
										},
									},
									SourceRef: "Event_1f35b4w",
									TargetRef: "Activity_0wdnxu3",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0rs1jsj",
										},
									},
									SourceRef: "Activity_0wdnxu3",
									TargetRef: "Event_0pi17ux",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1luuks1",
										},
									},
									SourceRef: "Event_0pi17ux",
									TargetRef: "Activity_0javi3i",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1vs1xz8",
										},
									},
									SourceRef: "Activity_0javi3i",
									TargetRef: "Event_1cb9pew",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0uy5u27",
										},
									},
									SourceRef: "Activity_0s6eb2u",
									TargetRef: "Gateway_0y5g78s",
								},
							},
							SubProcesses: []instance.SubProcess{
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Activity_0ffbdg4",
												},
												Name: "Expanded Sub-Process 1",
											},
											Incoming: []string{"Flow_1iyrc7z"},
											Outgoing: []string{"Flow_1b1ivbn"},
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
																	ID: "Event_1umydnx",
																},
																Name: "Start Event 2",
															},
															Outgoing: []string{"Flow_0pqufnb"},
														},
													},
												},
											},
										},
										SequenceFlows: []instance.SequenceFlow{
											{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Flow_0pqufnb",
													},
												},
												SourceRef: "Event_1umydnx",
												TargetRef: "Activity_1n1hhyt",
											},
											{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Flow_1qzdagc",
													},
												},
												SourceRef: "Activity_1n1hhyt",
												TargetRef: "Event_1bycw4y",
											},
										},
										UserTasks: []instance.UserTask{
											{
												Task: instance.Task{
													Activity: instance.Activity{
														FlowNode: instance.FlowNode{
															FlowElement: instance.FlowElement{
																BaseElement: instance.BaseElement{
																	ID: "Activity_1n1hhyt",
																},
																Name: "User Task 7 Standard Loop",
															},
															Incoming: []string{"Flow_0pqufnb"},
															Outgoing: []string{"Flow_1qzdagc"},
														},
														LoopCharacteristicsElements: instance.LoopCharacteristicsElements{
															StandardLoopCharacteristics: []instance.StandardLoopCharacteristics{{}},
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
																	ID: "Event_1bycw4y",
																},
																Name: "End Event 2",
															},
															Incoming: []string{"Flow_1qzdagc"},
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
													ID: "Activity_03q0xwc",
												},
												Name: "Collapsed Sub-Process 1 Multi-instances",
											},
											Incoming: []string{"Flow_10j3phf"},
											Outgoing: []string{"Flow_02vtnpn"},
										},
										LoopCharacteristicsElements: instance.LoopCharacteristicsElements{
											MultiInstanceLoopCharacteristics: []instance.MultiInstanceLoopCharacteristics{{}},
										},
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
														ID: "Event_128e9tk",
													},
													Name: "Intermediate Event Signal Throw 1",
												},
												Incoming: []string{"Flow_0uffa0a"},
												Outgoing: []string{"Flow_10j3phf"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											SignalEventDefinitions: []instance.SignalEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "SignalEventDefinition_13bzj0y",
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
														ID: "Event_0pi17ux",
													},
													Name: "Intermediate Event Conditional Catch",
												},
												Incoming: []string{"Flow_0rs1jsj"},
												Outgoing: []string{"Flow_1luuks1"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											ConditionalEventDefinitions: []instance.ConditionalEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "ConditionalEventDefinition_1hzmt1w",
															},
														},
													},
													Condition: instance.ExpressionUnMarshal{
														Type:                   "tFormalExpression",
														ExpressionSubstitution: &instance.FormalExpression{},
													},
												},
											},
										},
									},
								},
							},
							ServiceTasks: []instance.ServiceTask{
								{
									Task: instance.Task{
										Activity: instance.Activity{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Activity_00kk0w1",
													},
													Name: "Service Task 4",
												},
												Incoming: []string{"Flow_1fmiuuz"},
												Outgoing: []string{"Flow_0uffa0a"},
											},
										},
									},
								},
							},
							InclusiveGatewaies: []instance.InclusiveGateway{
								{
									Gateway: instance.Gateway{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Gateway_0iz1sti",
												},
												Name: "Inclusive Gateway",
											},
											Incoming: []string{"Flow_0vyan0i"},
											Outgoing: []string{"Flow_1fmiuuz", "Flow_1lz1d6w"},
										},
									},
									Default: "Flow_1lz1d6w",
								},
							},
							UserTasks: []instance.UserTask{
								{
									Task: instance.Task{
										Activity: instance.Activity{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Activity_07vdczu",
													},
													Name: "User Task 3",
												},
												Incoming: []string{"Flow_1qqutfv"},
												Outgoing: []string{"Flow_0vyan0i"},
											},
											Properties: []instance.Property{
												{
													BaseElement: instance.BaseElement{
														ID: "Property_0qjltqr",
													},
													Name: "__targetRef_placeholder",
												},
											},
											DataInputAssociations: []instance.DataInputAssociation{
												{
													DataAssociation: instance.DataAssociation{
														BaseElement: instance.BaseElement{
															ID: "DataInputAssociation_1442h6j",
														},
														SourceRef: "DataObjectReference_1rkw6dp",
														TargetRef: "Property_0qjltqr",
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
														ID: "Activity_0s6eb2u",
													},
													Name: "User Task 8",
												},
												Incoming: []string{"Flow_1b1ivbn"},
												Outgoing: []string{"Flow_0uy5u27"},
											},
										},
									},
								},
							},
							SendTasks: []instance.SendTask{
								{
									Task: instance.Task{
										Activity: instance.Activity{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Activity_0pn331p",
													},
													Name: "Send Task 2",
												},
												Incoming: []string{"Flow_0xky4tu"},
												Outgoing: []string{"Flow_1qqutfv"},
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
													ID: "Activity_1uu0yo4",
												},
												Name: "Abstract Task 1",
											},
											Incoming: []string{"Flow_0gzdgm8"},
											Outgoing: []string{"Flow_0xky4tu"},
										},
									},
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Activity_0thkytf",
												},
												Name: "Task 5",
											},
											Incoming: []string{"Flow_02vtnpn"},
											Outgoing: []string{"Flow_0t9173h"},
										},
									},
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Activity_0ayd8ln",
												},
												Name: "Task 6",
											},
											Incoming: []string{"Flow_1exfm6h"},
											Outgoing: []string{"Flow_1md83y0"},
										},
									},
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Activity_0wdnxu3",
												},
												Name: "Task 9",
											},
											Incoming: []string{"Flow_0i7jvk9"},
											Outgoing: []string{"Flow_0rs1jsj"},
										},
									},
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Activity_0javi3i",
												},
												Name: "Task 10",
											},
											Incoming: []string{"Flow_1luuks1"},
											Outgoing: []string{"Flow_1vs1xz8"},
										},
									},
								},
							},
							StartEvents: []instance.StartEvent{
								{
									CatchEvent: instance.CatchEvent{
										Event: instance.Event{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "StartEvent_1",
													},
													Name: "Start Event 1 Timer",
												},
												Outgoing: []string{"Flow_0gzdgm8"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											TimerEventDefinitions: []instance.TimerEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "TimerEventDefinition_1wg06e7",
															},
														},
													},
												},
											},
										},
									},
								},
							},
							CallActivities: []instance.CallActivity{
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Activity_0qnc8vy",
												},
												Name: "Call Activity calling a Global User Task",
											},
											Incoming: []string{"Flow_1lz1d6w"},
											Outgoing: []string{"Flow_1iyrc7z"},
										},
									},
								},
							},
							DataObjectReferenes: []instance.DataObjectReference{
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "DataObjectReference_1rkw6dp",
										},
										Name: "Data Object",
									},
									DataObjectRef: "DataObject_0v253gd",
								},
							},
							DataObjects: []instance.DataObject{
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "DataObject_0v253gd",
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
														ID: "Event_1f35b4w",
													},
												},
												Outgoing: []string{"Flow_0i7jvk9"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											EscalationEventDefinitions: []instance.EscalationEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "EscalationEventDefinition_1e7kboz",
															},
														},
													},
												},
											},
										},
									},
									CancelActivity: false,
									AttachedToRef:  "Activity_0s6eb2u",
								},
								{
									CatchEvent: instance.CatchEvent{
										Event: instance.Event{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Event_1rsgo7a",
													},
													Name: "Boundary Intermediate Event Non-Interrupting Conditional",
												},
												Outgoing: []string{"Flow_1exfm6h"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											ConditionalEventDefinitions: []instance.ConditionalEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "ConditionalEventDefinition_1uszs8x",
															},
														},
													},
													Condition: instance.ExpressionUnMarshal{
														Type:                   "tFormalExpression",
														ExpressionSubstitution: &instance.FormalExpression{},
													},
												},
											},
										},
									},
									CancelActivity: false,
									AttachedToRef:  "Activity_0thkytf",
								},
							},
							EndEvents: []instance.EndEvent{
								{
									ThrowEvent: instance.ThrowEvent{
										Event: instance.Event{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Event_0403g4k",
													},
													Name: "End Event 1 Message",
												},
												Incoming: []string{"Flow_19165ws"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											MessageEventDefinitions: []instance.MessageEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "MessageEventDefinition_00eayw1",
															},
														},
													},
												},
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
														ID: "Event_1cb9pew",
													},
													Name: "End Event 3 Signal",
												},
												Incoming: []string{"Flow_1vs1xz8"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											SignalEventDefinitions: []instance.SignalEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "SignalEventDefinition_1v2uxma",
															},
														},
													},
												},
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
													ID: "Gateway_0y5g78s",
												},
												Name: "Parallel Gateway 2",
											},
											Incoming: []string{"Flow_0t9173h", "Flow_1md83y0", "Flow_0uy5u27"},
											Outgoing: []string{"Flow_19165ws"},
										},
									},
								},
							},
						},
						Artifacts: instance.Artifacts{
							TextAnnotations: []instance.TextAnnotation{
								{
									Artifact: instance.Artifact{
										BaseElement: instance.BaseElement{
											ID: "TextAnnotation_06uwg6s",
										},
									},
									Text: "Annotation",
								},
							},
							Associations: []instance.Association{
								{
									Artifact: instance.Artifact{
										BaseElement: instance.BaseElement{
											ID: "Association_0ahr1rq",
										},
									},
									SourceRef: "Activity_0ffbdg4",
									TargetRef: "TextAnnotation_06uwg6s",
								},
							},
						},
					},
					{
						CallableElement: instance.CallableElement{
							RootElement: instance.RootElement{
								BaseElement: instance.BaseElement{
									ID: "Process_1xz7va4",
								},
							},
						},
						IsExecutable: false,
						LaneSets: []instance.LaneSet{
							{
								BaseElement: instance.BaseElement{
									ID: "LaneSet_02xvcik",
								},
								Lanes: []instance.Lane{
									{
										BaseElement: instance.BaseElement{
											ID: "Lane_07fps3z",
										},
										Name: "Lane 1",
										FlowNodeRefs: []string{
											"Event_0n9dn9q",
											"Gateway_1n5jtse",
											"Activity_0dbod77",
											"Event_0yxib88",
											"Event_0dwi53b",
											"Activity_1i3ap44",
											"Event_0b4vgd7",
											"Activity_1xwdlgt",
											"Activity_0u4otp8",
											"Activity_1gsxnu3",
											"Activity_1kayloh",
											"Event_1f1af1t",
											"Activity_01h1c3z",
											"Activity_0wxomwg",
											"Event_1qwowa6",
											"Gateway_156l12u",
											"Activity_1fag12l",
											"Activity_1a394xz",
											"Activity_17ocq3z",
											"Activity_1ig20jv",
											"Event_1gbx6rp",
											"Event_1brbeoj",
											"Event_1e5rzkm",
											"Event_087hfz9",
											"Event_0drv2q5",
											"Event_0gdj0lz",
											"Event_1qgdkdf",
											"Activity_1qy7xgp",
											"Event_009wd9m",
											"Event_1lph52r",
											"Event_1r9ahil",
										},
									},
									{
										BaseElement: instance.BaseElement{
											ID: "Lane_1eeat7r",
										},
										Name: "Lane 2",
										FlowNodeRefs: []string{
											"Activity_1c0usbk",
											"Activity_19dq1xp",
											"Activity_11l0l4d",
											"Activity_1cwg4vs",
											"Event_1rx2a8f",
											"Activity_1rcfmjq",
											"Activity_06qihxs",
											"Event_15p2cvq",
											"Activity_0xe0m5g",
											"Event_14urgdq",
											"Activity_15jh1tp",
											"Gateway_0ynnz6a",
											"Activity_0ahqya2",
											"Event_1qd721q",
											"Event_1ofkxhs",
											"Gateway_0gp6gne",
											"Gateway_198yhv3",
											"Event_0sedygx",
											"Event_0kp736q",
										},
									},
								},
							},
						},
						FlowElements: instance.FlowElements{
							SequenceFlows: []instance.SequenceFlow{
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0m8w25k",
										},
									},
									SourceRef: "Activity_0dbod77",
									TargetRef: "Gateway_1n5jtse",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_13lye0k",
										},
									},
									SourceRef: "Gateway_1n5jtse",
									TargetRef: "Event_0n9dn9q",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0xjmqod",
										},
									},
									SourceRef: "Gateway_1n5jtse",
									TargetRef: "Event_0gdj0lz",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_051lggk",
										},
									},
									SourceRef: "Gateway_1n5jtse",
									TargetRef: "Event_1qgdkdf",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_138p74x",
										},
									},
									SourceRef: "Event_0yxib88",
									TargetRef: "Activity_0dbod77",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1j9rxo8",
										},
									},
									SourceRef: "Event_0n9dn9q",
									TargetRef: "Activity_1fag12l",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1cxmbe8",
										},
									},
									SourceRef: "Event_1qgdkdf",
									TargetRef: "Activity_1a394xz",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0xx9p19",
										},
									},
									SourceRef: "Activity_1fag12l",
									TargetRef: "Gateway_156l12u",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_16g4olt",
										},
										Name: "Default Sequence Flow 2",
									},
									SourceRef: "Gateway_156l12u",
									TargetRef: "Event_1qwowa6",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0vv68hn",
										},
									},
									SourceRef: "Event_1qwowa6",
									TargetRef: "Activity_0wxomwg",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0dk5zr9",
										},
									},
									SourceRef: "Event_087hfz9",
									TargetRef: "Activity_17ocq3z",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0prsfbf",
										},
									},
									SourceRef: "Gateway_156l12u",
									TargetRef: "Activity_17ocq3z",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0npffus",
										},
									},
									SourceRef: "Activity_01h1c3z",
									TargetRef: "Activity_1qy7xgp",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_10q4nrz",
										},
									},
									SourceRef: "Event_009wd9m",
									TargetRef: "Event_1lph52r",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0rx4vqb",
										},
									},
									SourceRef: "Activity_1qy7xgp",
									TargetRef: "Activity_1i3ap44",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0qvdwb3",
										},
									},
									SourceRef: "Activity_1i3ap44",
									TargetRef: "Event_0dwi53b",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0yk73ad",
										},
									},
									SourceRef: "Activity_0wxomwg",
									TargetRef: "Activity_1xwdlgt",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_07zdeu5",
										},
									},
									SourceRef: "Activity_1xwdlgt",
									TargetRef: "Event_0b4vgd7",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_01k51sm",
										},
									},
									SourceRef: "Activity_1ig20jv",
									TargetRef: "Activity_0u4otp8",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0zgmk02",
										},
									},
									SourceRef: "Activity_17ocq3z",
									TargetRef: "Event_1f1af1t",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0zh9gvr",
										},
									},
									SourceRef: "Event_1f1af1t",
									TargetRef: "Activity_0u4otp8",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0twe4nl",
										},
									},
									SourceRef: "Event_1r9ahil",
									TargetRef: "Activity_1gsxnu3",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0qd2cjd",
										},
									},
									SourceRef: "Activity_1gsxnu3",
									TargetRef: "Event_0b4vgd7",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1ef3aqe",
										},
									},
									SourceRef: "Activity_0u4otp8",
									TargetRef: "Event_0b4vgd7",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1p5rcqt",
										},
									},
									SourceRef: "Event_0drv2q5",
									TargetRef: "Activity_1kayloh",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0vzv0ky",
										},
									},
									SourceRef: "Activity_1a394xz",
									TargetRef: "Activity_1ig20jv",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_02c98k2",
										},
									},
									SourceRef: "Event_1qd721q",
									TargetRef: "Activity_0ahqya2",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0jmy9bt",
										},
									},
									SourceRef: "Activity_0ahqya2",
									TargetRef: "Gateway_0ynnz6a",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_163ttxe",
										},
									},
									SourceRef: "Gateway_0ynnz6a",
									TargetRef: "Activity_15jh1tp",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1yz425v",
										},
									},
									SourceRef: "Activity_15jh1tp",
									TargetRef: "Event_14urgdq",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_11zr1a3",
										},
									},
									SourceRef: "Event_14urgdq",
									TargetRef: "Gateway_0gp6gne",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0uaxzdx",
										},
									},
									SourceRef: "Event_1e5rzkm",
									TargetRef: "Activity_1c0usbk",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_00uq1g7",
										},
									},
									SourceRef: "Activity_1c0usbk",
									TargetRef: "Gateway_0gp6gne",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0ax4o0j",
										},
									},
									SourceRef: "Gateway_0gp6gne",
									TargetRef: "Activity_0xe0m5g",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1a59vjg",
										},
									},
									SourceRef: "Event_1gbx6rp",
									TargetRef: "Activity_1rcfmjq",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1k2of3s",
										},
									},
									SourceRef: "Activity_1rcfmjq",
									TargetRef: "Event_1rx2a8f",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0oyt612",
										},
									},
									SourceRef: "Event_1brbeoj",
									TargetRef: "Activity_06qihxs",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_07hgdqa",
										},
									},
									SourceRef: "Activity_06qihxs",
									TargetRef: "Event_0sedygx",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0dh0nnk",
										},
									},
									SourceRef: "Activity_1kayloh",
									TargetRef: "Event_0sedygx",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0fop3s1",
										},
									},
									SourceRef: "Gateway_0ynnz6a",
									TargetRef: "Activity_19dq1xp",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0xqi2l5",
										},
									},
									SourceRef: "Event_1ofkxhs",
									TargetRef: "Activity_19dq1xp",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0ox4w4d",
										},
									},
									SourceRef: "Activity_19dq1xp",
									TargetRef: "Activity_11l0l4d",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1i60rky",
										},
									},
									SourceRef: "Event_15p2cvq",
									TargetRef: "Activity_1cwg4vs",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0r5tyx6",
										},
									},
									SourceRef: "Activity_1cwg4vs",
									TargetRef: "Event_0kp736q",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0ly8iho",
										},
									},
									SourceRef: "Activity_11l0l4d",
									TargetRef: "Gateway_198yhv3",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0rkhydx",
										},
									},
									SourceRef: "Activity_0xe0m5g",
									TargetRef: "Gateway_198yhv3",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1tqpq63",
										},
									},
									SourceRef: "Gateway_198yhv3",
									TargetRef: "Event_0sedygx",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0l823y4",
										},
									},
									SourceRef: "Event_0gdj0lz",
									TargetRef: "Activity_01h1c3z",
								},
							},
							IntermediateThrowEvents: []instance.IntermediateThrowEvent{
								{
									ThrowEvent: instance.ThrowEvent{
										Event: instance.Event{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Event_1qwowa6",
													},
													Name: "Intermediate Event Message Throw",
												},
												Incoming: []string{"Flow_16g4olt"},
												Outgoing: []string{"Flow_0vv68hn"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											MessageEventDefinitions: []instance.MessageEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "MessageEventDefinition_0io8u5r",
															},
														},
													},
												},
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
														ID: "Event_1f1af1t",
													},
												},
												Incoming: []string{"Flow_0zgmk02"},
												Outgoing: []string{"Flow_0zh9gvr"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											EscalationEventDefinitions: []instance.EscalationEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "EscalationEventDefinition_185ama0",
															},
														},
													},
												},
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
														ID: "Event_1lph52r",
													},
													Name: "Intermediate End Event",
												},
												Incoming: []string{"Flow_10q4nrz"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											LinkEventDefinitions: []instance.LinkEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "LinkEventDefinition_0a59ub9",
															},
														},
													},
													Name: "",
												},
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
														ID: "Event_14urgdq",
													},
													Name: "Intermediate Event Signal Throw 2",
												},
												Incoming: []string{"Flow_1yz425v"},
												Outgoing: []string{"Flow_11zr1a3"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											SignalEventDefinitions: []instance.SignalEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "SignalEventDefinition_0jhkcy9",
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
														ID: "Event_0n9dn9q",
													},
													Name: "Intermediate Event Message Catch",
												},
												Incoming: []string{"Flow_13lye0k"},
												Outgoing: []string{"Flow_1j9rxo8"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											MessageEventDefinitions: []instance.MessageEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "MessageEventDefinition_0ox4djh",
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
														ID: "Event_0gdj0lz",
													},
													Name: "Intermediate Event Timer Catch",
												},
												Incoming: []string{"Flow_0xjmqod"},
												Outgoing: []string{"Flow_0l823y4"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											TimerEventDefinitions: []instance.TimerEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "TimerEventDefinition_1vvsiz9",
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
														ID: "Event_1qgdkdf",
													},
													Name: "Intermediate Event Message Catch 2",
												},
												Incoming: []string{"Flow_051lggk"},
												Outgoing: []string{"Flow_1cxmbe8"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											MessageEventDefinitions: []instance.MessageEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "MessageEventDefinition_1nec7v4",
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
														ID: "Event_1ofkxhs",
													},
													Name: "Intermediate Event Link",
												},
												Outgoing: []string{"Flow_0xqi2l5"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											LinkEventDefinitions: []instance.LinkEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "LinkEventDefinition_0p5xwz6",
															},
														},
													},
													Name: "",
												},
											},
										},
									},
								},
							},
							EventBasedGatewaies: []instance.EventBasedGateway{
								{
									Gateway: instance.Gateway{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Gateway_1n5jtse",
												},
												Name: "Event Base Gateway 3",
											},
											Incoming: []string{"Flow_0m8w25k"},
											Outgoing: []string{"Flow_13lye0k", "Flow_0xjmqod", "Flow_051lggk"},
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
													ID: "Activity_0dbod77",
												},
												Name: "Task 11",
											},
											Incoming: []string{"Flow_138p74x"},
											Outgoing: []string{"Flow_0m8w25k"},
										},
									},
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Activity_1a394xz",
												},
												Name: "Task 21",
											},
											Incoming: []string{"Flow_1cxmbe8"},
											Outgoing: []string{"Flow_0vzv0ky"},
										},
									},
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Activity_17ocq3z",
												},
												Name: "Task 18",
											},
											Incoming: []string{"Flow_0dk5zr9", "Flow_0prsfbf"},
											Outgoing: []string{"Flow_0zgmk02"},
										},
									},
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Activity_1c0usbk",
												},
												Name: "Task 27",
											},
											Incoming: []string{"Flow_0uaxzdx"},
											Outgoing: []string{"Flow_00uq1g7"},
										},
									},
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Activity_11l0l4d",
												},
												Name: "Task 32",
											},
											Incoming: []string{"Flow_0ox4w4d"},
											Outgoing: []string{"Flow_0ly8iho"},
										},
									},
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Activity_1cwg4vs",
												},
												Name: "Task 33",
											},
											Incoming: []string{"Flow_1i60rky"},
											Outgoing: []string{"Flow_0r5tyx6"},
										},
									},
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Activity_1rcfmjq",
												},
												Name: "Task 29",
											},
											Incoming: []string{"Flow_1a59vjg"},
											Outgoing: []string{"Flow_1k2of3s"},
										},
									},
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Activity_06qihxs",
												},
												Name: "Task 30",
											},
											Incoming: []string{"Flow_0oyt612"},
											Outgoing: []string{"Flow_07hgdqa"},
										},
									},
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Activity_1xwdlgt",
												},
												Name: "Task 17",
											},
											Incoming: []string{"Flow_0yk73ad"},
											Outgoing: []string{"Flow_07zdeu5"},
										},
									},
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Activity_0u4otp8",
												},
												Name: "Task 23",
											},
											Incoming: []string{"Flow_01k51sm", "Flow_0zh9gvr"},
											Outgoing: []string{"Flow_1ef3aqe"},
										},
									},
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Activity_1gsxnu3",
												},
												Name: "Task 19",
											},
											Incoming: []string{"Flow_0twe4nl"},
											Outgoing: []string{"Flow_0qd2cjd"},
										},
									},
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Activity_1kayloh",
												},
												Name: "Task 24",
											},
											Incoming: []string{"Flow_1p5rcqt"},
											Outgoing: []string{"Flow_0dh0nnk"},
										},
									},
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Activity_0xe0m5g",
												},
												Name: "Task 28",
											},
											Incoming: []string{"Flow_0ax4o0j"},
											Outgoing: []string{"Flow_0rkhydx"},
										},
									},
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Activity_15jh1tp",
												},
												Name: "Task 26",
											},
											Incoming: []string{"Flow_163ttxe"},
											Outgoing: []string{"Flow_1yz425v"},
										},
									},
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Activity_0ahqya2",
												},
												Name: "Task 25",
											},
											Incoming: []string{"Flow_02c98k2"},
											Outgoing: []string{"Flow_0jmy9bt"},
										},
									},
								},
							},
							StartEvents: []instance.StartEvent{
								{
									CatchEvent: instance.CatchEvent{
										Event: instance.Event{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Event_0yxib88",
													},
													Name: "Start Event 2 Message",
												},
												Outgoing: []string{"Flow_138p74x"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											MessageEventDefinitions: []instance.MessageEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "MessageEventDefinition_1gmfphr",
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
														ID: "Event_1qd721q",
													},
													Name: "Start Event 6 Signal",
												},
												Outgoing: []string{"Flow_02c98k2"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											SignalEventDefinitions: []instance.SignalEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "SignalEventDefinition_02lr7ru",
															},
														},
													},
												},
											},
										},
									},
								},
							},
							DataStoreReferences: []instance.DataStoreReference{
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "DataStoreReference_0t1ift6",
										},
										Name: "Data Store Reference",
									},
								},
							},
							SubProcesses: []instance.SubProcess{
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Activity_01h1c3z",
												},
												Name: "Expanded Call Activity",
											},
											Incoming: []string{"Flow_0l823y4"},
											Outgoing: []string{"Flow_0npffus"},
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
																	ID: "Event_05w6wy8",
																},
																Name: "Start Event 3",
															},
															Outgoing: []string{"Flow_18dr9jm"},
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
																	ID: "Event_1wq0sy2",
																},
																Name: "Start Event 4 Conditional",
															},
															Outgoing: []string{"Flow_1sl2jua"},
														},
													},
													EventDefinitions: instance.EventDefinitions{
														ConditionalEventDefinitions: []instance.ConditionalEventDefinition{
															{
																EventDefinition: instance.EventDefinition{
																	RootElement: instance.RootElement{
																		BaseElement: instance.BaseElement{
																			ID: "ConditionalEventDefinition_0cq4mx7",
																		},
																	},
																},
																Condition: instance.ExpressionUnMarshal{
																	Type:                   "tFormalExpression",
																	ExpressionSubstitution: &instance.FormalExpression{},
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
														ID: "Flow_18dr9jm",
													},
												},
												SourceRef: "Event_05w6wy8",
												TargetRef: "Activity_1lrpf00",
											},
											{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Flow_1sl2jua",
													},
												},
												SourceRef: "Event_1wq0sy2",
												TargetRef: "Activity_1lrpf00",
											},
											{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Flow_13o472r",
													},
												},
												SourceRef: "Activity_1lrpf00",
												TargetRef: "Activity_0g3cmxi",
											},
											{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Flow_0agposr",
													},
												},
												SourceRef: "Activity_0g3cmxi",
												TargetRef: "Activity_0pyeguq",
											},
											{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Flow_0e3d4iw",
													},
												},
												SourceRef: "Event_0mg1ari",
												TargetRef: "Event_0a5qvbg",
											},
											{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Flow_10obnjq",
													},
												},
												SourceRef: "Activity_0pyeguq",
												TargetRef: "Event_153rv56",
											},
										},
										UserTasks: []instance.UserTask{
											{
												Task: instance.Task{
													Activity: instance.Activity{
														FlowNode: instance.FlowNode{
															FlowElement: instance.FlowElement{
																BaseElement: instance.BaseElement{
																	ID: "Activity_1lrpf00",
																},
																Name: "User Task 12 Multi-inst Seq.",
															},
															Incoming: []string{"Flow_18dr9jm", "Flow_1sl2jua"},
															Outgoing: []string{"Flow_13o472r"},
														},
														LoopCharacteristicsElements: instance.LoopCharacteristicsElements{
															MultiInstanceLoopCharacteristics: []instance.MultiInstanceLoopCharacteristics{{
																IsSequential: true,
															}},
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
																	ID: "Activity_0g3cmxi",
																},
																Name: "User Task 13",
															},
															Incoming: []string{"Flow_13o472r"},
															Outgoing: []string{"Flow_0agposr"},
														},
													},
												},
											},
										},
										ServiceTasks: []instance.ServiceTask{
											{
												Task: instance.Task{
													Activity: instance.Activity{
														FlowNode: instance.FlowNode{
															FlowElement: instance.FlowElement{
																BaseElement: instance.BaseElement{
																	ID: "Activity_0pyeguq",
																},
																Name: "Service Task 14",
															},
															Incoming: []string{"Flow_0agposr"},
															Outgoing: []string{"Flow_10obnjq"},
														},
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
																	ID: "Event_0mg1ari",
																},
																Name: "Boundary Intermediate Event Interrupting Message",
															},
															Outgoing: []string{"Flow_0e3d4iw"},
														},
													},
													EventDefinitions: instance.EventDefinitions{
														MessageEventDefinitions: []instance.MessageEventDefinition{
															{
																EventDefinition: instance.EventDefinition{
																	RootElement: instance.RootElement{
																		BaseElement: instance.BaseElement{
																			ID: "MessageEventDefinition_0latlfm",
																		},
																	},
																},
															},
														},
													},
												},
												AttachedToRef: "Activity_0g3cmxi",
											},
										},
										EndEvents: []instance.EndEvent{
											{
												ThrowEvent: instance.ThrowEvent{
													Event: instance.Event{
														FlowNode: instance.FlowNode{
															FlowElement: instance.FlowElement{
																BaseElement: instance.BaseElement{
																	ID: "Event_0a5qvbg",
																},
																Name: "End Event 5 Terminate",
															},
															Incoming: []string{"Flow_0e3d4iw"},
														},
													},
													EventDefinitions: instance.EventDefinitions{
														TerminateEventDefinitions: []instance.TerminateEventDefinition{
															{
																EventDefinition: instance.EventDefinition{
																	RootElement: instance.RootElement{
																		BaseElement: instance.BaseElement{
																			ID: "TerminateEventDefinition_14p6uf0",
																		},
																	},
																},
															},
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
																	ID: "Event_153rv56",
																},
																Name: "End Event 4",
															},
															Incoming: []string{"Flow_10obnjq"},
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
													ID: "Activity_1fag12l",
												},
												Name: "Collapsed Sub-Process 2",
											},
											Incoming: []string{"Flow_1j9rxo8"},
											Outgoing: []string{"Flow_0xx9p19"},
										},
									},
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Activity_19dq1xp",
												},
												Name: "Expanded Sub-Process 3",
											},
											Incoming: []string{"Flow_0fop3s1", "Flow_0xqi2l5"},
											Outgoing: []string{"Flow_0ox4w4d"},
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
																	ID: "Event_0wm5nle",
																},
																Name: "Start Event 7",
															},
															Outgoing: []string{"Flow_01v892m"},
														},
													},
												},
											},
										},
										SequenceFlows: []instance.SequenceFlow{
											{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Flow_01v892m",
													},
												},
												SourceRef: "Event_0wm5nle",
												TargetRef: "Event_16rqalg",
											},
											{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Flow_0shj0yz",
													},
												},
												SourceRef: "Event_16rqalg",
												TargetRef: "Activity_03j7y8k",
											},
											{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Flow_1nnas72",
													},
												},
												SourceRef: "Activity_03j7y8k",
												TargetRef: "Gateway_10qt5ud",
											},
											{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Flow_0bwg7yc",
													},
												},
												SourceRef: "Gateway_10qt5ud",
												TargetRef: "Event_0ivkjpo",
											},
											{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Flow_1ierw66",
													},
												},
												SourceRef: "Gateway_10qt5ud",
												TargetRef: "Event_1wir0h9",
											},
										},
										IntermediateCatchEvents: []instance.IntermediateCatchEvent{
											{
												CatchEvent: instance.CatchEvent{
													Event: instance.Event{
														FlowNode: instance.FlowNode{
															FlowElement: instance.FlowElement{
																BaseElement: instance.BaseElement{
																	ID: "Event_16rqalg",
																},
																Name: "Intermediate Event Signal Catch",
															},
															Incoming: []string{"Flow_01v892m"},
															Outgoing: []string{"Flow_0shj0yz"},
														},
													},
													EventDefinitions: instance.EventDefinitions{
														SignalEventDefinitions: []instance.SignalEventDefinition{
															{
																EventDefinition: instance.EventDefinition{
																	RootElement: instance.RootElement{
																		BaseElement: instance.BaseElement{
																			ID: "SignalEventDefinition_0gdrh6v",
																		},
																	},
																},
															},
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
																ID: "Activity_03j7y8k",
															},
															Name: "Task 31",
														},
														Incoming: []string{"Flow_0shj0yz"},
														Outgoing: []string{"Flow_1nnas72"},
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
																ID: "Gateway_10qt5ud",
															},
															Name: "Exclusive Gateway",
														},
														Incoming: []string{"Flow_1nnas72"},
														Outgoing: []string{"Flow_0bwg7yc", "Flow_1ierw66"},
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
																	ID: "Event_0ivkjpo",
																},
																Name: "End Event 12",
															},
															Incoming: []string{"Flow_0bwg7yc"},
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
																	ID: "Event_1wir0h9",
																},
																Name: "End Event 13 Error",
															},
															Incoming: []string{"Flow_1ierw66"},
														},
													},
													EventDefinitions: instance.EventDefinitions{
														ErrorEventDefinitions: []instance.ErrorEventDefinition{
															{
																EventDefinition: instance.EventDefinition{
																	RootElement: instance.RootElement{
																		BaseElement: instance.BaseElement{
																			ID: "ErrorEventDefinition_13if476",
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
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Activity_1ig20jv",
												},
												Name: "Expanded Sub-Process 2",
											},
											Incoming: []string{"Flow_0vzv0ky"},
											Outgoing: []string{"Flow_01k51sm"},
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
																	ID: "Event_05uva8p",
																},
																Name: "Start Event 5 None",
															},
															Outgoing: []string{"Flow_0ted34p"},
														},
													},
												},
											},
										},
										SequenceFlows: []instance.SequenceFlow{
											{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Flow_0ted34p",
													},
												},
												SourceRef: "Event_05uva8p",
												TargetRef: "Activity_1k00lfn",
											},
											{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Flow_07jjywp",
													},
												},
												SourceRef: "Activity_1k00lfn",
												TargetRef: "Event_0o1k45w",
											},
										},
										ServiceTasks: []instance.ServiceTask{
											{
												Task: instance.Task{
													Activity: instance.Activity{
														FlowNode: instance.FlowNode{
															FlowElement: instance.FlowElement{
																BaseElement: instance.BaseElement{
																	ID: "Activity_1k00lfn",
																},
																Name: "Service Task 22",
															},
															Incoming: []string{"Flow_0ted34p"},
															Outgoing: []string{"Flow_07jjywp"},
														},
														LoopCharacteristicsElements: instance.LoopCharacteristicsElements{
															MultiInstanceLoopCharacteristics: []instance.MultiInstanceLoopCharacteristics{{}},
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
																	ID: "Event_0o1k45w",
																},
																Name: "End Event 8 None",
															},
															Incoming: []string{"Flow_07jjywp"},
														},
													},
												},
											},
										},
									},
								},
							},
							CallActivities: []instance.CallActivity{
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Activity_0wxomwg",
												},
												Name: "Collapsed Call Activity",
											},
											Incoming: []string{"Flow_0vv68hn"},
											Outgoing: []string{"Flow_0yk73ad"},
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
														ID: "Event_087hfz9",
													},
													Name: "Boundary Intermediate Event Non-Interrupting Escalation",
												},
												Outgoing: []string{"Flow_0dk5zr9"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											EscalationEventDefinitions: []instance.EscalationEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "EscalationEventDefinition_0pw5off",
															},
														},
													},
												},
											},
										},
									},
									CancelActivity: false,
									AttachedToRef:  "Activity_0wxomwg",
								},
								{
									CatchEvent: instance.CatchEvent{
										Event: instance.Event{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Event_1e5rzkm",
													},
													Name: "Boundary Intermediate Event Interrupting Timer",
												},
												Outgoing: []string{"Flow_0uaxzdx"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											TimerEventDefinitions: []instance.TimerEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "TimerEventDefinition_13e8rvq",
															},
														},
													},
												},
											},
										},
									},
									AttachedToRef: "Activity_1a394xz",
								},
								{
									CatchEvent: instance.CatchEvent{
										Event: instance.Event{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Event_15p2cvq",
													},
												},
												Outgoing: []string{"Flow_1i60rky"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											SignalEventDefinitions: []instance.SignalEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "SignalEventDefinition_0e6ibwj",
															},
														},
													},
												},
											},
										},
									},
									AttachedToRef: "Activity_11l0l4d",
								},
								{
									CatchEvent: instance.CatchEvent{
										Event: instance.Event{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Event_1brbeoj",
													},
													Name: "Boundary Intermediate Event Non-Interrupting Timer",
												},
												Outgoing: []string{"Flow_0oyt612"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											TimerEventDefinitions: []instance.TimerEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "TimerEventDefinition_0d4yvld",
															},
														},
													},
												},
											},
										},
									},
									CancelActivity: false,
									AttachedToRef:  "Activity_1ig20jv",
								},
								{
									CatchEvent: instance.CatchEvent{
										Event: instance.Event{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Event_1gbx6rp",
													},
													Name: "Boundary Intermediate Event Interrupting Error",
												},
												Outgoing: []string{"Flow_1a59vjg"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											ErrorEventDefinitions: []instance.ErrorEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "ErrorEventDefinition_17g5j05",
															},
														},
													},
												},
											},
										},
									},
									AttachedToRef: "Activity_1ig20jv",
								},
								{
									CatchEvent: instance.CatchEvent{
										Event: instance.Event{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Event_0drv2q5",
													},
													Name: "Boundary Intermediate Event Non-Interrupting Signal",
												},
												Outgoing: []string{"Flow_1p5rcqt"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											SignalEventDefinitions: []instance.SignalEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "SignalEventDefinition_168in84",
															},
														},
													},
												},
											},
										},
									},
									CancelActivity: false,
									AttachedToRef:  "Activity_0u4otp8",
								},
								{
									CatchEvent: instance.CatchEvent{
										Event: instance.Event{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Event_009wd9m",
													},
													Name: "Boundary Intermediate Event Interrupting Conditional",
												},
												Outgoing: []string{"Flow_10q4nrz"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											ConditionalEventDefinitions: []instance.ConditionalEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "ConditionalEventDefinition_14j1wsr",
															},
														},
													},
													Condition: instance.ExpressionUnMarshal{
														Type:                   "tFormalExpression",
														ExpressionSubstitution: &instance.FormalExpression{},
													},
												},
											},
										},
									},
									AttachedToRef: "Activity_1qy7xgp",
								},
								{
									CatchEvent: instance.CatchEvent{
										Event: instance.Event{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Event_1r9ahil",
													},
													Name: "Boundary Intermediate Event Non-Interrupting Message",
												},
												Outgoing: []string{"Flow_0twe4nl"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											MessageEventDefinitions: []instance.MessageEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "MessageEventDefinition_1qa9e6d",
															},
														},
													},
												},
											},
										},
									},
									AttachedToRef: "Activity_1xwdlgt",
								},
							},
							ExclusiveGatewaies: []instance.ExclusiveGateway{
								{
									Gateway: instance.Gateway{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Gateway_156l12u",
												},
												Name: "Exclusive Gateway 4",
											},
											Incoming: []string{"Flow_0xx9p19"},
											Outgoing: []string{"Flow_16g4olt", "Flow_0prsfbf"},
										},
									},
									Default: "Flow_16g4olt",
								},
								{
									Gateway: instance.Gateway{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Gateway_198yhv3",
												},
												Name: "Parallel Gateway 7",
											},
											Incoming: []string{"Flow_0ly8iho", "Flow_0rkhydx"},
											Outgoing: []string{"Flow_1tqpq63"},
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
														ID: "Event_1rx2a8f",
													},
													Name: "End Event 10",
												},
												Incoming: []string{"Flow_1k2of3s"},
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
														ID: "Event_0dwi53b",
													},
													Name: "End Event 6 Message",
												},
												Incoming: []string{"Flow_0qvdwb3"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											MessageEventDefinitions: []instance.MessageEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "MessageEventDefinition_1tw5clq",
															},
														},
													},
												},
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
														ID: "Event_0b4vgd7",
													},
													Name: "End Event 7 None",
												},
												Incoming: []string{"Flow_07zdeu5", "Flow_0qd2cjd", "Flow_1ef3aqe"},
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
														ID: "Event_0sedygx",
													},
													Name: "End Event 11 Escalation",
												},
												Incoming: []string{"Flow_07hgdqa", "Flow_0dh0nnk", "Flow_1tqpq63"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											EscalationEventDefinitions: []instance.EscalationEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "EscalationEventDefinition_0frigdc",
															},
														},
													},
												},
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
														ID: "Event_0kp736q",
													},
													Name: "End Event 14",
												},
												Incoming: []string{"Flow_0r5tyx6"},
											},
										},
									},
								},
							},
							ReceiveTasks: []instance.ReceiveTask{
								{
									Task: instance.Task{
										Activity: instance.Activity{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Activity_1i3ap44",
													},
													Name: "Receive Task 16",
												},
												Incoming: []string{"Flow_0rx4vqb"},
												Outgoing: []string{"Flow_0qvdwb3"},
											},
										},
									},
								},
							},
							ServiceTasks: []instance.ServiceTask{
								{
									Task: instance.Task{
										Activity: instance.Activity{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Activity_1qy7xgp",
													},
													Name: "Service Task 15",
												},
												Incoming: []string{"Flow_0npffus"},
												Outgoing: []string{"Flow_0rx4vqb"},
											},
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
													ID: "Gateway_0ynnz6a",
												},
												Name: "Parallel Gateway 5",
											},
											Incoming: []string{"Flow_0jmy9bt"},
											Outgoing: []string{"Flow_163ttxe", "Flow_0fop3s1"},
										},
									},
								},
							},
							InclusiveGatewaies: []instance.InclusiveGateway{
								{
									Gateway: instance.Gateway{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Gateway_0gp6gne",
												},
												Name: "Inclusive Gateway",
											},
											Incoming: []string{"Flow_11zr1a3", "Flow_00uq1g7"},
											Outgoing: []string{"Flow_0ax4o0j"},
										},
									},
								},
							},
						},
					},
				},
				Categories: []instance.Category{
					{
						RootElement: instance.RootElement{
							BaseElement: instance.BaseElement{
								ID: "Category_10r7ibr",
							},
						},
						CategoryValues: []instance.CategoryValue{
							{
								BaseElement: instance.BaseElement{
									ID: "CategoryValue_02um9y9",
								},
								Value: "Group",
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
