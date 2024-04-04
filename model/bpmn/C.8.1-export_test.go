package bpmn

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sparrow-community/bpm/model/bpmn/instance"
)

func TestC_8_1_export(t *testing.T) {
	// create test use ./test/C.8.1-export.bpmn
	path := "./test/C.8.1-export.bpmn"
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
				Collaborations: []instance.Collaboration{
					{
						RootElement: instance.RootElement{
							BaseElement: instance.BaseElement{
								ID: "Collaboration_05psygt",
							},
						},
						Participants: []instance.Participant{
							{
								BaseElement: instance.BaseElement{
									ID: "Participant_11h1ii2",
								},
								Name:       "Vacation Request",
								ProcessRef: "Process_1xl5gyi",
							},
						},
					},
				},
				Processes: []instance.Process{
					{
						CallableElement: instance.CallableElement{
							RootElement: instance.RootElement{
								BaseElement: instance.BaseElement{
									ID: "Process_1xl5gyi",
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
														ID: "Event_0q4cbsy",
													},
													Name: "Vacation Request Received",
												},
												Outgoing: []string{"Flow_1916suv"},
											},
										},
									},
								},
							},
							SequenceFlows: []instance.SequenceFlow{
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1916suv",
										},
									},
									SourceRef: "Event_0q4cbsy",
									TargetRef: "Activity_0lrykpf",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_05r3ppa",
										},
									},
									SourceRef: "Activity_0lrykpf",
									TargetRef: "Activity_13m56bo",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1229gb2",
										},
									},
									SourceRef: "Activity_13m56bo",
									TargetRef: "Gateway_1m1epyv",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1slje8w",
										},
										Name: "Approved",
									},
									SourceRef: "Gateway_1m1epyv",
									TargetRef: "Activity_0kxqh4y",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_17w7s53",
										},
									},
									SourceRef: "Activity_0kxqh4y",
									TargetRef: "Activity_0h1298a",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_06f2rqz",
										},
										Name: "Manual Validation Required",
									},
									SourceRef: "Gateway_1m1epyv",
									TargetRef: "Activity_14g794r",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1bn59o5",
										},
									},
									SourceRef: "Activity_14g794r",
									TargetRef: "Gateway_0bpdi0b",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_18wfjpz",
										},
										Name: "Approved",
									},
									SourceRef: "Gateway_0bpdi0b",
									TargetRef: "Activity_1exp9va",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1kjzd51",
										},
									},
									SourceRef: "Activity_1exp9va",
									TargetRef: "Activity_0qxgfc9",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1s1ikgt",
										},
									},
									SourceRef: "Activity_0qxgfc9",
									TargetRef: "Event_0fselnt",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0dhwpuk",
										},
									},
									SourceRef: "Activity_0h1298a",
									TargetRef: "Event_0jqpsr9",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_16iauun",
										},
										Name: "Refused",
									},
									SourceRef: "Gateway_0bpdi0b",
									TargetRef: "Activity_02k6h6m",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1uc40yw",
										},
									},
									SourceRef: "Activity_02k6h6m",
									TargetRef: "Event_10bj8rf",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1x3mfjv",
										},
										Name: "Refused",
									},
									SourceRef: "Gateway_1m1epyv",
									TargetRef: "Activity_0wd8gll",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1x4un6h",
										},
									},
									SourceRef: "Activity_0wd8gll",
									TargetRef: "Event_0990itm",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_15v234s",
										},
									},
									SourceRef: "Event_1aoguwi",
									TargetRef: "Event_0c4c8gy",
								},
							},
							ServiceTasks: []instance.ServiceTask{
								{
									Task: instance.Task{
										Activity: instance.Activity{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Activity_0lrykpf",
													},
													Name: "Fetch Vacation Information",
												},
												Incoming: []string{"Flow_1916suv"},
												Outgoing: []string{"Flow_05r3ppa"},
											},
											Properties: []instance.Property{
												{
													BaseElement: instance.BaseElement{
														ID: "Property_1ffrmni",
													},
													Name: "__targetRef_placeholder",
												},
											},
											DataInputAssociations: []instance.DataInputAssociation{
												{
													DataAssociation: instance.DataAssociation{
														BaseElement: instance.BaseElement{
															ID: "DataInputAssociation_19594jt",
														},
														SourceRef: "DataObjectReference_0s5y8ld",
														TargetRef: "Property_1ffrmni",
													},
												},
											},
											DataOutputAssociations: []instance.DataOutputAssociation{
												{
													DataAssociation: instance.DataAssociation{
														BaseElement: instance.BaseElement{
															ID: "DataOutputAssociation_1gvweaq",
														},
														TargetRef: "DataObjectReference_0uia8i9",
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
														ID: "Activity_0h1298a",
													},
													Name: "Update Remaining Vacation",
												},
												Incoming: []string{"Flow_17w7s53"},
												Outgoing: []string{"Flow_0dhwpuk"},
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
														ID: "Activity_0qxgfc9",
													},
													Name: "Update Remaining Vacation",
												},
												Incoming: []string{"Flow_1kjzd51"},
												Outgoing: []string{"Flow_1s1ikgt"},
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
														ID: "Event_1aoguwi",
													},
												},
												Outgoing: []string{"Flow_15v234s"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											ErrorEventDefinitions: []instance.ErrorEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "ErrorEventDefinition_1kchpzk",
															},
														},
													},
												},
											},
										},
									},
									AttachedToRef: "Activity_0lrykpf",
								},
							},
							DataObjectReferenes: []instance.DataObjectReference{
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "DataObjectReference_0s5y8ld",
										},
										Name: "Employee Badge Number",
									},
									DataObjectRef: "DataObject_13kw65u",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "DataObjectReference_0uia8i9",
										},
										Name: "Current Vacation Status",
									},
									DataObjectRef: "DataObject_0ta7zty",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "DataObjectReference_1m7aoov",
										},
										Name: "To",
									},
									DataObjectRef: "DataObject_1grql3s",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "DataObjectReference_06x7a95",
										},
										Name: "From",
									},
									DataObjectRef: "DataObject_1r856yi",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "DataObjectReference_19p7a9l",
										},
										Name: "Vacation Approval",
									},
									DataObjectRef: "DataObject_0msvzum",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "DataObjectReference_0ynczgz",
										},
										Name: "Reason",
									},
									DataObjectRef: "DataObject_0v2smxx",
								},
							},
							DataObjects: []instance.DataObject{
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "DataObject_13kw65u",
										},
									},
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "DataObject_0ta7zty",
										},
									},
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "DataObject_1grql3s",
										},
									},
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "DataObject_1r856yi",
										},
									},
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "DataObject_0msvzum",
										},
									},
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "DataObject_0v2smxx",
										},
									},
								},
							},
							BusinessRuleTasks: []instance.BusinessRuleTask{
								{
									Task: instance.Task{
										Activity: instance.Activity{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Activity_13m56bo",
													},
													Name: "Vacation Approval",
												},
												Incoming: []string{"Flow_05r3ppa"},
												Outgoing: []string{"Flow_1229gb2"},
											},
											Properties: []instance.Property{
												{
													BaseElement: instance.BaseElement{
														ID: "Property_1c2el8a",
													},
													Name: "__targetRef_placeholder",
												},
											},
											DataInputAssociations: []instance.DataInputAssociation{
												{
													DataAssociation: instance.DataAssociation{
														BaseElement: instance.BaseElement{
															ID: "DataInputAssociation_03oginu",
														},
														SourceRef: "DataObjectReference_1m7aoov",
														TargetRef: "Property_1c2el8a",
													},
												},
												{
													DataAssociation: instance.DataAssociation{
														BaseElement: instance.BaseElement{
															ID: "DataInputAssociation_1j9yja8",
														},
														SourceRef: "DataObjectReference_06x7a95",
														TargetRef: "Property_1c2el8a",
													},
												},
											},
											DataOutputAssociations: []instance.DataOutputAssociation{
												{
													DataAssociation: instance.DataAssociation{
														BaseElement: instance.BaseElement{
															ID: "DataOutputAssociation_06bl88p",
														},
														TargetRef: "DataObjectReference_19p7a9l",
													},
												},
												{
													DataAssociation: instance.DataAssociation{
														BaseElement: instance.BaseElement{
															ID: "DataOutputAssociation_0jmelr8",
														},
														TargetRef: "DataObjectReference_0ynczgz",
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
													ID: "Gateway_1m1epyv",
												},
												Name: "",
											},
											Incoming: []string{"Flow_1229gb2"},
											Outgoing: []string{"Flow_1slje8w", "Flow_06f2rqz", "Flow_1x3mfjv"},
										},
									},
									Default: "Flow_1x3mfjv",
								},
								{
									Gateway: instance.Gateway{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Gateway_0bpdi0b",
												},
												Name: "",
											},
											Incoming: []string{"Flow_1bn59o5"},
											Outgoing: []string{"Flow_18wfjpz", "Flow_16iauun"},
										},
									},
									Default: "Flow_16iauun",
								},
							},
							SendTasks: []instance.SendTask{
								{
									Task: instance.Task{
										Activity: instance.Activity{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "Activity_0kxqh4y",
													},
													Name: "Notify Employee of Approval",
												},
												Incoming: []string{"Flow_1slje8w"},
												Outgoing: []string{"Flow_17w7s53"},
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
														ID: "Activity_1exp9va",
													},
													Name: "Notify Employee of Approval",
												},
												Incoming: []string{"Flow_18wfjpz"},
												Outgoing: []string{"Flow_1kjzd51"},
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
														ID: "Activity_02k6h6m",
													},
													Name: "Notify Employee of Refusal",
												},
												Incoming: []string{"Flow_16iauun"},
												Outgoing: []string{"Flow_1uc40yw"},
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
														ID: "Activity_0wd8gll",
													},
													Name: "Notify Employee of Refusal",
												},
												Incoming: []string{"Flow_1x3mfjv"},
												Outgoing: []string{"Flow_1x4un6h"},
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
														ID: "Activity_14g794r",
													},
													Name: "Manually Approve Vacation",
												},
												Incoming: []string{"Flow_06f2rqz"},
												Outgoing: []string{"Flow_1bn59o5"},
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
														ID: "Event_0fselnt",
													},
													Name: "Vacation Approved by Manager",
												},
												Incoming: []string{"Flow_1s1ikgt"},
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
														ID: "Event_0jqpsr9",
													},
													Name: "Vacation Approved Automatically",
												},
												Incoming: []string{"Flow_0dhwpuk"},
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
														ID: "Event_10bj8rf",
													},
													Name: "Vacation Refused by Manager",
												},
												Incoming: []string{"Flow_1uc40yw"},
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
														ID: "Event_0990itm",
													},
													Name: "Vacation Refused Automatically",
												},
												Incoming: []string{"Flow_1x4un6h"},
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
														ID: "Event_0c4c8gy",
													},
													Name: "Employee not found",
												},
												Incoming: []string{"Flow_15v234s"},
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
