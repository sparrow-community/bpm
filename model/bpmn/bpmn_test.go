package bpmn

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sparrow-community/bpm/model/bpmn/instance"
)

// test files srouce from https://github.com/bpmn-miwg/bpmn-miwg-test-suite/tree/master/bpmn.io%20(Camunda%20Modeler)%2011.5.0

func TestBpmnModelInstanceFromFile_A_1_0_export(t *testing.T) {
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

func TestBpmnModelInstanceFromFile_A_1_0_roundtrip(t *testing.T) {
	// create test use ./test/A.1.0-roundtrip.bpmn
	path := "./test/A.1.0-roundtrip.bpmn"
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
			ID:              "_1373649849716",
			Name:            "A.1.0",
			TargetNamespace: "http://www.trisotech.com/definitions/_1373649849716",
			Rootlemnts: instance.Rootlemnts{
				Processes: []instance.Process{
					{
						CallableElement: instance.CallableElement{
							RootElement: instance.RootElement{
								BaseElement: instance.BaseElement{
									ID: "WFP-6-",
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
														ID: "_93c466ab-b271-4376-a427-f4c353d55ce8",
													},
													Name: "Start Event",
												},
												Outgoing: []string{"_e16564d7-0c4c-413e-95f6-f668a3f851fb"},
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
													ID: "_ec59e164-68b4-4f94-98de-ffb1c58a84af",
												},
												Name: "Task 1",
											},
											Incoming: []string{"_e16564d7-0c4c-413e-95f6-f668a3f851fb"},
											Outgoing: []string{"_d77dd5ec-e4e7-420e-bbe7-8ac9cd1df599"},
										},
									},
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "_820c21c0-45f3-473b-813f-06381cc637cd",
												},
												Name: "Task 2",
											},
											Incoming: []string{"_d77dd5ec-e4e7-420e-bbe7-8ac9cd1df599"},
											Outgoing: []string{"_2aa47410-1b0e-4f8b-ad54-d6f798080cb4"},
										},
									},
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "_e70a6fcb-913c-4a7b-a65d-e83adc73d69c",
												},
												Name: "Task 3",
											},
											Incoming: []string{"_2aa47410-1b0e-4f8b-ad54-d6f798080cb4"},
											Outgoing: []string{"_8e8fe679-eb3b-4c43-a4d6-891e7087ff80"},
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
														ID: "_a47df184-085b-49f7-bb82-031c84625821",
													},
													Name: "End Event",
												},
												Incoming: []string{"_8e8fe679-eb3b-4c43-a4d6-891e7087ff80"},
											},
										},
									},
								},
							},
							SequenceFlows: []instance.SequenceFlow{
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_e16564d7-0c4c-413e-95f6-f668a3f851fb",
										},
										Name: "",
									},
									SourceRef: "_93c466ab-b271-4376-a427-f4c353d55ce8",
									TargetRef: "_ec59e164-68b4-4f94-98de-ffb1c58a84af",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_d77dd5ec-e4e7-420e-bbe7-8ac9cd1df599",
										},
									},
									SourceRef: "_ec59e164-68b4-4f94-98de-ffb1c58a84af",
									TargetRef: "_820c21c0-45f3-473b-813f-06381cc637cd",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_2aa47410-1b0e-4f8b-ad54-d6f798080cb4",
										},
									},
									SourceRef: "_820c21c0-45f3-473b-813f-06381cc637cd",
									TargetRef: "_e70a6fcb-913c-4a7b-a65d-e83adc73d69c",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_8e8fe679-eb3b-4c43-a4d6-891e7087ff80",
										},
									},
									SourceRef: "_e70a6fcb-913c-4a7b-a65d-e83adc73d69c",
									TargetRef: "_a47df184-085b-49f7-bb82-031c84625821",
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

