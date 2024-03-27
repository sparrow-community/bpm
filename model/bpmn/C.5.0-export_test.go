package bpmn

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sparrow-community/bpm/model/bpmn/instance"
)

func TestC_5_0_export(t *testing.T) {
	// create test use ./test/C.5.0-export.bpmn
	path := "./test/C.5.0-export.bpmn"
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
			ID:              "Definitions_09f86rk",
			TargetNamespace: "http://bpmn.io/schema/bpmn",
			Exporter:        "Camunda Modeler",
			ExporterVersion: "5.2.0",
			RootElemnts: instance.RootElemnts{
				Collaborations: []instance.Collaboration{
					{
						RootElement: instance.RootElement{
							BaseElement: instance.BaseElement{
								ID: "Collaboration_10f1am5",
							},
						},
						Participants: []instance.Participant{
							{
								BaseElement: instance.BaseElement{
									ID: "Participant_0dcjjya",
								},
								Name:       "Bank",
								ProcessRef: "Process_18ixeuz",
							},
						},
					},
				},
				Processes: []instance.Process{
					{
						CallableElement: instance.CallableElement{
							RootElement: instance.RootElement{
								BaseElement: instance.BaseElement{
									ID: "Process_18ixeuz",
								},
							},
						},
						IsExecutable: true,
						LaneSets: []instance.LaneSet{
							{
								BaseElement: instance.BaseElement{
									ID: "LaneSet_0b0ox2x",
								},
								Lanes: []instance.Lane{
									{
										BaseElement: instance.BaseElement{
											ID: "Lane_0zg7l9h",
										},
										Name: "Head of Market Service",
										FlowNodeRefs: []string{
											"Activity_1exyjv9",
											"Gateway_1kweljx",
											"Activity_0gwbifa",
											"Event_150agrk",
										},
									},
									{
										BaseElement: instance.BaseElement{
											ID: "Lane_1tifwj0",
										},
										Name: "Corporate Account Manager",
										FlowNodeRefs: []string{
											"Activity_1qr8zif",
											"Gateway_0bfi7td",
											"Activity_1jbdjra",
											"Event_1k9lc9s",
										},
									},
									{
										BaseElement: instance.BaseElement{
											ID: "Lane_0fioqbn",
										},
										Name: "Private Customer Account Manager",
										FlowNodeRefs: []string{
											"StartEvent_1",
											"Activity_0qyqx4q",
											"Activity_02iaxsg",
											"Gateway_0w4yy7d",
											"Gateway_0nz6kpd",
											"Activity_0sffmrs",
											"Activity_059ixil",
											"Activity_1nw3l11",
											"Gateway_0olrjos",
											"Activity_09zja0c",
											"Activity_12aarlh",
											"Gateway_0x30xhl",
											"Activity_11byy83",
											"Gateway_0g55arm",
											"Gateway_1v08mnm",
											"Activity_1kfeipf",
											"Activity_10xe0k2",
											"Activity_0miaffi",
											"Event_1e3rnn0",
											"Activity_0mwrxc9",
											"Gateway_1dzkbyo",
											"Gateway_0qgud4j",
											"Activity_0qi7bn3",
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
														ID: "StartEvent_1",
													},
													Name: "Customer interested in Bank offer",
												},
												Outgoing: []string{"Flow_1axad4z"},
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
														ID: "Activity_0qyqx4q",
													},
													Name: "Interview customer",
												},
												Incoming: []string{"Flow_1axad4z"},
												Outgoing: []string{"Flow_0xot8mo"},
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
														ID: "Activity_02iaxsg",
													},
													Name: "Prove/Provide identity",
												},
												Incoming: []string{"Flow_0xot8mo"},
												Outgoing: []string{"Flow_1tuddf9"},
											},
											DataOutputAssociations: []instance.DataOutputAssociation{
												{
													DataAssociation: instance.DataAssociation{
														BaseElement: instance.BaseElement{
															ID: "DataOutputAssociation_0lzuywp",
														},
														TargetRef: "DataObjectReference_1hls48h",
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
														ID: "Activity_1qr8zif",
													},
													Name: "Document the identity of the economic owner",
												},
												Incoming: []string{"Flow_0wsn5cd"},
												Outgoing: []string{"Flow_0068ri5"},
											},
											Properties: []instance.Property{
												{
													BaseElement: instance.BaseElement{
														ID: "Property_00sy1c4",
													},
													Name: "__targetRef_placeholder",
												},
											},
											DataInputAssociations: []instance.DataInputAssociation{
												{
													DataAssociation: instance.DataAssociation{
														BaseElement: instance.BaseElement{
															ID: "DataInputAssociation_162nud5",
														},
														SourceRef: "DataObjectReference_1hls48h",
														TargetRef: "Property_00sy1c4",
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
														ID: "Activity_1jbdjra",
													},
													Name: "End business relation",
												},
												Incoming: []string{"Flow_1rojgff"},
												Outgoing: []string{"Flow_0i1ef99"},
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
														ID: "Activity_0sffmrs",
													},
													Name: "Obtain supporting data and documents of the customer",
												},
												Incoming: []string{"Flow_07rh70d"},
												Outgoing: []string{"Flow_1kmt6vy"},
											},
											Properties: []instance.Property{
												{
													BaseElement: instance.BaseElement{
														ID: "Property_0kqr8wd",
													},
													Name: "__targetRef_placeholder",
												},
											},
											DataInputAssociations: []instance.DataInputAssociation{
												{
													DataAssociation: instance.DataAssociation{
														BaseElement: instance.BaseElement{
															ID: "DataInputAssociation_1mzopaj",
														},
														SourceRef: "DataObjectReference_1hls48h",
														TargetRef: "Property_0kqr8wd",
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
														ID: "Activity_059ixil",
													},
													Name: "Check customer documents",
												},
												Incoming: []string{"Flow_1kmt6vy"},
												Outgoing: []string{"Flow_1bsig71"},
											},
											Properties: []instance.Property{
												{
													BaseElement: instance.BaseElement{
														ID: "Property_0p17pfd",
													},
													Name: "__targetRef_placeholder",
												},
											},
											DataInputAssociations: []instance.DataInputAssociation{
												{
													DataAssociation: instance.DataAssociation{
														BaseElement: instance.BaseElement{
															ID: "DataInputAssociation_1kz9ntg",
														},
														SourceRef: "DataObjectReference_1hls48h",
														TargetRef: "Property_0p17pfd",
													},
												},
											},
											DataOutputAssociations: []instance.DataOutputAssociation{
												{
													DataAssociation: instance.DataAssociation{
														BaseElement: instance.BaseElement{
															ID: "DataOutputAssociation_0hnkqhf",
														},
														TargetRef: "DataObjectReference_0orm9f6",
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
														ID: "Activity_1nw3l11",
													},
													Name: "Find documents in customer file",
												},
												Incoming: []string{"Flow_1c2yylu"},
												Outgoing: []string{"Flow_0w9hvzb"},
											},
											Properties: []instance.Property{
												{
													BaseElement: instance.BaseElement{
														ID: "Property_09oqjh7",
													},
													Name: "__targetRef_placeholder",
												},
											},
											DataInputAssociations: []instance.DataInputAssociation{
												{
													DataAssociation: instance.DataAssociation{
														BaseElement: instance.BaseElement{
															ID: "DataInputAssociation_0j754wv",
														},
														SourceRef: "DataObjectReference_04bhv2e",
														TargetRef: "Property_09oqjh7",
													},
												},
											},
											DataOutputAssociations: []instance.DataOutputAssociation{
												{
													DataAssociation: instance.DataAssociation{
														BaseElement: instance.BaseElement{
															ID: "DataOutputAssociation_1b3koo9",
														},
														TargetRef: "DataStoreReference_06oj71q",
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
														ID: "Activity_09zja0c",
													},
													Name: "Add personal data",
												},
												Incoming: []string{"Flow_0yqxpk3"},
												Outgoing: []string{"Flow_1uklf5b"},
											},
											DataOutputAssociations: []instance.DataOutputAssociation{
												{
													DataAssociation: instance.DataAssociation{
														BaseElement: instance.BaseElement{
															ID: "DataOutputAssociation_0zvt05l",
														},
														TargetRef: "DataStoreReference_06oj71q",
													},
												},
												{
													DataAssociation: instance.DataAssociation{
														BaseElement: instance.BaseElement{
															ID: "DataOutputAssociation_02hqoca",
														},
														TargetRef: "DataObjectReference_1k37h1h",
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
														ID: "Activity_12aarlh",
													},
													Name: "Perform know your customer (KYC) activities",
												},
												Incoming: []string{"Flow_1bicxco"},
												Outgoing: []string{"Flow_00x75yi"},
											},
											Properties: []instance.Property{
												{
													BaseElement: instance.BaseElement{
														ID: "Property_19yj73k",
													},
													Name: "__targetRef_placeholder",
												},
											},
											DataInputAssociations: []instance.DataInputAssociation{
												{
													DataAssociation: instance.DataAssociation{
														BaseElement: instance.BaseElement{
															ID: "DataInputAssociation_1gxtza7",
														},
														SourceRef: "DataObjectReference_1k37h1h",
														TargetRef: "Property_19yj73k",
													},
												},
											},
											DataOutputAssociations: []instance.DataOutputAssociation{
												{
													DataAssociation: instance.DataAssociation{
														BaseElement: instance.BaseElement{
															ID: "DataOutputAssociation_1swc9uo",
														},
														TargetRef: "DataStoreReference_06oj71q",
													},
												},
												{
													DataAssociation: instance.DataAssociation{
														BaseElement: instance.BaseElement{
															ID: "DataOutputAssociation_0t3m61i",
														},
														TargetRef: "DataObjectReference_134mvuu",
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
														ID: "Activity_11byy83",
													},
													Name: "Perform risk assessment of the customer",
												},
												Incoming: []string{"Flow_0dn002z"},
												Outgoing: []string{"Flow_0d94m6a"},
											},
											Properties: []instance.Property{
												{
													BaseElement: instance.BaseElement{
														ID: "Property_19jaan0",
													},
													Name: "__targetRef_placeholder",
												},
											},
											DataInputAssociations: []instance.DataInputAssociation{
												{
													DataAssociation: instance.DataAssociation{
														BaseElement: instance.BaseElement{
															ID: "DataInputAssociation_1hyzf08",
														},
														SourceRef: "DataObjectReference_134mvuu",
														TargetRef: "Property_19jaan0",
													},
												},
											},
											DataOutputAssociations: []instance.DataOutputAssociation{
												{
													DataAssociation: instance.DataAssociation{
														BaseElement: instance.BaseElement{
															ID: "DataOutputAssociation_0105tem",
														},
														TargetRef: "DataStoreReference_06oj71q",
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
														ID: "Activity_1exyjv9",
													},
													Name: "Check risk and decide about approval",
												},
												Incoming: []string{"Flow_04jrss9"},
												Outgoing: []string{"Flow_1tqxpdq"},
											},
											Properties: []instance.Property{
												{
													BaseElement: instance.BaseElement{
														ID: "Property_1xasvkv",
													},
													Name: "__targetRef_placeholder",
												},
											},
											DataInputAssociations: []instance.DataInputAssociation{
												{
													DataAssociation: instance.DataAssociation{
														BaseElement: instance.BaseElement{
															ID: "DataInputAssociation_03zkllz",
														},
														SourceRef: "DataObjectReference_134mvuu",
														TargetRef: "Property_1xasvkv",
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
														ID: "Activity_1kfeipf",
													},
													Name: "Document risk assessment",
												},
												Incoming: []string{"Flow_1u4qwnh"},
												Outgoing: []string{"Flow_0jhqlw7"},
											},
											Properties: []instance.Property{
												{
													BaseElement: instance.BaseElement{
														ID: "Property_0tms0mp",
													},
													Name: "__targetRef_placeholder",
												},
											},
											DataInputAssociations: []instance.DataInputAssociation{
												{
													DataAssociation: instance.DataAssociation{
														BaseElement: instance.BaseElement{
															ID: "DataInputAssociation_11xrg0f",
														},
														SourceRef: "DataObjectReference_134mvuu",
														TargetRef: "Property_0tms0mp",
													},
												},
											},
											DataOutputAssociations: []instance.DataOutputAssociation{
												{
													DataAssociation: instance.DataAssociation{
														BaseElement: instance.BaseElement{
															ID: "DataOutputAssociation_12u836g",
														},
														TargetRef: "DataStoreReference_06oj71q",
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
														ID: "Activity_0miaffi",
													},
													Name: "Create customer in the system",
												},
												Incoming: []string{"Flow_0p0czrh"},
												Outgoing: []string{"Flow_0i620wv"},
											},
											Properties: []instance.Property{
												{
													BaseElement: instance.BaseElement{
														ID: "Property_0sk0tl4",
													},
													Name: "__targetRef_placeholder",
												},
											},
											DataInputAssociations: []instance.DataInputAssociation{
												{
													DataAssociation: instance.DataAssociation{
														BaseElement: instance.BaseElement{
															ID: "DataInputAssociation_0c1plix",
														},
														SourceRef: "DataObjectReference_08zd1rz",
														TargetRef: "Property_0sk0tl4",
													},
												},
												{
													DataAssociation: instance.DataAssociation{
														BaseElement: instance.BaseElement{
															ID: "DataInputAssociation_07nlhhx",
														},
														SourceRef: "DataStoreReference_0u436el",
														TargetRef: "Property_0sk0tl4",
													},
												},
												{
													DataAssociation: instance.DataAssociation{
														BaseElement: instance.BaseElement{
															ID: "DataInputAssociation_0zey6s5",
														},
														SourceRef: "DataStoreReference_14tehfn",
														TargetRef: "Property_0sk0tl4",
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
														ID: "Activity_0mwrxc9",
													},
													Name: "Copy, sign, and scan documents",
												},
												Incoming: []string{"Flow_1ukqiu1"},
												Outgoing: []string{"Flow_1c2yylu"},
											},
											Properties: []instance.Property{
												{
													BaseElement: instance.BaseElement{
														ID: "Property_184fmz8",
													},
													Name: "__targetRef_placeholder",
												},
											},
											DataInputAssociations: []instance.DataInputAssociation{
												{
													DataAssociation: instance.DataAssociation{
														BaseElement: instance.BaseElement{
															ID: "DataInputAssociation_0j7zcu6",
														},
														SourceRef: "DataObjectReference_0orm9f6",
														TargetRef: "Property_184fmz8",
													},
												},
											},
											DataOutputAssociations: []instance.DataOutputAssociation{
												{
													DataAssociation: instance.DataAssociation{
														BaseElement: instance.BaseElement{
															ID: "DataOutputAssociation_1yk6ahc",
														},
														TargetRef: "DataObjectReference_04bhv2e",
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
														ID: "Activity_0qi7bn3",
													},
													Name: "Complete data and documents",
												},
												Incoming: []string{"Flow_0a1q7lz"},
												Outgoing: []string{"Flow_0cuw6ks"},
											},
										},
									},
								},
							},
							DataObjectReferenes: []instance.DataObjectReference{
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "DataObjectReference_1hls48h",
										},
										Name: "ID document [for analysis]",
									},
									DataObjectRef: "DataObject_0d365z3",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "DataObjectReference_0orm9f6",
										},
										Name: "ID document [analysed]",
									},
									DataObjectRef: "DataObject_0z2mzsc",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "DataObjectReference_1k37h1h",
										},
										Name: "Customer data [for KYC analysis]",
									},
									DataObjectRef: "DataObject_06w6727",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "DataObjectReference_134mvuu",
										},
										Name: "Customer data [for risk assessment]",
									},
									DataObjectRef: "DataObject_1fjn0xu",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "DataObjectReference_08zd1rz",
										},
									},
									DataObjectRef: "DataObject_0qzxnu6",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "DataObjectReference_04bhv2e",
										},
										Name: "ID document [scanned]",
									},
									DataObjectRef: "DataObject_0p5084t",
								},
							},
							DataObjects: []instance.DataObject{
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "DataObject_0d365z3",
										},
									},
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "DataObject_0z2mzsc",
										},
									},
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "DataObject_06w6727",
										},
									},
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "DataObject_1fjn0xu",
										},
									},
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "DataObject_0qzxnu6",
										},
									},
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "DataObject_0p5084t",
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
													ID: "Gateway_0w4yy7d",
												},
												Name: "Legal entity or individual?",
											},
											Incoming: []string{"Flow_1tuddf9"},
											Outgoing: []string{"Flow_0seyzm2", "Flow_0wsn5cd"},
										},
									},
								},
								{
									Gateway: instance.Gateway{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Gateway_0nz6kpd",
												},
											},
											Incoming: []string{"Flow_0seyzm2", "Flow_03olk7p"},
											Outgoing: []string{"Flow_07rh70d"},
										},
									},
								},
								{
									Gateway: instance.Gateway{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Gateway_0bfi7td",
												},
												Name: "Identity of the economic owner certifiable?",
											},
											Incoming: []string{"Flow_0068ri5"},
											Outgoing: []string{"Flow_03olk7p", "Flow_1rojgff"},
										},
									},
								},
								{
									Gateway: instance.Gateway{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Gateway_0g55arm",
												},
												Name: "Subject to approval?",
											},
											Incoming: []string{"Flow_0d94m6a"},
											Outgoing: []string{"Flow_0tstl3d", "Flow_04jrss9"},
										},
									},
								},
								{
									Gateway: instance.Gateway{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Gateway_1kweljx",
												},
												Name: "Approval?",
											},
											Incoming: []string{"Flow_1tqxpdq"},
											Outgoing: []string{"Flow_1n2v14t", "Flow_1dnwq9r"},
										},
									},
								},
								{
									Gateway: instance.Gateway{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Gateway_1v08mnm",
												},
											},
											Incoming: []string{"Flow_0tstl3d", "Flow_1n2v14t"},
											Outgoing: []string{"Flow_1u4qwnh"},
										},
									},
								},
								{
									Gateway: instance.Gateway{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Gateway_1dzkbyo",
												},
												Name: "Data complete?",
											},
											Incoming: []string{"Flow_1bsig71"},
											Outgoing: []string{"Flow_0474q31", "Flow_0a1q7lz"},
										},
									},
								},
								{
									Gateway: instance.Gateway{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Gateway_0qgud4j",
												},
											},
											Incoming: []string{"Flow_0474q31", "Flow_0cuw6ks"},
											Outgoing: []string{"Flow_1ukqiu1"},
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
														ID: "Event_1k9lc9s",
													},
													Name: "Business relation ended",
												},
												Incoming: []string{"Flow_0i1ef99"},
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
														ID: "Event_150agrk",
													},
													Name: "No business relation created",
												},
												Incoming: []string{"Flow_0zdt6ay"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											SignalEventDefinitions: []instance.SignalEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "SignalEventDefinition_0wo77f7",
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
														ID: "Event_1e3rnn0",
													},
													Name: "Identity determined and new customer created",
												},
												Incoming: []string{"Flow_0i620wv"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											SignalEventDefinitions: []instance.SignalEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "SignalEventDefinition_14powuq",
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
											ID: "Flow_0dn002z",
										},
									},
									SourceRef: "Gateway_0x30xhl",
									TargetRef: "Activity_11byy83",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_00x75yi",
										},
									},
									SourceRef: "Activity_12aarlh",
									TargetRef: "Gateway_0x30xhl",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1uklf5b",
										},
									},
									SourceRef: "Activity_09zja0c",
									TargetRef: "Gateway_0x30xhl",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1bicxco",
										},
									},
									SourceRef: "Gateway_0olrjos",
									TargetRef: "Activity_12aarlh",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0yqxpk3",
										},
									},
									SourceRef: "Gateway_0olrjos",
									TargetRef: "Activity_09zja0c",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0w9hvzb",
										},
									},
									SourceRef: "Activity_1nw3l11",
									TargetRef: "Gateway_0olrjos",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1bsig71",
										},
									},
									SourceRef: "Activity_059ixil",
									TargetRef: "Gateway_1dzkbyo",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1kmt6vy",
										},
									},
									SourceRef: "Activity_0sffmrs",
									TargetRef: "Activity_059ixil",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_07rh70d",
										},
									},
									SourceRef: "Gateway_0nz6kpd",
									TargetRef: "Activity_0sffmrs",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0i1ef99",
										},
									},
									SourceRef: "Activity_1jbdjra",
									TargetRef: "Event_1k9lc9s",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1rojgff",
										},
										Name: "No",
									},
									SourceRef: "Gateway_0bfi7td",
									TargetRef: "Activity_1jbdjra",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_03olk7p",
										},
										Name: "Yes",
									},
									SourceRef: "Gateway_0bfi7td",
									TargetRef: "Gateway_0nz6kpd",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0068ri5",
										},
									},
									SourceRef: "Activity_1qr8zif",
									TargetRef: "Gateway_0bfi7td",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0wsn5cd",
										},
										Name: "Legal Entity",
									},
									SourceRef: "Gateway_0w4yy7d",
									TargetRef: "Activity_1qr8zif",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0seyzm2",
										},
										Name: "Individual Person",
									},
									SourceRef: "Gateway_0w4yy7d",
									TargetRef: "Gateway_0nz6kpd",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1tuddf9",
										},
									},
									SourceRef: "Activity_02iaxsg",
									TargetRef: "Gateway_0w4yy7d",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0xot8mo",
										},
									},
									SourceRef: "Activity_0qyqx4q",
									TargetRef: "Activity_02iaxsg",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1axad4z",
										},
									},
									SourceRef: "StartEvent_1",
									TargetRef: "Activity_0qyqx4q",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0d94m6a",
										},
									},
									SourceRef: "Activity_11byy83",
									TargetRef: "Gateway_0g55arm",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0tstl3d",
										},
										Name: "No",
									},
									SourceRef: "Gateway_0g55arm",
									TargetRef: "Gateway_1v08mnm",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_04jrss9",
										},
										Name: "Yes",
									},
									SourceRef: "Gateway_0g55arm",
									TargetRef: "Activity_1exyjv9",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1tqxpdq",
										},
									},
									SourceRef: "Activity_1exyjv9",
									TargetRef: "Gateway_1kweljx",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1n2v14t",
										},
										Name: "Yes",
									},
									SourceRef: "Gateway_1kweljx",
									TargetRef: "Gateway_1v08mnm",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1u4qwnh",
										},
									},
									SourceRef: "Gateway_1v08mnm",
									TargetRef: "Activity_1kfeipf",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1dnwq9r",
										},
										Name: "No",
									},
									SourceRef: "Gateway_1kweljx",
									TargetRef: "Activity_0gwbifa",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0zdt6ay",
										},
									},
									SourceRef: "Activity_0gwbifa",
									TargetRef: "Event_150agrk",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0jhqlw7",
										},
									},
									SourceRef: "Activity_1kfeipf",
									TargetRef: "Activity_10xe0k2",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0p0czrh",
										},
									},
									SourceRef: "Activity_10xe0k2",
									TargetRef: "Activity_0miaffi",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0i620wv",
										},
									},
									SourceRef: "Activity_0miaffi",
									TargetRef: "Event_1e3rnn0",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1c2yylu",
										},
									},
									SourceRef: "Activity_0mwrxc9",
									TargetRef: "Activity_1nw3l11",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0474q31",
										},
										Name: "Yes",
									},
									SourceRef: "Gateway_1dzkbyo",
									TargetRef: "Gateway_0qgud4j",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0a1q7lz",
										},
										Name: "No",
									},
									SourceRef: "Gateway_1dzkbyo",
									TargetRef: "Activity_0qi7bn3",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1ukqiu1",
										},
									},
									SourceRef: "Gateway_0qgud4j",
									TargetRef: "Activity_0mwrxc9",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0cuw6ks",
										},
									},
									SourceRef: "Activity_0qi7bn3",
									TargetRef: "Gateway_0qgud4j",
								},
							},
							ParallelGatewaies: []instance.ParallelGateway{
								{
									Gateway: instance.Gateway{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Gateway_0olrjos",
												},
											},
											Incoming: []string{"Flow_0w9hvzb"},
											Outgoing: []string{"Flow_0yqxpk3", "Flow_1bicxco"},
										},
									},
								},
								{
									Gateway: instance.Gateway{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Gateway_0x30xhl",
												},
											},
											Incoming: []string{"Flow_1uklf5b", "Flow_00x75yi"},
											Outgoing: []string{"Flow_0dn002z"},
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
													ID: "Activity_0gwbifa",
												},
												Name: "Reject customer request",
											},
											Incoming: []string{"Flow_1dnwq9r"},
											Outgoing: []string{"Flow_0zdt6ay"},
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
													ID: "Activity_10xe0k2",
												},
												Name: "Check for connected clients",
											},
											Incoming: []string{"Flow_0jhqlw7"},
											Outgoing: []string{"Flow_0p0czrh"},
										},
										Properties: []instance.Property{
											{
												BaseElement: instance.BaseElement{
													ID: "Property_1loybx1",
												},
												Name: "__targetRef_placeholder",
											},
										},
										DataInputAssociations: []instance.DataInputAssociation{
											{
												DataAssociation: instance.DataAssociation{
													BaseElement: instance.BaseElement{
														ID: "DataInputAssociation_1qqmaw2",
													},
													SourceRef: "DataObjectReference_134mvuu",
													TargetRef: "Property_1loybx1",
												},
											},
										},
										DataOutputAssociations: []instance.DataOutputAssociation{
											{
												DataAssociation: instance.DataAssociation{
													BaseElement: instance.BaseElement{
														ID: "DataOutputAssociation_0p01q4m",
													},
													TargetRef: "DataObjectReference_08zd1rz",
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
											ID: "DataStoreReference_06oj71q",
										},
										Name: "Customer Data (temporary storage)",
									},
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "DataStoreReference_0u436el",
										},
										Name: "Bank System",
									},
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "DataStoreReference_14tehfn",
										},
										Name: "Customer Data (temporary storage)",
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
