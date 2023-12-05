package bpmn

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sparrow-community/bpm/model/bpmn/instance"
)

func TestA_3_0_roudtrip(t *testing.T) {
	path := "./test/A.3.0-roundtrip.bpmn"
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
			ID:              "_1373649919111",
			Name:            "A.3.0",
			TargetNamespace: "http://www.trisotech.com/definitions/_1373649919111",
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
													ID: "_1ac4b759-40e3-4dfb-b0e3-ad1d201d6c3d",
												},
												Name: "Start Event",
											},
											Outgoing: []string{"_83f6ca65-43f7-496e-a7eb-2a4a2fc28f22"},
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
													ID: "_65f5459f-44ae-436d-a089-a91d6d78075b",
												},
												Name: "Task 1",
											},
											Incoming: []string{"_83f6ca65-43f7-496e-a7eb-2a4a2fc28f22"},
											Outgoing: []string{"_68ba9b96-b1e9-4691-bc25-a36bf5731502"},
										},
									},
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "_9fad8da5-a28c-4b6b-bb71-fbd5c65b9681",
												},
												Name: "Task 4",
											},
											Incoming: []string{"_7742093f-cd2c-415e-be71-d2514bc559c9"},
											Outgoing: []string{"_c425e783-e839-4990-9a2c-28b7341d9b2e"},
										},
									},
								},
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "_72204cd7-709c-4656-9554-3ae29b3844ce",
												},
												Name: "Task 3",
											},
											Incoming: []string{"_fe023d13-58bc-4f08-b60a-ebe4489f4190"},
											Outgoing: []string{"_88b9f814-764e-492b-b38d-d5e8dfa68400"},
										},
									},
								},

								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "_2d2d0d29-896f-49f9-8109-77a7304309c5",
												},
												Name: "Task 2",
											},
											Incoming: []string{"_250377d0-628d-463f-95f6-1f4ceed9bd8a"},
											Outgoing: []string{"_719b757a-fc92-46bd-8d10-cca5a5bbf3bf"},
										},
									},
								},
							},
							SequenceFlows: []instance.SequenceFlow{
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_83f6ca65-43f7-496e-a7eb-2a4a2fc28f22",
										},
									},
									SourceRef: "_1ac4b759-40e3-4dfb-b0e3-ad1d201d6c3d",
									TargetRef: "_65f5459f-44ae-436d-a089-a91d6d78075b",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_68ba9b96-b1e9-4691-bc25-a36bf5731502",
										},
									},
									SourceRef: "_65f5459f-44ae-436d-a089-a91d6d78075b",
									TargetRef: "_1ae31d1b-2559-4f78-a3ec-47986a49db48",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_7742093f-cd2c-415e-be71-d2514bc559c9",
										},
									},
									SourceRef: "_178e16eb-4c9e-4ea0-9644-7c5fb2b71825",
									TargetRef: "_9fad8da5-a28c-4b6b-bb71-fbd5c65b9681",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_fe023d13-58bc-4f08-b60a-ebe4489f4190",
										},
									},
									SourceRef: "_428dcbf5-8e5e-48e0-9c0c-d93003fa8c82",
									TargetRef: "_72204cd7-709c-4656-9554-3ae29b3844ce",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_250377d0-628d-463f-95f6-1f4ceed9bd8a",
										},
									},
									SourceRef: "_1ae31d1b-2559-4f78-a3ec-47986a49db48",
									TargetRef: "_2d2d0d29-896f-49f9-8109-77a7304309c5",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_719b757a-fc92-46bd-8d10-cca5a5bbf3bf",
										},
									},
									SourceRef: "_2d2d0d29-896f-49f9-8109-77a7304309c5",
									TargetRef: "_ce253897-4300-4b24-b71f-4c9535698c70",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_88b9f814-764e-492b-b38d-d5e8dfa68400",
										},
									},
									SourceRef: "_72204cd7-709c-4656-9554-3ae29b3844ce",
									TargetRef: "_ce253897-4300-4b24-b71f-4c9535698c70",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_c425e783-e839-4990-9a2c-28b7341d9b2e",
										},
									},
									SourceRef: "_9fad8da5-a28c-4b6b-bb71-fbd5c65b9681",
									TargetRef: "_10ce0b26-1b3e-46a2-85a5-62538ed2da13",
								},
							},
							SubProcess: []instance.SubProcess{
								{
									Activity: instance.Activity{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "_1ae31d1b-2559-4f78-a3ec-47986a49db48",
												},
												Name: "Collapsed\nSub-Process",
											},
											Incoming: []string{"_68ba9b96-b1e9-4691-bc25-a36bf5731502"},
											Outgoing: []string{"_250377d0-628d-463f-95f6-1f4ceed9bd8a"},
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
														ID: "_428dcbf5-8e5e-48e0-9c0c-d93003fa8c82",
													},
													Name: "Boundary Intermediate Event Non-Interrupting Message",
												},
												Outgoing: []string{"_fe023d13-58bc-4f08-b60a-ebe4489f4190"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											MessageEventDefinitions: []instance.MessageEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "",
															},
														},
													},
												},
											},
										},
									},
									CancelActivity: false,
									AttachedToRef:  "_1ae31d1b-2559-4f78-a3ec-47986a49db48",
								},
								{
									CatchEvent: instance.CatchEvent{
										Event: instance.Event{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "_178e16eb-4c9e-4ea0-9644-7c5fb2b71825",
													},
													Name: "Boundary Intermediate Event Interrupting Escalation",
												},
												Outgoing: []string{"_7742093f-cd2c-415e-be71-d2514bc559c9"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											EscalationEventDefinitions: []instance.EscalationEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "",
															},
														},
													},
												},
											},
										},
									},
									AttachedToRef: "_1ae31d1b-2559-4f78-a3ec-47986a49db48",
								},
							},
							EndEvents: []instance.EndEvent{
								{
									ThrowEvent: instance.ThrowEvent{
										Event: instance.Event{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "_ce253897-4300-4b24-b71f-4c9535698c70",
													},
													Name: "End Event 1",
												},
												Incoming: []string{"_719b757a-fc92-46bd-8d10-cca5a5bbf3bf", "_88b9f814-764e-492b-b38d-d5e8dfa68400"},
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
														ID: "_10ce0b26-1b3e-46a2-85a5-62538ed2da13",
													},
													Name: "End Event 2",
												},
												Incoming: []string{"_c425e783-e839-4990-9a2c-28b7341d9b2e"},
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
		t.Errorf("Unexpected result (-want, +got):\n%s", diff)
	}
}