func TestBpmnModelInstanceFromFile_A_2_0_export(t *testing.T) {
	// create test use ./test/A.2.0-export.bpmn
	path := "./test/A.2.0-export.bpmn"
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
			TargetNamespace: "http://bpmn.io/bpmn",
			ID:              "sid-38422fae-e03e-43a3-bef4-bd33b32041b2",
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
							StartEvents: []instance.StartEvent{{
								CatchEvent: instance.CatchEvent{
									Event: instance.Event{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Event_072o7cv",
												},
												Name: "Start Event",
											},
											Outgoing: []string{"Flow_12fenxe"},
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
													ID: "Activity_0opq70y",
												},
												Name: "Task 1",
											},
											Incoming: []string{"Flow_12fenxe"},
											Outgoing: []string{"Flow_1yupnqc"},
										},
									},
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Activity_1ljp29t",
												},
												Name: "Task 2",
											},
											Incoming: []string{"Flow_0dd1rck"},
											Outgoing: []string{"Flow_0y0m2k8"},
										},
									},
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Activity_0jhawx0",
												},
												Name: "Task 3",
											},
											Incoming: []string{"Flow_0x796n6"},
											Outgoing: []string{"Flow_1lk8qao"},
										},
									},
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Activity_0ddly78",
												},
												Name: "Task 4",
											},
											Incoming: []string{"Flow_1801a2c"},
											Outgoing: []string{"Flow_17lrcjr"},
										},
									},
								},
							},
							SequenceFlows: []instance.SequenceFlow{
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_12fenxe",
										},
									},
									SourceRef: "Event_072o7cv",
									TargetRef: "Activity_0opq70y",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1yupnqc",
										},
									},
									SourceRef: "Activity_0opq70y",
									TargetRef: "Gateway_03s9abx",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0dd1rck",
										},
									},
									SourceRef: "Gateway_03s9abx",
									TargetRef: "Activity_1ljp29t",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0x796n6",
										},
									},
									SourceRef: "Gateway_03s9abx",
									TargetRef: "Activity_0jhawx0",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1801a2c",
										},
									},
									SourceRef: "Gateway_03s9abx",
									TargetRef: "Activity_0ddly78",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_0y0m2k8",
										},
									},
									SourceRef: "Activity_1ljp29t",
									TargetRef: "Event_1d5wxn1",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_1lk8qao",
										},
									},
									SourceRef: "Activity_0jhawx0",
									TargetRef: "Gateway_03haizn",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_17lrcjr",
										},
									},
									SourceRef: "Activity_0ddly78",
									TargetRef: "Gateway_03haizn",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "Flow_05zv16r",
										},
									},
									SourceRef: "Gateway_03haizn",
									TargetRef: "Event_1d5wxn1",
								},
							},
							ExclusiveGatewaies: []instance.ExclusiveGateway{
								{
									Gateway: instance.Gateway{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Gateway_03s9abx",
												},
												Name: "Gateway (Split Flow)",
											},
											Incoming: []string{"Flow_1yupnqc"},
											Outgoing: []string{"Flow_0dd1rck", "Flow_0x796n6", "Flow_1801a2c"},
										},
									},
								},
								{
									Gateway: instance.Gateway{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "Gateway_03haizn",
												},
												Name: "Gateway (Merge Flows)",
											},
											Incoming: []string{"Flow_1lk8qao", "Flow_17lrcjr"},
											Outgoing: []string{"Flow_05zv16r"},
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
														ID: "Event_1d5wxn1",
													},
													Name: "End Event",
												},
												Incoming: []string{"Flow_0y0m2k8", "Flow_05zv16r"},
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

