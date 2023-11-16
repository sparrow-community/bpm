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
