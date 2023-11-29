package bpmn

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sparrow-community/bpm/model/bpmn/instance"
)

func TestA_2_0_rountrip(t *testing.T) {
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
