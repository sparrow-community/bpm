package bpmn

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sparrow-community/bpm/model/bpmn/instance"
)

func TestC_6_0_export(t *testing.T) {
	// create test use ./test/C.6.0-export.bpmn
	path := "./test/C.6.0-export.bpmn"
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
			ID:              "Definitions_13oah7q",
			TargetNamespace: "http://bpmn.io/schema/bpmn",
			Exporter:        "Camunda Modeler",
			ExporterVersion: "5.2.0",
			RootElemnts: instance.RootElemnts{
				Processes: []instance.Process{
					{
						CallableElement: instance.CallableElement{
							RootElement: instance.RootElement{
								BaseElement: instance.BaseElement{
									ID: "Process_19noqni",
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
														ID: "StartEvent_1",
													},
													Name: "Receive Travel Request",
												},
												Outgoing: []string{"Flow_1xi649k"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											MessageEventDefinitions: []instance.MessageEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "MessageEventDefinition_0ein9t6",
															},
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
											ID: "Flow_1xi649k",
										},
									},
									SourceRef: "StartEvent_1",
									TargetRef: "Activity_1qdxrgj",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1p58yis",
										},
									},
									SourceRef: "Activity_1qdxrgj",
									TargetRef: "Gateway_1ersh6n",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_128dw9w",
										},
									},
									SourceRef: "Gateway_1ersh6n",
									TargetRef: "Event_0w821nf",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_15l6xor",
										},
									},
									SourceRef: "Gateway_1ersh6n",
									TargetRef: "Event_1gu9t77",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1a3h5v6",
										},
									},
									SourceRef: "Gateway_1ersh6n",
									TargetRef: "Event_19meht8",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1uw2kdx",
										},
									},
									SourceRef: "Event_0w821nf",
									TargetRef: "Activity_1bidfcm",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1dytiuq",
										},
									},
									SourceRef: "Event_1gu9t77",
									TargetRef: "Activity_1g04x9h",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1domhhx",
										},
										Name: "24 Hours",
									},
									SourceRef: "Event_0isfp1w",
									TargetRef: "Activity_1g04x9h",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0gx0bqr",
										},
									},
									SourceRef: "Event_19meht8",
									TargetRef: "Activity_084p2mw",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0yqzl55",
										},
									},
									SourceRef: "Activity_1g04x9h",
									TargetRef: "Event_0lsnfz2",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1aavu4q",
										},
									},
									SourceRef: "Activity_084p2mw",
									TargetRef: "Event_1yyb73n",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_075xnug",
										},
									},
									SourceRef: "Activity_1bidfcm",
									TargetRef: "Activity_0p5xveb",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0ufn18k",
										},
									},
									SourceRef: "Event_0gv16hd",
									TargetRef: "Activity_05lmtrf",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0vi45od",
										},
									},
									SourceRef: "Activity_0p5xveb",
									TargetRef: "Activity_039ic8d",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1p554j5",
										},
									},
									SourceRef: "Event_0wnb2z5",
									TargetRef: "Event_0qemotd",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1y9j47t",
										},
									},
									SourceRef: "Event_0qemotd",
									TargetRef: "Activity_13l7j32",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0ld9uyo",
										},
									},
									SourceRef: "Activity_039ic8d",
									TargetRef: "Activity_1a4r8ya",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1fz8fw8",
										},
									},
									SourceRef: "Activity_13l7j32",
									TargetRef: "Event_0aabyn5",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1wfqz94",
										},
									},
									SourceRef: "Activity_1a4r8ya",
									TargetRef: "Event_0l6vh4o",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0d43f6n",
										},
									},
									SourceRef: "Activity_05lmtrf",
									TargetRef: "Event_0rk6pwf",
								},
							},
							SendTasks: []instance.SendTask{
								{
									Task: instance.Task{
										Activity: instance.Activity{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Activity_1qdxrgj",
													},
													Name: "Make Flights and Hotel Offer",
												},
												Incoming: []string{"Flow_1xi649k"},
												Outgoing: []string{"Flow_1p58yis"},
											},
											Properties: []instance.Property{
												{
													BaseElement: instance.BaseElement{
														ID: "Property_0jmwx6a",
													},
													Name: "__targetRef_placeholder",
												},
											},
											DataInputAssociations: []instance.DataInputAssociation{
												{
													DataAssociation: instance.DataAssociation{
														BaseElement: instance.BaseElement{
															ID: "DataInputAssociation_17xmvz3",
														},
														SourceRef: "DataObjectReference_0vc6qe4",
														TargetRef: "Property_0jmwx6a",
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
														ID: "Activity_1bidfcm",
													},
													Name: "Request Credit Card Information",
												},
												Incoming: []string{"Flow_1uw2kdx"},
												Outgoing: []string{"Flow_075xnug"},
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
														ID: "Activity_1g04x9h",
													},
													Name: "Notify Customer Offer Expired",
												},
												Incoming: []string{"Flow_1dytiuq", "Flow_1domhhx"},
												Outgoing: []string{"Flow_0yqzl55"},
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
														ID: "Activity_05lmtrf",
													},
													Name: "Notify Failed Booking",
												},
												Incoming: []string{"Flow_0ufn18k"},
												Outgoing: []string{"Flow_0d43f6n"},
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
														ID: "Activity_13l7j32",
													},
													Name: "Notify Failed Credit Transaction",
												},
												Incoming: []string{"Flow_1y9j47t"},
												Outgoing: []string{"Flow_1fz8fw8"},
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
														ID: "Activity_1a4r8ya",
													},
													Name: "Confirm Booking",
												},
												Incoming: []string{"Flow_0ld9uyo"},
												Outgoing: []string{"Flow_1wfqz94"},
											},
											DataOutputAssociations: []instance.DataOutputAssociation{
												{
													DataAssociation: instance.DataAssociation{
														BaseElement: instance.BaseElement{
															ID: "DataOutputAssociation_0yhidap",
														},
														TargetRef: "DataObjectReference_0hy72ip",
													},
												},
											},
										},
									},
								},
							},
							DataObjectReferenes: []instance.DataObjectReference{
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "DataObjectReference_0vc6qe4",
										},
										Name: "Travel Request",
									},
									DataObjectRef: "DataObject_0acgqpe",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "DataObjectReference_0hy72ip",
										},
										Name: "Itinerary",
									},
									DataObjectRef: "DataObject_1h6jn3y",
								},
							},
							DataObjects: []instance.DataObject{
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "DataObject_0acgqpe",
										},
									},
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "DataObject_1h6jn3y",
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
													ID: "Gateway_1ersh6n",
												},
											},
											Incoming: []string{"Flow_1p58yis"},
											Outgoing: []string{"Flow_128dw9w", "Flow_15l6xor", "Flow_1a3h5v6"},
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
														ID: "Event_0w821nf",
													},
													Name: "Offer Approved",
												},
												Incoming: []string{"Flow_128dw9w"},
												Outgoing: []string{"Flow_1uw2kdx"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											MessageEventDefinitions: []instance.MessageEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "MessageEventDefinition_12s4ok2",
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
														ID: "Event_1gu9t77",
													},
													Name: "24 Hours",
												},
												Incoming: []string{"Flow_15l6xor"},
												Outgoing: []string{"Flow_1dytiuq"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											TimerEventDefinitions: []instance.TimerEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "TimerEventDefinition_0gimu3s",
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
														ID: "Event_19meht8",
													},
													Name: "Cancel Request",
												},
												Incoming: []string{"Flow_1a3h5v6"},
												Outgoing: []string{"Flow_0gx0bqr"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											MessageEventDefinitions: []instance.MessageEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "MessageEventDefinition_0ylxz6n",
															},
														},
													},
												},
											},
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
														ID: "Event_0qemotd",
													},
													Name: "Booking",
												},
												Incoming: []string{"Flow_1p554j5"},
												Outgoing: []string{"Flow_1y9j47t"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											CompensateEventDefinitions: []instance.CompensateEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "CompensateEventDefinition_1nra8kx",
															},
														},
													},
												},
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
														ID: "Event_0isfp1w",
													},
												},
												Outgoing: []string{"Flow_1domhhx"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											TimerEventDefinitions: []instance.TimerEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "TimerEventDefinition_0ry7jh8",
															},
														},
													},
												},
											},
										},
									},
									AttachedToRef: "Activity_1bidfcm",
								},
								{
									CatchEvent: instance.CatchEvent{
										Event: instance.Event{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Event_0gv16hd",
													},
												},
												Outgoing: []string{"Flow_0ufn18k"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											ErrorEventDefinitions: []instance.ErrorEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "ErrorEventDefinition_10wgwwl",
															},
														},
													},
												},
											},
										},
									},
									AttachedToRef: "Activity_0p5xveb",
								},
								{
									CatchEvent: instance.CatchEvent{
										Event: instance.Event{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Event_0wnb2z5",
													},
												},
												Outgoing: []string{"Flow_1p554j5"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											ErrorEventDefinitions: []instance.ErrorEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "ErrorEventDefinition_1tcdwud",
															},
														},
													},
												},
											},
										},
									},
									AttachedToRef: "Activity_039ic8d",
								},
							},
							ServiceTasks: []instance.ServiceTask{
								{
									Task: instance.Task{
										Activity: instance.Activity{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Activity_084p2mw",
													},
													Name: "Update Customer Record",
												},
												Incoming: []string{"Flow_0gx0bqr"},
												Outgoing: []string{"Flow_1aavu4q"},
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
														ID: "Activity_039ic8d",
													},
													Name: "Charge Credit Card",
												},
												Incoming: []string{"Flow_0vi45od"},
												Outgoing: []string{"Flow_0ld9uyo"},
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
														ID: "Event_0lsnfz2",
													},
													Name: "Offer Expired",
												},
												Incoming: []string{"Flow_0yqzl55"},
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
														ID: "Event_1yyb73n",
													},
													Name: "Request Cancelled",
												},
												Incoming: []string{"Flow_1aavu4q"},
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
														ID: "Event_0aabyn5",
													},
													Name: "Failed Credit Transaction",
												},
												Incoming: []string{"Flow_1fz8fw8"},
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
														ID: "Event_0l6vh4o",
													},
													Name: "Booking Confirmed",
												},
												Incoming: []string{"Flow_1wfqz94"},
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
														ID: "Event_0rk6pwf",
													},
													Name: "Failed Booking",
												},
												Incoming: []string{"Flow_0d43f6n"},
											},
										},
									},
								},
							},
							SubProcesses: []instance.SubProcess{
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Activity_0p5xveb",
												},
												Name: "Make Booking",
											},
											Incoming: []string{"Flow_075xnug"},
											Outgoing: []string{"Flow_0vi45od"},
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
																	ID: "Event_1fywxat",
																},
															},
															Outgoing: []string{"Flow_0wb651w"},
														},
													},
												},
											},
										},
										SequenceFlows: []instance.SequenceFlow{
											{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Flow_0wb651w",
													},
												},
												SourceRef: "Event_1fywxat",
												TargetRef: "Gateway_14mg5pf",
											},
											{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Flow_01h7e6d",
													},
												},
												SourceRef: "Gateway_14mg5pf",
												TargetRef: "Activity_13sg203",
											},
											{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Flow_04o2v8w",
													},
												},
												SourceRef: "Activity_13sg203",
												TargetRef: "Gateway_1mo08sa",
											},
											{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Flow_1h0awz9",
													},
												},
												SourceRef: "Gateway_14mg5pf",
												TargetRef: "Activity_0qz49yv",
											},
											{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Flow_1l3h2oz",
													},
												},
												SourceRef: "Activity_0qz49yv",
												TargetRef: "Gateway_1mo08sa",
											},
											{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Flow_0stizar",
													},
												},
												SourceRef: "Gateway_1mo08sa",
												TargetRef: "Event_1c52ias",
											},
										},
										ParallelGatewaies: []instance.ParallelGateway{
											{
												Gateway: instance.Gateway{
													FlowNode: instance.FlowNode{
														FlowElement: instance.FlowElement{
															BaseElement: instance.BaseElement{
																ID: "Gateway_14mg5pf",
															},
														},
														Incoming: []string{"Flow_0wb651w"},
														Outgoing: []string{"Flow_01h7e6d", "Flow_1h0awz9"},
													},
												},
											},
											{
												Gateway: instance.Gateway{
													FlowNode: instance.FlowNode{
														FlowElement: instance.FlowElement{
															BaseElement: instance.BaseElement{
																ID: "Gateway_1mo08sa",
															},
														},
														Incoming: []string{"Flow_04o2v8w", "Flow_1l3h2oz"},
														Outgoing: []string{"Flow_0stizar"},
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
																	ID: "Activity_13sg203",
																},
																Name: "Book Flight",
															},
															Incoming: []string{"Flow_01h7e6d"},
															Outgoing: []string{"Flow_04o2v8w"},
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
																	ID: "Activity_0hgj2bs",
																},
																Name: "Cancel Flight",
															},
														},
														IsForCompensation: true,
													},
												},
											},
											{
												Task: instance.Task{
													Activity: instance.Activity{
														FlowNode: instance.FlowNode{
															FlowElement: instance.FlowElement{
																BaseElement: instance.BaseElement{
																	ID: "Activity_0qz49yv",
																},
																Name: "Book Hotel",
															},
															Incoming: []string{"Flow_1h0awz9"},
															Outgoing: []string{"Flow_1l3h2oz"},
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
																	ID: "Activity_1n0lwxw",
																},
																Name: "Cancel Hotel",
															},
														},
														IsForCompensation: true,
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
																	ID: "Event_0oxjqip",
																},
																Name: "Flight",
															},
														},
													},
													EventDefinitions: instance.EventDefinitions{
														CompensateEventDefinitions: []instance.CompensateEventDefinition{
															{
																EventDefinition: instance.EventDefinition{
																	RootElement: instance.RootElement{
																		BaseElement: instance.BaseElement{
																			ID: "CompensateEventDefinition_0l7f698",
																		},
																	},
																},
															},
														},
													},
												},
												AttachedToRef: "Activity_13sg203",
											},
											{
												CatchEvent: instance.CatchEvent{
													Event: instance.Event{
														FlowNode: instance.FlowNode{
															FlowElement: instance.FlowElement{
																BaseElement: instance.BaseElement{
																	ID: "Event_1gsyz0h",
																},
																Name: "Hotel",
															},
														},
													},
													EventDefinitions: instance.EventDefinitions{
														CompensateEventDefinitions: []instance.CompensateEventDefinition{
															{
																EventDefinition: instance.EventDefinition{
																	RootElement: instance.RootElement{
																		BaseElement: instance.BaseElement{
																			ID: "CompensateEventDefinition_11wluxa",
																		},
																	},
																},
															},
														},
													},
												},
												AttachedToRef: "Activity_0qz49yv",
											},
										},
										EndEvents: []instance.EndEvent{
											{
												ThrowEvent: instance.ThrowEvent{
													Event: instance.Event{
														FlowNode: instance.FlowNode{
															FlowElement: instance.FlowElement{
																BaseElement: instance.BaseElement{
																	ID: "Event_1c52ias",
																},
																Name: "Travel Booked",
															},
															Incoming: []string{"Flow_0stizar"},
														},
													},
												},
											},
										},
										SubProcesses: []instance.SubProcess{
											{
												Activity: instance.Activity{
													FlowNode: instance.FlowNode{
														FlowElement: instance.FlowElement{
															BaseElement: instance.BaseElement{
																ID: "Activity_1t020b1",
															},
														},
													},
												},
												TriggeredByEvent: true,
												FlowElements: instance.FlowElements{
													StartEvents: []instance.StartEvent{
														{
															CatchEvent: instance.CatchEvent{
																Event: instance.Event{
																	FlowNode: instance.FlowNode{
																		FlowElement: instance.FlowElement{
																			BaseElement: instance.BaseElement{
																				ID: "Event_0hlskm4",
																			},
																		},
																		Outgoing: []string{"Flow_01qcqu0"},
																	},
																},
																EventDefinitions: instance.EventDefinitions{
																	CompensateEventDefinitions: []instance.CompensateEventDefinition{
																		{
																			EventDefinition: instance.EventDefinition{
																				RootElement: instance.RootElement{
																					BaseElement: instance.BaseElement{
																						ID: "CompensateEventDefinition_11z58bi",
																					},
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
																	ID: "Flow_01qcqu0",
																},
															},
															SourceRef: "Event_0hlskm4",
															TargetRef: "Gateway_0azwu61",
														},
														{
															FlowElement: instance.FlowElement{
																BaseElement: instance.BaseElement{
																	ID: "Flow_0ck9rq3",
																},
															},
															SourceRef: "Gateway_0azwu61",
															TargetRef: "Event_1o7y58x",
														},
														{
															FlowElement: instance.FlowElement{
																BaseElement: instance.BaseElement{
																	ID: "Flow_168c72o",
																},
															},
															SourceRef: "Gateway_0azwu61",
															TargetRef: "Event_17sn5te",
														},
														{
															FlowElement: instance.FlowElement{
																BaseElement: instance.BaseElement{
																	ID: "Flow_1f1d6rn",
																},
															},
															SourceRef: "Event_17sn5te",
															TargetRef: "Gateway_1udyyri",
														},
														{
															FlowElement: instance.FlowElement{
																BaseElement: instance.BaseElement{
																	ID: "Flow_170itut",
																},
															},
															SourceRef: "Event_1o7y58x",
															TargetRef: "Gateway_1udyyri",
														},
														{
															FlowElement: instance.FlowElement{
																BaseElement: instance.BaseElement{
																	ID: "Flow_0hvvmqj",
																},
															},
															SourceRef: "Gateway_1udyyri",
															TargetRef: "Event_0bfvt7c",
														},
													},
													ParallelGatewaies: []instance.ParallelGateway{
														{
															Gateway: instance.Gateway{
																FlowNode: instance.FlowNode{
																	FlowElement: instance.FlowElement{
																		BaseElement: instance.BaseElement{
																			ID: "Gateway_0azwu61",
																		},
																	},
																	Incoming: []string{"Flow_01qcqu0"},
																	Outgoing: []string{"Flow_0ck9rq3", "Flow_168c72o"},
																},
															},
														},
														{
															Gateway: instance.Gateway{
																FlowNode: instance.FlowNode{
																	FlowElement: instance.FlowElement{
																		BaseElement: instance.BaseElement{
																			ID: "Gateway_1udyyri",
																		},
																	},
																	Incoming: []string{"Flow_1f1d6rn", "Flow_170itut"},
																	Outgoing: []string{"Flow_0hvvmqj"},
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
																				ID: "Event_1o7y58x",
																			},
																			Name: "Flight",
																		},
																		Incoming: []string{"Flow_0ck9rq3"},
																		Outgoing: []string{"Flow_170itut"},
																	},
																},
																EventDefinitions: instance.EventDefinitions{
																	CompensateEventDefinitions: []instance.CompensateEventDefinition{
																		{
																			EventDefinition: instance.EventDefinition{
																				RootElement: instance.RootElement{
																					BaseElement: instance.BaseElement{
																						ID: "CompensateEventDefinition_1wju1u7",
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
																				ID: "Event_17sn5te",
																			},
																			Name: "Hotel",
																		},
																		Incoming: []string{"Flow_168c72o"},
																		Outgoing: []string{"Flow_1f1d6rn"},
																	},
																},
																EventDefinitions: instance.EventDefinitions{
																	CompensateEventDefinitions: []instance.CompensateEventDefinition{
																		{
																			EventDefinition: instance.EventDefinition{
																				RootElement: instance.RootElement{
																					BaseElement: instance.BaseElement{
																						ID: "CompensateEventDefinition_08lflw7",
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
																				ID: "Event_0bfvt7c",
																			},
																		},
																		Incoming: []string{"Flow_0hvvmqj"},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
									Artifacts: instance.Artifacts{
										Associations: []instance.Association{
											{
												Artifact: instance.Artifact{
													BaseElement: instance.BaseElement{
														ID: "Association_03vh68e",
													},
												},
												AssociationDirection: "One",
												SourceRef:            "Event_0oxjqip",
												TargetRef:            "Activity_0hgj2bs",
											},
											{
												Artifact: instance.Artifact{
													BaseElement: instance.BaseElement{
														ID: "Association_0srffpa",
													},
												},
												AssociationDirection: "One",
												SourceRef:            "Event_1gsyz0h",
												TargetRef:            "Activity_1n0lwxw",
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

	if diff := cmp.Diff(excepted, modelInstance); diff != "" {
		t.Errorf("Mismatch in %s (-want +got):\n%s", path, diff)
	}

}
