package bpmn

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sparrow-community/bpm/model/bpmn/instance"
)

func TestC_8_0_roundtrip(t *testing.T) {
	// create test use ./test/C.8.0-roundtrip.bpmn
	path := "./test/C.8.0-roundtrip.bpmn"
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
			ID:              "Definitions__a9e3a63e-db87-4fa1-9091-328acae4f490",
			Name:            "Vacation Request - (colors)",
			TargetNamespace: "http://www.softwareag.com/aris/bpmn2",
			Exporter:        "Workflow Modeler",
			ExporterVersion: "7.14.1.202203211419",
			RootElemnts: instance.RootElemnts{
				ItemDefinitions: []instance.ItemDefinition{
					{
						RootElement: instance.RootElement{
							BaseElement: instance.BaseElement{
								ID: "_triso-default-bpmnItemDefinition-string_id",
							},
						},
						StructureRef: "feel:string",
					},
				},
				Processes: []instance.Process{
					{
						CallableElement: instance.CallableElement{
							RootElement: instance.RootElement{
								BaseElement: instance.BaseElement{
									ID:                "VacationRequestProcess",
									Documentation:     []instance.Documentation{{Value: "Vacation Request - BPMN MIWG demo for 2022"}},
									ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil, nil, nil, nil, nil}},
								},
							},
							Name: "Vacation Request - (i18n)",
							IoSpecification: instance.IoSpecification{
								InputOutputSpecification: instance.InputOutputSpecification{
									DataInputs: []instance.DataInput{
										{
											BaseElement: instance.BaseElement{
												ID:                "dataInputdataInput_8b9aa28f-5974-4087-9895-0467c25635dc",
												ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil, nil, nil, nil}},
											},
											Name: "Employee Badge Number",
										},
										{
											BaseElement: instance.BaseElement{
												ID:                "dataInputdataInput_2bed4da9-f468-4fe0-a37c-ae5d4c17e323",
												ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil, nil, nil, nil}},
											},
											Name: "From",
										},
										{
											BaseElement: instance.BaseElement{
												ID:                "dataInputdataInput_103fd816-b422-420d-aec9-c28b54f620db",
												ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil, nil, nil, nil}},
											},
											Name: "To",
										},
									},
									DataOutputs: []instance.DataOutput{
										{
											BaseElement: instance.BaseElement{
												ID:                "dataOutputdataOutput_dac8ee76-f637-4fd9-8357-6a87fd11ef41",
												ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil, nil, nil, nil}},
											},
											Name: "Reason",
										},
										{
											BaseElement: instance.BaseElement{
												ID:                "dataOutputdataOutput_c2d108b2-e6c0-47e6-a3cb-aa5321bcaeaf",
												ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil, nil, nil, nil}},
											},
											Name: "Vacation Approval",
										},
									},
									InputSets: []instance.InputSet{
										{
											BaseElement: instance.BaseElement{
												ID: "_4d1f4479-b14c-4dd0-9675-607384e9dd81",
											},
											DataInputRefs: []string{"dataInputdataInput_8b9aa28f-5974-4087-9895-0467c25635dc", "dataInputdataInput_2bed4da9-f468-4fe0-a37c-ae5d4c17e323", "dataInputdataInput_103fd816-b422-420d-aec9-c28b54f620db"},
										},
									},
									OutputSets: []instance.OutputSet{
										{
											BaseElement: instance.BaseElement{
												ID: "_10b372b4-3cdd-4ddd-aa58-129b5df7b2f8",
											},
											DataOutputRefs: []string{"dataOutputdataOutput_dac8ee76-f637-4fd9-8357-6a87fd11ef41", "dataOutputdataOutput_c2d108b2-e6c0-47e6-a3cb-aa5321bcaeaf"},
										},
									},
								},
							},
						},
						ProcessType:  "None",
						IsExecutable: false,
						FlowElements: instance.FlowElements{
							StartEvents: []instance.StartEvent{
								{
									CatchEvent: instance.CatchEvent{
										Event: instance.Event{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID:                "_b1625a52-aaf0-4694-86cb-7af891212ac6",
														ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil, nil, nil, nil, nil, nil, nil, nil, nil}},
													},
													Name: "Vacation Request Received",
												},
												Outgoing: []string{"_b02b5bfa-5629-4af5-b1bc-ac2a86d88adb"},
											},
										},
									},
								},
							},
							SequenceFlows: []instance.SequenceFlow{
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID:                "_b02b5bfa-5629-4af5-b1bc-ac2a86d88adb",
											ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil, nil, nil}},
										},
									},
									SourceRef: "_b1625a52-aaf0-4694-86cb-7af891212ac6",
									TargetRef: "_2b960d84-feb1-46a9-a1a1-c300dd996b99",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID:                "_81ed1b89-5a4d-4028-a245-4fd38b866b6d",
											ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil, nil, nil}},
										},
									},
									SourceRef: "_5e16a4e0-0f23-482a-be47-d3edbc5741ba",
									TargetRef: "_1cd5fe29-b3ec-4f21-a1aa-57773f0729ca",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID:                "_2a5ee73c-c495-48a0-945b-3dabab28ca91",
											ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil, nil, nil, nil, nil, nil}},
										},
										Name: "Refused",
									},
									SourceRef: "_42367c5f-d084-44ee-90c7-960d1ab02a3b",
									TargetRef: "_9ed61a6a-7cc1-4ed1-86d8-03482b0983c9",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID:                "_d79c991e-446c-47d1-ac9d-9d0113e35b93",
											ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil, nil, nil}},
										},
										Name: "",
									},
									SourceRef: "_2b960d84-feb1-46a9-a1a1-c300dd996b99",
									TargetRef: "_1a818a94-ba6f-413b-a7e8-6f8fd2a11e32",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID:                "_93ac178c-3efc-48ae-8a99-8d29a0c4aaf6",
											ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil, nil, nil}},
										},
									},
									SourceRef: "_02232e32-c3d2-473c-a15d-9c5dca00eadc",
									TargetRef: "_3ae826ca-5f38-43c4-be3a-35d1157aa27f",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID:                "_130d903a-f211-4b3e-9711-9061f5978d94",
											ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil, nil, nil}},
										},
									},
									SourceRef: "_9ed61a6a-7cc1-4ed1-86d8-03482b0983c9",
									TargetRef: "_1688f604-5edf-4187-ad9e-18f74bfede53",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID:                "_d3a575d8-cd41-4246-ad7a-7a068de6679f",
											ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil, nil, nil, nil, nil, nil}},
										},
										Name: "Refused",
									},
									SourceRef: "_64bb8b55-d348-41d6-9ea8-f333b1d8cb69",
									TargetRef: "_02232e32-c3d2-473c-a15d-9c5dca00eadc",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID:                "_f2b0da63-d841-4457-ad85-7d86c8b5c1d2",
											ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil, nil, nil, nil, nil, nil}},
										},
										Name: "Approved",
									},
									SourceRef: "_64bb8b55-d348-41d6-9ea8-f333b1d8cb69",
									TargetRef: "_a97c1a48-faba-447b-bfa6-7aa81a6fe0a0",
									ConditionExpression: instance.ExpressionUnMarshal{
										Type: "tFormalExpression",
										ExpressionSubstitution: &instance.FormalExpression{
											Expression: instance.Expression{
												BaseElementWithMixedContent: instance.BaseElementWithMixedContent{
													ID: "_f2b0da63-d841-4457-ad85-7d86c8b5c1d2condExpr",
												},
											},
											Value: `Vacation Approval = "Approved"`,
										},
									},
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID:                "_a5441c65-06f8-4972-8158-e6f61f904841",
											ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil, nil, nil}},
										},
									},
									SourceRef: "_a97c1a48-faba-447b-bfa6-7aa81a6fe0a0",
									TargetRef: "_5e16a4e0-0f23-482a-be47-d3edbc5741ba",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID:                "_188d38f6-b73d-4c62-964b-affc6fe040c1",
											ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil, nil, nil}},
										},
									},
									SourceRef: "_93ec9873-edf1-4549-b052-961994ec8234",
									TargetRef: "_4b72053b-8ebb-4ae6-99c6-7c93cf1c1d1b",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID:                "_325973e7-0bc8-4136-b6df-be1e681d8608",
											ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil, nil, nil, nil, nil, nil}},
										},
										Name: "Approved",
									},
									SourceRef: "_42367c5f-d084-44ee-90c7-960d1ab02a3b",
									TargetRef: "_93ec9873-edf1-4549-b052-961994ec8234",
									ConditionExpression: instance.ExpressionUnMarshal{
										Type: "tFormalExpression",
										ExpressionSubstitution: &instance.FormalExpression{
											Expression: instance.Expression{
												BaseElementWithMixedContent: instance.BaseElementWithMixedContent{
													ID: "_325973e7-0bc8-4136-b6df-be1e681d8608condExpr",
												},
											},
											Value: `Vacation Approval = "Approved"`,
										},
									},
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID:                "_3c50d7bb-5ecc-4d34-82a9-a5fa41d9c018",
											ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil, nil, nil}},
										},
										Name: "",
									},
									SourceRef: "_79523269-7444-4b01-90e9-e23957a9d020",
									TargetRef: "_64bb8b55-d348-41d6-9ea8-f333b1d8cb69",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID:                "_f250e125-23eb-45cd-9c61-191d172df093",
											ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil, nil, nil}},
										},
									},
									SourceRef: "_f8fcb377-3d7d-4138-9a7e-6ab58b97e29d",
									TargetRef: "_b4d636eb-b501-4462-93c8-04652db10307",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID:                "_68ef27a7-011e-4ec0-ad45-41f52093de68",
											ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil, nil, nil}},
										},
										Name: "",
									},
									SourceRef: "_4b72053b-8ebb-4ae6-99c6-7c93cf1c1d1b",
									TargetRef: "_6677ef80-82df-4951-919d-1f36123b681b",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID:                "_0a1c4f20-509f-4aeb-baf9-acc762f4fdf9",
											ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil, nil, nil, nil, nil, nil}},
										},
										Name: "Manual Validation Required",
									},
									SourceRef: "_42367c5f-d084-44ee-90c7-960d1ab02a3b",
									TargetRef: "_79523269-7444-4b01-90e9-e23957a9d020",
									ConditionExpression: instance.ExpressionUnMarshal{
										Type: "tFormalExpression",
										ExpressionSubstitution: &instance.FormalExpression{
											Expression: instance.Expression{
												BaseElementWithMixedContent: instance.BaseElementWithMixedContent{
													ID: "_0a1c4f20-509f-4aeb-baf9-acc762f4fdf9condExpr",
												},
											},
											Value: `Vacation Approval = "Manual Validation Required"`,
										},
									},
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID:                "_ac1d82df-9bf6-404b-be70-b052e311f8b7",
											ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil, nil, nil}},
										},
									},
									SourceRef: "_1a818a94-ba6f-413b-a7e8-6f8fd2a11e32",
									TargetRef: "_42367c5f-d084-44ee-90c7-960d1ab02a3b",
								},
							},
							ServiceTasks: []instance.ServiceTask{
								{
									Task: instance.Task{
										Activity: instance.Activity{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "_5e16a4e0-0f23-482a-be47-d3edbc5741ba",
														ExtensionElements: instance.ExtensionElements{
															Any: []xml.Token{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil},
														},
														Any: nil,
													},
													Name: "Update Remaining Vacation",
												},
												Incoming: []string{"_a5441c65-06f8-4972-8158-e6f61f904841"},
												Outgoing: []string{"_81ed1b89-5a4d-4028-a245-4fd38b866b6d"},
											},
										},
									},
									Implementation: "##WebService",
								},
								{
									Task: instance.Task{
										Activity: instance.Activity{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "_4b72053b-8ebb-4ae6-99c6-7c93cf1c1d1b",
														ExtensionElements: instance.ExtensionElements{
															Any: []xml.Token{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil},
														},
													},
													Name: "Update Remaining Vacation",
												},
												Incoming: []string{"_188d38f6-b73d-4c62-964b-affc6fe040c1"},
												Outgoing: []string{"_68ef27a7-011e-4ec0-ad45-41f52093de68"},
											},
										},
									},
									Implementation: "##WebService",
								},
								{
									Task: instance.Task{
										Activity: instance.Activity{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "_2b960d84-feb1-46a9-a1a1-c300dd996b99",
														ExtensionElements: instance.ExtensionElements{
															Any: []xml.Token{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil},
														},
													},
													Name: "Fetch Vacation Information",
												},
												Incoming: []string{"_b02b5bfa-5629-4af5-b1bc-ac2a86d88adb"},
												Outgoing: []string{"_d79c991e-446c-47d1-ac9d-9d0113e35b93"},
											},
											IoSpecification: instance.IoSpecification{
												InputOutputSpecification: instance.InputOutputSpecification{
													DataInputs: []instance.DataInput{
														{
															BaseElement: instance.BaseElement{
																ID: "DataInput__2b960d84-feb1-46a9-a1a1-c300dd996b99",
															},
														},
													},
													DataOutputs: []instance.DataOutput{
														{
															BaseElement: instance.BaseElement{
																ID: "DataOutput__2b960d84-feb1-46a9-a1a1-c300dd996b99",
															},
														},
													},
													InputSets: []instance.InputSet{
														{
															BaseElement: instance.BaseElement{
																ID: "_74206bda-6edc-4a98-a433-af5446be8809",
															},
															DataInputRefs: []string{"DataInput__2b960d84-feb1-46a9-a1a1-c300dd996b99"},
														},
													},
													OutputSets: []instance.OutputSet{
														{
															BaseElement: instance.BaseElement{
																ID: "_cd959c80-ff41-401a-9282-b76dcf95aaf2",
															},
															DataOutputRefs: []string{"DataOutput__2b960d84-feb1-46a9-a1a1-c300dd996b99"},
														},
													},
												},
											},
											DataInputAssociations: []instance.DataInputAssociation{
												{
													DataAssociation: instance.DataAssociation{
														BaseElement: instance.BaseElement{
															ID:                "_f4f41761-cdd8-4b90-a956-9182345df488",
															ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil, nil}},
														},
														SourceRef: "dataInputdataInput_8b9aa28f-5974-4087-9895-0467c25635dc",
														TargetRef: "DataInput__2b960d84-feb1-46a9-a1a1-c300dd996b99",
													},
												},
											},
											DataOutputAssociations: []instance.DataOutputAssociation{
												{
													DataAssociation: instance.DataAssociation{
														BaseElement: instance.BaseElement{
															ID:                "_40d3cb58-31bb-47a4-9591-032a38011de3",
															ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil, nil}},
														},
														SourceRef: "DataOutput__2b960d84-feb1-46a9-a1a1-c300dd996b99",
														TargetRef: "Reference__775c93ab-82e5-4a62-97f0-d305eda76c92",
													},
												},
											},
										},
									},
									Implementation: "##WebService",
								},
							},
							BoundaryEvents: []instance.BoundaryEvent{
								{
									CatchEvent: instance.CatchEvent{
										Event: instance.Event{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID:                "_f8fcb377-3d7d-4138-9a7e-6ab58b97e29d",
														ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil}},
													},
												},
												Outgoing: []string{"_f250e125-23eb-45cd-9c61-191d172df093"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											ErrorEventDefinitions: []instance.ErrorEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "Definition__f8fcb377-3d7d-4138-9a7e-6ab58b97e29d",
															},
														},
													},
												},
											},
										},
									},
									AttachedToRef: "_2b960d84-feb1-46a9-a1a1-c300dd996b99",
								},
							},
							DataObjectReferenes: []instance.DataObjectReference{
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID:                "Reference__775c93ab-82e5-4a62-97f0-d305eda76c92",
											ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil}},
										},
										Name: "Current Vacation Status",
									},
									DataObjectRef: "_775c93ab-82e5-4a62-97f0-d305eda76c92",
								},
							},
							DataObjects: []instance.DataObject{
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID:                "_775c93ab-82e5-4a62-97f0-d305eda76c92",
											ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil, nil, nil}},
										},
										Name: "Current Vacation Status",
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
														ID:                "_1a818a94-ba6f-413b-a7e8-6f8fd2a11e32",
														ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}},
													},
													Name: "Vacation Approval",
												},
												Incoming: []string{"_d79c991e-446c-47d1-ac9d-9d0113e35b93"},
												Outgoing: []string{"_ac1d82df-9bf6-404b-be70-b052e311f8b7"},
											},
											IoSpecification: instance.IoSpecification{
												InputOutputSpecification: instance.InputOutputSpecification{
													DataInputs: []instance.DataInput{
														{
															BaseElement: instance.BaseElement{
																ID:                "DataInput__1a818a94-ba6f-413b-a7e8-6f8fd2a11e32_1",
																ExtensionElements: instance.ExtensionElements{},
															},
															Name: "",
														},
														{
															BaseElement: instance.BaseElement{
																ID:                "DataInput__1a818a94-ba6f-413b-a7e8-6f8fd2a11e32",
																ExtensionElements: instance.ExtensionElements{},
															},
															Name: "",
														},
													},
													DataOutputs: []instance.DataOutput{
														{
															BaseElement: instance.BaseElement{
																ID: "DataOutput__1a818a94-ba6f-413b-a7e8-6f8fd2a11e32_1",
															},
															Name: "",
														},
														{
															BaseElement: instance.BaseElement{
																ID: "DataOutput__1a818a94-ba6f-413b-a7e8-6f8fd2a11e32",
															},
															Name: "",
														},
													},
													InputSets: []instance.InputSet{
														{
															BaseElement: instance.BaseElement{
																ID: "_7522b1b6-fa6b-4e86-8bd7-d77f8d0db55a",
															},
															DataInputRefs: []string{"DataInput__1a818a94-ba6f-413b-a7e8-6f8fd2a11e32_1", "DataInput__1a818a94-ba6f-413b-a7e8-6f8fd2a11e32"},
														},
													},
													OutputSets: []instance.OutputSet{
														{
															BaseElement: instance.BaseElement{
																ID: "_7821d757-2e54-4340-8286-6b138f27ca51",
															},
															DataOutputRefs: []string{"DataOutput__1a818a94-ba6f-413b-a7e8-6f8fd2a11e32_1", "DataOutput__1a818a94-ba6f-413b-a7e8-6f8fd2a11e32"},
														},
													},
												},
											},
											DataInputAssociations: []instance.DataInputAssociation{
												{
													DataAssociation: instance.DataAssociation{
														BaseElement: instance.BaseElement{
															ID:                "_d7aa7a34-bb93-49ae-87d7-f3ccd25add3f",
															ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil, nil}},
														},
														SourceRef: "dataInputdataInput_2bed4da9-f468-4fe0-a37c-ae5d4c17e323",
														TargetRef: "DataInput__1a818a94-ba6f-413b-a7e8-6f8fd2a11e32_1",
													},
												},
												{
													DataAssociation: instance.DataAssociation{
														BaseElement: instance.BaseElement{
															ID:                "_bfe6b00e-a1ce-4b12-b868-f445ebc77d1d",
															ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil, nil}},
														},
														SourceRef: "dataInputdataInput_103fd816-b422-420d-aec9-c28b54f620db",
														TargetRef: "DataInput__1a818a94-ba6f-413b-a7e8-6f8fd2a11e32",
													},
												},
											},
											DataOutputAssociations: []instance.DataOutputAssociation{
												{
													DataAssociation: instance.DataAssociation{
														BaseElement: instance.BaseElement{
															ID:                "_3dcf2a50-5a30-4a47-956f-7962eb907737",
															ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil, nil}},
														},
														SourceRef: "DataOutput__1a818a94-ba6f-413b-a7e8-6f8fd2a11e32_1",
														TargetRef: "dataOutputdataOutput_c2d108b2-e6c0-47e6-a3cb-aa5321bcaeaf",
													},
												},
												{
													DataAssociation: instance.DataAssociation{
														BaseElement: instance.BaseElement{
															ID:                "_66f6f24d-1cad-4c93-b9b3-b93193015a18",
															ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil, nil}},
														},
														SourceRef: "DataOutput__1a818a94-ba6f-413b-a7e8-6f8fd2a11e32",
														TargetRef: "dataOutputdataOutput_dac8ee76-f637-4fd9-8357-6a87fd11ef41",
													},
												},
											},
										},
									},
									Implementation: "##unspecified",
								},
							},
							ExclusiveGatewaies: []instance.ExclusiveGateway{
								{
									Gateway: instance.Gateway{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID:                "_42367c5f-d084-44ee-90c7-960d1ab02a3b",
													ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil, nil, nil}},
												},
												Name: "",
											},
											Incoming: []string{"_ac1d82df-9bf6-404b-be70-b052e311f8b7"},
											Outgoing: []string{"_2a5ee73c-c495-48a0-945b-3dabab28ca91", "_325973e7-0bc8-4136-b6df-be1e681d8608", "_0a1c4f20-509f-4aeb-baf9-acc762f4fdf9"},
										},
										GatewayDirection: "Diverging",
									},
									Default: "_2a5ee73c-c495-48a0-945b-3dabab28ca91",
								},
								{
									Gateway: instance.Gateway{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID:                "_64bb8b55-d348-41d6-9ea8-f333b1d8cb69",
													ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil, nil, nil}},
												},
												Name: "",
											},
											Incoming: []string{"_3c50d7bb-5ecc-4d34-82a9-a5fa41d9c018"},
											Outgoing: []string{"_d3a575d8-cd41-4246-ad7a-7a068de6679f", "_f2b0da63-d841-4457-ad85-7d86c8b5c1d2"},
										},
										GatewayDirection: "Diverging",
									},
									Default: "_d3a575d8-cd41-4246-ad7a-7a068de6679f",
								},
							},
							SendTasks: []instance.SendTask{
								{
									Task: instance.Task{
										Activity: instance.Activity{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID:                "_a97c1a48-faba-447b-bfa6-7aa81a6fe0a0",
														ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}},
													},
													Name: "Notify Employee of Approval",
												},
												Incoming: []string{"_f2b0da63-d841-4457-ad85-7d86c8b5c1d2"},
												Outgoing: []string{"_a5441c65-06f8-4972-8158-e6f61f904841"},
											},
										},
									},
									Implementation: "##WebService",
								},
								{
									Task: instance.Task{
										Activity: instance.Activity{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID:                "_02232e32-c3d2-473c-a15d-9c5dca00eadc",
														ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}},
													},
													Name: "Notify Employee of Refusal",
												},
												Incoming: []string{"_d3a575d8-cd41-4246-ad7a-7a068de6679f"},
												Outgoing: []string{"_93ac178c-3efc-48ae-8a99-8d29a0c4aaf6"},
											},
										},
									},
									Implementation: "##WebService",
								},
								{
									Task: instance.Task{
										Activity: instance.Activity{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID:                "_9ed61a6a-7cc1-4ed1-86d8-03482b0983c9",
														ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}},
													},
													Name: "Notify Employee of Refusal",
												},
												Incoming: []string{"_2a5ee73c-c495-48a0-945b-3dabab28ca91"},
												Outgoing: []string{"_130d903a-f211-4b3e-9711-9061f5978d94"},
											},
										},
									},
									Implementation: "##WebService",
								},
								{
									Task: instance.Task{
										Activity: instance.Activity{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID:                "_93ec9873-edf1-4549-b052-961994ec8234",
														ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}},
													},
													Name: "Notify Employee of Approval",
												},
												Incoming: []string{"_325973e7-0bc8-4136-b6df-be1e681d8608"},
												Outgoing: []string{"_188d38f6-b73d-4c62-964b-affc6fe040c1"},
											},
										},
									},
									Implementation: "##WebService",
								},
							},
							UserTasks: []instance.UserTask{
								{
									Task: instance.Task{
										Activity: instance.Activity{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID:                "_79523269-7444-4b01-90e9-e23957a9d020",
														ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil, nil, nil, nil, nil, nil, nil}},
													},
													Name: "Manually Approve Vacation",
												},
												Incoming: []string{"_0a1c4f20-509f-4aeb-baf9-acc762f4fdf9"},
												Outgoing: []string{"_3c50d7bb-5ecc-4d34-82a9-a5fa41d9c018"},
											},
										},
									},
									Implementation: "##unspecified",
								},
							},
							EndEvents: []instance.EndEvent{
								{
									ThrowEvent: instance.ThrowEvent{
										Event: instance.Event{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID:                "_1cd5fe29-b3ec-4f21-a1aa-57773f0729ca",
														ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil, nil, nil, nil}},
													},
													Name: "Vacation Approved by Manager",
												},
												Incoming: []string{"_81ed1b89-5a4d-4028-a245-4fd38b866b6d"},
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
														ID:                "_6677ef80-82df-4951-919d-1f36123b681b",
														ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil, nil, nil, nil}},
													},
													Name: "Vacation Approved Automatically",
												},
												Incoming: []string{"_68ef27a7-011e-4ec0-ad45-41f52093de68"},
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
														ID:                "_3ae826ca-5f38-43c4-be3a-35d1157aa27f",
														ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil, nil, nil, nil}},
													},
													Name: "Vacation Refused by Manager",
												},
												Incoming: []string{"_93ac178c-3efc-48ae-8a99-8d29a0c4aaf6"},
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
														ID:                "_b4d636eb-b501-4462-93c8-04652db10307",
														ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil, nil, nil, nil}},
													},
													Name: "Employee not found",
												},
												Incoming: []string{"_f250e125-23eb-45cd-9c61-191d172df093"},
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
														ID:                "_1688f604-5edf-4187-ad9e-18f74bfede53",
														ExtensionElements: instance.ExtensionElements{Any: []xml.Token{nil, nil, nil, nil}},
													},
													Name: "Vacation Refused Automatically",
												},
												Incoming: []string{"_130d903a-f211-4b3e-9711-9061f5978d94"},
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