func TestBpmnModelInstanceFromFile_A_2_0_roundtrip(t *testing.T) {
	// create test use ./test/A.2.0-roundtrip.bpmn
	path := "./test/A.2.0-roundtrip.bpmn"
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
			TargetNamespace: "http://www.trisotech.com/definitions/_1373649889746",
			ID:              "_1373649889746",
			Name:            "A.2.0",
			Rootlemnts: instance.Rootlemnts{
				Processes: []instance.Process{
					{
						CallableElement: instance.CallableElement{
							RootElement: instance.RootElement{
								BaseElement: instance.BaseElement{
									ID: "WFP-6-",
								},
							},
						},
						IsExecutable: false,
						FlowElements: instance.FlowElements{
							StartEvents: []instance.StartEvent{{
								CatchEvent: instance.CatchEvent{
									Event: instance.Event{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "_6b5db6a9-037a-49ad-9201-09201e2aaa97",
												},
												Name: "Start Event",
											},
											Outgoing: []string{"_b50f530c-3450-4e1a-b81f-ea346dc6e1cb"},
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
													ID: "_5a972b87-735d-454a-b31c-f52fb3afc5c7",
												},
												Name: "Task 1",
											},
											Incoming: []string{"_b50f530c-3450-4e1a-b81f-ea346dc6e1cb"},
											Outgoing: []string{"_fe74c141-8843-4b00-a704-5e5e13be53b0"},
										},
									},
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "_4f7d62d7-f0e6-46bc-be00-69e02da38f65",
												},
												Name: "Task 2",
											},
											Incoming: []string{"_f1478fb7-98c4-4c01-8c15-68bd04c91535"},
											Outgoing: []string{"_a3d40a56-9b7f-417e-911e-d39e7f18b90c"},
										},
									},
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "_e6eb725a-34bc-45c7-aed0-9f9596cd7bee",
												},
												Name: "Task 3",
											},
											Incoming: []string{"_a1570a53-28d2-41b1-a3a2-3e50c00d747e"},
											Outgoing: []string{"_e9ebc7c7-995d-46db-86ce-d823bc2b4687"},
										},
									},
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "_7d399717-1aba-47ac-8d7d-8aaa033255e0",
												},
												Name: "Task 4",
											},
											Incoming: []string{"_20ebb3c1-5178-4c7c-a91d-23e58f2aa73b"},
											Outgoing: []string{"_698b593f-18eb-42ea-b8cd-bcd51e1514cc"},
										},
									},
								},
							},
							SequenceFlows: []instance.SequenceFlow{
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_b50f530c-3450-4e1a-b81f-ea346dc6e1cb",
										},
									},
									SourceRef: "_6b5db6a9-037a-49ad-9201-09201e2aaa97",
									TargetRef: "_5a972b87-735d-454a-b31c-f52fb3afc5c7",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_fe74c141-8843-4b00-a704-5e5e13be53b0",
										},
									},
									SourceRef: "_5a972b87-735d-454a-b31c-f52fb3afc5c7",
									TargetRef: "_35fe57a7-1302-44e2-bf58-032f11af7ecb",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_f1478fb7-98c4-4c01-8c15-68bd04c91535",
										},
									},
									SourceRef: "_35fe57a7-1302-44e2-bf58-032f11af7ecb",
									TargetRef: "_4f7d62d7-f0e6-46bc-be00-69e02da38f65",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_a3d40a56-9b7f-417e-911e-d39e7f18b90c",
										},
									},
									SourceRef: "_4f7d62d7-f0e6-46bc-be00-69e02da38f65",
									TargetRef: "_258f51eb-b764-4a71-b681-3a01cca14143",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_e9ebc7c7-995d-46db-86ce-d823bc2b4687",
										},
									},
									SourceRef: "_e6eb725a-34bc-45c7-aed0-9f9596cd7bee",
									TargetRef: "_33c66216-391c-49c2-aa19-d8f0b7f5f91d",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_698b593f-18eb-42ea-b8cd-bcd51e1514cc",
										},
									},
									SourceRef: "_7d399717-1aba-47ac-8d7d-8aaa033255e0",
									TargetRef: "_33c66216-391c-49c2-aa19-d8f0b7f5f91d",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_d4ce87c6-1373-45d6-a3b4-fbb2a04ee2e5",
										},
									},
									SourceRef: "_33c66216-391c-49c2-aa19-d8f0b7f5f91d",
									TargetRef: "_258f51eb-b764-4a71-b681-3a01cca14143",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_a1570a53-28d2-41b1-a3a2-3e50c00d747e",
										},
									},
									SourceRef: "_35fe57a7-1302-44e2-bf58-032f11af7ecb",
									TargetRef: "_e6eb725a-34bc-45c7-aed0-9f9596cd7bee",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_20ebb3c1-5178-4c7c-a91d-23e58f2aa73b",
										},
									},
									SourceRef: "_35fe57a7-1302-44e2-bf58-032f11af7ecb",
									TargetRef: "_7d399717-1aba-47ac-8d7d-8aaa033255e0",
								},
							},
							ExclusiveGatewaies: []instance.ExclusiveGateway{
								{
									Gateway: instance.Gateway{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "_35fe57a7-1302-44e2-bf58-032f11af7ecb",
												},
												Name: "Gateway\n(Split Flow)",
											},
											Incoming: []string{"_fe74c141-8843-4b00-a704-5e5e13be53b0"},
											Outgoing: []string{"_f1478fb7-98c4-4c01-8c15-68bd04c91535", "_a1570a53-28d2-41b1-a3a2-3e50c00d747e", "_20ebb3c1-5178-4c7c-a91d-23e58f2aa73b"},
										},
									},
								},
								{
									Gateway: instance.Gateway{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "_33c66216-391c-49c2-aa19-d8f0b7f5f91d",
												},
												Name: "Gateway\n(Merge Flows)",
											},
											Incoming: []string{"_e9ebc7c7-995d-46db-86ce-d823bc2b4687", "_698b593f-18eb-42ea-b8cd-bcd51e1514cc"},
											Outgoing: []string{"_d4ce87c6-1373-45d6-a3b4-fbb2a04ee2e5"},
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
														ID: "_258f51eb-b764-4a71-b681-3a01cca14143",
													},
													Name: "End Event",
												},
												Incoming: []string{"_a3d40a56-9b7f-417e-911e-d39e7f18b90c", "_d4ce87c6-1373-45d6-a3b4-fbb2a04ee2e5"},
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

