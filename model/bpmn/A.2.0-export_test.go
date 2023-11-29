package bpmn

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sparrow-community/bpm/model/bpmn/instance"
)

func TestA_2_0_export(t *testing.T) {
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