func TestBpmnModelInstanceFromFile_A_2_0_rountrip(t *testing.T) {
	// create test use ./test/A.2.0-rountrip.bpmn
	path := "./test/A.2.0-rountrip.bpmn"
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
			TargetNamespace: "http://www.trisotech.com/definitions/_1373649889746",
			ID:              "_1373649889746",
			Name:            "A.2.0",
			Rootlemnts: instance.Rootlemnts{
				Processes: []instance.Process{
					{
						CallableElement: instance.CallableElement{
							RootElement: instance.RootElement{
								BaseElement: instance.BaseElement{
									ID: "WFP-6-",
								},
							},
						},
						IsExecutable: false,
						FlowElements: instance.FlowElements{
							StartEvents: []instance.StartEvent{{
								CatchEvent: instance.CatchEvent{
									Event: instance.Event{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "_6b5db6a9-037a-49ad-9201-09201e2aaa97",
												},
												Name: "Start Event",
											},
											Outgoing: []string{"_b50f530c-3450-4e1a-b81f-ea346dc6e1cb"},
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
													ID: "_5a972b87-735d-454a-b31c-f52fb3afc5c7",
												},
												Name: "Task 1",
											},
											Incoming: []string{"_b50f530c-3450-4e1a-b81f-ea346dc6e1cb"},
											Outgoing: []string{"_fe74c141-8843-4b00-a704-5e5e13be53b0"},
										},
										StartQuantity:      1,
										CompletionQuantity: 1,
									},
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "_4f7d62d7-f0e6-46bc-be00-69e02da38f65",
												},
												Name: "Task 2",
											},
											Incoming: []string{"_f1478fb7-98c4-4c01-8c15-68bd04c91535"},
											Outgoing: []string{"_a3d40a56-9b7f-417e-911e-d39e7f18b90c"},
										},
										StartQuantity:      1,
										CompletionQuantity: 1,
									},
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "_e6eb725a-34bc-45c7-aed0-9f9596cd7bee",
												},
												Name: "Task 3",
											},
											Incoming: []string{"_a1570a53-28d2-41b1-a3a2-3e50c00d747e"},
											Outgoing: []string{"_e9ebc7c7-995d-46db-86ce-d823bc2b4687"},
										},
										StartQuantity:      1,
										CompletionQuantity: 1,
									},
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "_7d399717-1aba-47ac-8d7d-8aaa033255e0",
												},
												Name: "Task 4",
											},
											Incoming: []string{"_20ebb3c1-5178-4c7c-a91d-23e58f2aa73b"},
											Outgoing: []string{"_698b593f-18eb-42ea-b8cd-bcd51e1514cc"},
										},
										StartQuantity:      1,
										CompletionQuantity: 1,
									},
								},
							},
							SequenceFlows: []instance.SequenceFlow{
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_b50f530c-3450-4e1a-b81f-ea346dc6e1cb",
										},
									},
									SourceRef: "_6b5db6a9-037a-49ad-9201-09201e2aaa97",
									TargetRef: "_5a972b87-735d-454a-b31c-f52fb3afc5c7",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_fe74c141-8843-4b00-a704-5e5e13be53b0",
										},
									},
									SourceRef: "_5a972b87-735d-454a-b31c-f52fb3afc5c7",
									TargetRef: "_35fe57a7-1302-44e2-bf58-032f11af7ecb",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_f1478fb7-98c4-4c01-8c15-68bd04c91535",
										},
									},
									SourceRef: "_35fe57a7-1302-44e2-bf58-032f11af7ecb",
									TargetRef: "_4f7d62d7-f0e6-46bc-be00-69e02da38f65",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_a3d40a56-9b7f-417e-911e-d39e7f18b90c",
										},
									},
									SourceRef: "_4f7d62d7-f0e6-46bc-be00-69e02da38f65",
									TargetRef: "_258f51eb-b764-4a71-b681-3a01cca14143",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_e9ebc7c7-995d-46db-86ce-d823bc2b4687",
										},
									},
									SourceRef: "_e6eb725a-34bc-45c7-aed0-9f9596cd7bee",
									TargetRef: "_33c66216-391c-49c2-aa19-d8f0b7f5f91d",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_698b593f-18eb-42ea-b8cd-bcd51e1514cc",
										},
									},
									SourceRef: "_7d399717-1aba-47ac-8d7d-8aaa033255e0",
									TargetRef: "_33c66216-391c-49c2-aa19-d8f0b7f5f91d",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_d4ce87c6-1373-45d6-a3b4-fbb2a04ee2e5",
										},
									},
									SourceRef: "_33c66216-391c-49c2-aa19-d8f0b7f5f91d",
									TargetRef: "_258f51eb-b764-4a71-b681-3a01cca14143",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_a1570a53-28d2-41b1-a3a2-3e50c00d747e",
										},
									},
									SourceRef: "_35fe57a7-1302-44e2-bf58-032f11af7ecb",
									TargetRef: "_e6eb725a-34bc-45c7-aed0-9f9596cd7bee",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_20ebb3c1-5178-4c7c-a91d-23e58f2aa73b",
										},
									},
									SourceRef: "_35fe57a7-1302-44e2-bf58-032f11af7ecb",
									TargetRef: "_7d399717-1aba-47ac-8d7d-8aaa033255e0",
								},
							},
							ExclusiveGatewaies: []instance.ExclusiveGateway{
								{
									Gateway: instance.Gateway{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "_35fe57a7-1302-44e2-bf58-032f11af7ecb",
												},
												Name: "Gateway\n(Split Flow)",
											},
											Incoming: []string{"_fe74c141-8843-4b00-a704-5e5e13be53b0"},
											Outgoing: []string{"_f1478fb7-98c4-4c01-8c15-68bd04c91535", "_a1570a53-28d2-41b1-a3a2-3e50c00d747e", "_20ebb3c1-5178-4c7c-a91d-23e58f2aa73b"},
										},
										GatewayDirection: instance.GatewayDirectionUnspecified,
									},
								},
								{
									Gateway: instance.Gateway{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "_33c66216-391c-49c2-aa19-d8f0b7f5f91d",
												},
												Name: "Gateway\n(Merge Flows)",
											},
											Incoming: []string{"_e9ebc7c7-995d-46db-86ce-d823bc2b4687", "_698b593f-18eb-42ea-b8cd-bcd51e1514cc"},
											Outgoing: []string{"_d4ce87c6-1373-45d6-a3b4-fbb2a04ee2e5"},
										},
										GatewayDirection: instance.GatewayDirectionUnspecified,
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
														ID: "_258f51eb-b764-4a71-b681-3a01cca14143",
													},
													Name: "End Event",
												},
												Incoming: []string{"_a3d40a56-9b7f-417e-911e-d39e7f18b90c", "_d4ce87c6-1373-45d6-a3b4-fbb2a04ee2e5"},
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

func TestBpmnModelInstanceFromFile_A_2_1_export(t *testing.T) {
	// create test using ./test/A.2.1-export.bpmn
	path := "./test/A.2.1-export.bpmn"
	modelInstance, err := BpmnModelInstanceFromFile(path)
	if err != nil {
		t.Fatalf("BpmnModelInstanceFromFile error %s: %s", path, err)
	}
	if modelInstance == nil {
		t.Fatalf("modelInstance is nil, %s", path)
	}

	expected := &BpmnModelInstance{
		Definitions: &instance.Definitions{
			XMLName: xml.Name{},
		},
	}
	if diff := cmp.Diff(expected, modelInstance); diff != "" {
		t.Errorf("Unexpected result (-want, +got):\n%s", diff)
	}
}
