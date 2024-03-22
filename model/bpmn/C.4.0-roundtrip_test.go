package bpmn

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sparrow-community/bpm/model/bpmn/instance"
)

func TestC_4_0_roundtrip(t *testing.T) {
	// create test use ./test/C.4.0-roundtrip.bpmn
	path := "./test/C.4.0-roundtrip.bpmn"
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
			ID:              "_e7ccc12c-2987-48ae-92aa-08112101d9c9",
			Name:            "C.4.0",
			TargetNamespace: "http://www.trisotech.com/definitions/_e7ccc12c-2987-48ae-92aa-08112101d9c9",
			Exporter:        "BPMN Modeler",
			ExporterVersion: "6.2.0",
			RootElemnts: instance.RootElemnts{
				ItemDefinitions: []instance.ItemDefinition{
					{
						RootElement: instance.RootElement{
							BaseElement: instance.BaseElement{
								ID: "_triso-default-bpmnItemDefinition-string_id",
							},
						},
						StructureRef: "string",
					},
				},
				Signals: []instance.Signal{
					{
						RootElement: instance.RootElement{
							BaseElement: instance.BaseElement{
								ID: "_bf41cec3-da5e-4fc7-87f9-7520a720192b",
							},
						},
						Name: "New employee hired",
					},
				},
				Collaborations: []instance.Collaboration{
					{
						RootElement: instance.RootElement{
							BaseElement: instance.BaseElement{
								ID: "_085241a5-fb5c-44d3-8844-5366f865e353",
							},
						},
						Participants: []instance.Participant{
							{
								BaseElement: instance.BaseElement{
									ID: "_06745dd2-cbc6-4742-b11c-69281e5dcbdb",
								},
								Name:       "Money Bank",
								ProcessRef: "_42cba3a9-a8ab-40b5-b9a4-2e8f32be364e",
							},
						},
					},
					{
						RootElement: instance.RootElement{
							BaseElement: instance.BaseElement{
								ID: "_674b5898-d454-4114-a8eb-66c790366e94",
							},
						},
						Participants: []instance.Participant{
							{
								BaseElement: instance.BaseElement{
									ID: "_b5172765-a1c1-4c6e-8e57-9eabaf8cecdf",
								},
								Name:       "IT",
								ProcessRef: "_f0035388-f829-470c-b82b-0b15c3da3399",
							},
						},
					},
					{
						RootElement: instance.RootElement{
							BaseElement: instance.BaseElement{
								ID: "_bbdf487b-c3a3-4bfa-8ba3-79be59f08203",
							},
						},
						Participants: []instance.Participant{
							{
								BaseElement: instance.BaseElement{
									ID: "_1422af8d-518b-49f1-9563-9b6939de2818",
								},
								Name:       "Payroll",
								ProcessRef: "_da743a6f-d9e5-4fcf-8a96-d2fd5cfb73d4",
							},
						},
					},
					{
						RootElement: instance.RootElement{
							BaseElement: instance.BaseElement{
								ID: "_8d4bb408-4175-4fbb-8131-18c32b7aa022",
							},
						},
						Participants: []instance.Participant{
							{
								BaseElement: instance.BaseElement{
									ID: "_fab40827-7808-470a-93b6-4a3318bf0c0e",
								},
								Name:       "Facilities",
								ProcessRef: "_3486bf55-0a7f-4ff1-be15-1555669f58ad",
							},
						},
					},
				},
				DataStores: []instance.DataStore{
					{
						RootElement: instance.RootElement{
							BaseElement: instance.BaseElement{
								ID: "_d367181e-c25a-407e-a5fb-817a27698a30",
							},
						},
						ItemSubjectRef: "_triso-default-bpmnItemDefinition-string_id",
						Name:           "Employee Details",
						IsUnlimited:    false,
					},
					{
						RootElement: instance.RootElement{
							BaseElement: instance.BaseElement{
								ID: "_7c793f30-d450-4a8e-becf-f168b14317ea",
							},
						},
						Name:        "User Management",
						IsUnlimited: false,
					},
					{
						RootElement: instance.RootElement{
							BaseElement: instance.BaseElement{
								ID: "_2addcc9a-638e-4496-b84f-5b0e28fe6536",
							},
						},
						Name:        "Payroll system",
						IsUnlimited: false,
					},
				},
				Processes: []instance.Process{
					{
						CallableElement: instance.CallableElement{
							RootElement: instance.RootElement{
								BaseElement: instance.BaseElement{
									ID: "_42cba3a9-a8ab-40b5-b9a4-2e8f32be364e",
								},
							},
							Name: "Money Bank - Process",
						},
						IsClosed: false,
						LaneSets: []instance.LaneSet{
							{
								BaseElement: instance.BaseElement{
									ID: "_ff7ff8f6-a4f1-4e93-84e1-01cdb85eb755",
								},
								Name: "HR Department",
								Lanes: []instance.Lane{
									{
										BaseElement: instance.BaseElement{
											ID: "Lane_0w7h48m",
										},
										Name: "HR Department",
										FlowNodeRefs: []string{
											"_a4220c17-364f-4a08-ae9c-757a6468b295",
											"_f8973a92-3d84-4672-a1a3-b0df154121e1",
											"_f9e3cd76-809a-48b5-be1c-e84fc4324268",
											"_aa275782-c989-49ba-bf94-c58916ca7bb5",
											"_987b9b74-333a-4043-a72a-daadf667acc7",
											"_305ddf53-49a8-4105-ad06-70272a2332aa",
											"_0e71ed63-93f9-44b6-a89d-da9628652926",
											"_eba690b9-34ef-49e4-b265-1411809d9302",
											"_67944b4c-4950-45a2-a131-1c4679c6b433",
											"_4c95f4a0-f4ec-45ed-9fdb-7b236155d6f5",
											"_82da02ca-ee9a-4403-9f3b-aad030e089b9",
										},
									},
									{
										BaseElement: instance.BaseElement{
											ID: "_937b5086-463f-4c8c-837d-f5eee5cbc1f4",
										},
										Name: "Responsible Department",
										FlowNodeRefs: []string{
											"_986cf801-0780-49d3-91cd-2cc6d3c1aac3",
											"_855451b0-5298-48b2-a81d-84ecbcca0a85",
											"_72da5cee-0456-4c3c-ba8d-6dd085d6f52d",
											"_e3d3ac43-74a3-48ff-9a02-e64b1358cc34",
											"_80f70d22-fb42-403f-8bdb-6805e9467bb7",
											"_fe77c2f2-278f-4752-9d03-aa0c8a12af1e",
											"_db9147a9-7fbc-4657-a506-15e777f2cfd9",
											"_74e2cc7b-99ca-426b-ad53-ad70a56506aa",
											"_19808f32-dfb5-462d-aaa6-e662f9932dba",
											"_351b058e-c37c-4fb7-9d32-24075f53ce02",
											"_52401cbb-02b8-4eaf-84f1-1edbc0854a4a",
											"_36baf139-fb74-43ef-8936-d490238c2825",
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
														ID: "_a4220c17-364f-4a08-ae9c-757a6468b295",
													},
													Name: "Candidate accepted offer",
												},
												Outgoing: []string{"_6d1a9d5a-127d-477a-8dbe-268f125f20d8"},
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
														ID: "_f8973a92-3d84-4672-a1a3-b0df154121e1",
													},
													Name: "Send &#10;candidate Contract",
												},
												Incoming: []string{"_6d1a9d5a-127d-477a-8dbe-268f125f20d8", "_9ca6c9e6-4aff-4ec8-9fd3-2743c06483c3"},
												Outgoing: []string{"_4d3df0ce-e5be-443b-a379-8dcd4aae4de9"},
											},
										},
									},
									Implementation: "##unspecified",
								},
								{
									Task: instance.Task{
										Activity: instance.Activity{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "_aa275782-c989-49ba-bf94-c58916ca7bb5",
													},
													Name: "Get signature on contract and notify responsible department",
												},
												Incoming: []string{"_237c8380-5449-446e-a323-aad80181176d"},
												Outgoing: []string{"_41254af9-ca6b-49ae-a58e-c48f4c6120d2"},
											},
											IoSpecification: instance.IoSpecification{
												InputOutputSpecification: instance.InputOutputSpecification{
													BaseElement: instance.BaseElement{
														ID: "_d0960287-a964-4916-b4fe-b1238d9e3e75",
													},
													DataOutputs: []instance.DataOutput{
														{
															BaseElement: instance.BaseElement{
																ID: "_d0960287-a964-4916-b4fe-b1238d9e3e75",
															},
															ItemSubjectRef: "_triso-default-bpmnItemDefinition-string_id",
															Name:           "Employee Details",
														},
													},
													InputSets: []instance.InputSet{},
													OutputSets: []instance.OutputSet{
														{
															BaseElement: instance.BaseElement{
																ID: "_2850d02b-5efe-4fa8-82c9-eac458b2e4d8",
															},
															DataOutputRefs: []string{"_d0960287-a964-4916-b4fe-b1238d9e3e75"},
														},
													},
												},
											},
											DataOutputAssociations: []instance.DataOutputAssociation{
												{
													DataAssociation: instance.DataAssociation{
														BaseElement: instance.BaseElement{
															ID: "_5c5f0068-874e-4dbb-ad30-ae8389921645",
														},
														SourceRef: "_d0960287-a964-4916-b4fe-b1238d9e3e75",
														TargetRef: "_f00a95f6-043e-4ec1-8731-153a707e950c",
													},
												},
											},
										},
									},
									Implementation: "##unspecified",
								},
								{
									Task: instance.Task{
										Activity: instance.Activity{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "_987b9b74-333a-4043-a72a-daadf667acc7",
													},
													Name: "Review terms of contract",
												},
												Incoming: []string{"_7e9d8b8b-faa9-4264-858b-7454702c4ec2"},
												Outgoing: []string{"_9ca6c9e6-4aff-4ec8-9fd3-2743c06483c3"},
											},
										},
									},
									Implementation: "##unspecified",
								},
								{
									Task: instance.Task{
										Activity: instance.Activity{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "_0e71ed63-93f9-44b6-a89d-da9628652926",
													},
													Name: "Inform employee of company policies",
												},
												Incoming: []string{"_d1c906b1-8f68-4485-94ea-00e732fba73e"},
												Outgoing: []string{"_5ae80bdd-cda7-491b-b3f7-bce2c4b35a83"},
											},
										},
									},
									Implementation: "##unspecified",
								},
								{
									Task: instance.Task{
										Activity: instance.Activity{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "_eba690b9-34ef-49e4-b265-1411809d9302",
													},
													Name: "Introduce employee to company Mission, Vision and Values",
												},
												Incoming: []string{"_5ae80bdd-cda7-491b-b3f7-bce2c4b35a83"},
												Outgoing: []string{"_15f3cde0-eeb5-4ad2-820a-114815aa782d"},
											},
										},
									},
									Implementation: "##unspecified",
								},
								{
									Task: instance.Task{
										Activity: instance.Activity{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "_67944b4c-4950-45a2-a131-1c4679c6b433",
													},
													Name: "Perform training for time reports sick leave and holidays",
												},
												Incoming: []string{"_15f3cde0-eeb5-4ad2-820a-114815aa782d"},
												Outgoing: []string{"_2d93c5db-2e41-463b-a1a6-63a05616d912"},
											},
										},
									},
									Implementation: "##unspecified",
								},
								{
									Task: instance.Task{
										Activity: instance.Activity{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "_4c95f4a0-f4ec-45ed-9fdb-7b236155d6f5",
													},
													Name: "Register for medical insurance",
												},
												Incoming: []string{"_2d93c5db-2e41-463b-a1a6-63a05616d912"},
												Outgoing: []string{"_3104b7bb-a417-4557-8c1b-d95b1c7e36a4"},
											},
										},
									},
									Implementation: "##unspecified",
								},
								{
									Task: instance.Task{
										Activity: instance.Activity{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "_986cf801-0780-49d3-91cd-2cc6d3c1aac3",
													},
													Name: "Request preparations for a new employee",
												},
												Incoming: []string{"_0811a8cc-1fde-4319-a718-7fb3ce64ae0b"},
												Outgoing: []string{"_3bb98d19-0d0b-4f47-a097-30c60183e334"},
											},
										},
									},
									Implementation: "##unspecified",
								},
								{
									Task: instance.Task{
										Activity: instance.Activity{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "_72da5cee-0456-4c3c-ba8d-6dd085d6f52d",
													},
													Name: "Introduce new employee to the team",
												},
												Incoming: []string{"_cad69ba8-8b26-41f9-87a1-372a899dcbf3"},
												Outgoing: []string{"_2ef25e1e-bbdd-4cba-b411-6091ec8669d6"},
											},
										},
									},
									Implementation: "##unspecified",
								},
								{
									Task: instance.Task{
										Activity: instance.Activity{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "_e3d3ac43-74a3-48ff-9a02-e64b1358cc34",
													},
													Name: "Perform training for position",
												},
												Incoming: []string{"_2ef25e1e-bbdd-4cba-b411-6091ec8669d6"},
												Outgoing: []string{"_240df34c-86c2-49fe-b1ec-58ea25d35c1f"},
											},
										},
									},
									Implementation: "##unspecified",
								},
								{
									Task: instance.Task{
										Activity: instance.Activity{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "_351b058e-c37c-4fb7-9d32-24075f53ce02",
													},
													Name: "Compile welcome package",
												},
												Incoming: []string{"_ece7bc06-ccf7-49f5-b5f1-99e1103cf538"},
												Outgoing: []string{"_c8454490-8090-4cf7-a164-d0aded7bc4da"},
											},
										},
									},
									Implementation: "##unspecified",
								},
								{
									Task: instance.Task{
										Activity: instance.Activity{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "_52401cbb-02b8-4eaf-84f1-1edbc0854a4a",
													},
													Name: "Give employee welcome package",
												},
												Incoming: []string{"_c8454490-8090-4cf7-a164-d0aded7bc4da"},
												Outgoing: []string{"_0fdd6987-6236-4972-a7be-3431ed6654ec"},
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
													ID: "_f9e3cd76-809a-48b5-be1c-e84fc4324268",
												},
												Name: "Contract terms accepted ?",
											},
											Incoming: []string{"_4d3df0ce-e5be-443b-a379-8dcd4aae4de9"},
											Outgoing: []string{"_237c8380-5449-446e-a323-aad80181176d", "_7e9d8b8b-faa9-4264-858b-7454702c4ec2"},
										},
										GatewayDirection: "Diverging",
									},
								},
							},
							ParallelGatewaies: []instance.ParallelGateway{
								{
									Gateway: instance.Gateway{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "_305ddf53-49a8-4105-ad06-70272a2332aa",
												},
												Name: "Non-exclusive Gateway",
											},
											Incoming: []string{"_41254af9-ca6b-49ae-a58e-c48f4c6120d2"},
											Outgoing: []string{"_d1c906b1-8f68-4485-94ea-00e732fba73e", "_0811a8cc-1fde-4319-a718-7fb3ce64ae0b"},
										},
										GatewayDirection: "Diverging",
									},
								},
								{
									Gateway: instance.Gateway{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "_82da02ca-ee9a-4403-9f3b-aad030e089b9",
												},
												Name: "Non-exclusive Gateway",
											},
											Incoming: []string{"_3104b7bb-a417-4557-8c1b-d95b1c7e36a4", "_71b2c3f1-0d34-4cfb-956c-cedf83660dfb"},
											Outgoing: []string{"_cad69ba8-8b26-41f9-87a1-372a899dcbf3"},
										},
										GatewayDirection: "Converging",
									},
								},
								{
									Gateway: instance.Gateway{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "_80f70d22-fb42-403f-8bdb-6805e9467bb7",
												},
												Name: "Non-exclusive Gateway",
											},
											Incoming: []string{"_240df34c-86c2-49fe-b1ec-58ea25d35c1f"},
											Outgoing: []string{"_2be942a3-d921-4ad9-bce7-33e3411d98b4", "_4b186d86-fae1-4f55-a3b6-b90fa27affc1", "_c1307108-5b3d-48af-ae53-11b30dd31bed"},
										},
										GatewayDirection: "Diverging",
									},
								},
								{
									Gateway: instance.Gateway{
										FlowNode: instance.FlowNode{
											FlowElement: instance.FlowElement{
												BaseElement: instance.BaseElement{
													ID: "_19808f32-dfb5-462d-aaa6-e662f9932dba",
												},
												Name: "Non-exclusive Gateway",
											},
											Incoming: []string{"_e82e6ee3-24d6-419a-96c0-a147f1943c23", "_16e4630e-39c0-4e94-ab65-a8a622245adf", "_c8da08ab-0b59-4c21-9386-f857d624c471"},
											Outgoing: []string{"_ece7bc06-ccf7-49f5-b5f1-99e1103cf538"},
										},
										GatewayDirection: "Converging",
									},
								},
							},
							DataStoreReferences: []instance.DataStoreReference{
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_f00a95f6-043e-4ec1-8731-153a707e950c",
										},
										Name: "Employee Details",
									},
									DataStoreRef: "_d367181e-c25a-407e-a5fb-817a27698a30",
								},
							},
							IntermediateThrowEvents: []instance.IntermediateThrowEvent{
								{
									ThrowEvent: instance.ThrowEvent{
										Event: instance.Event{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "_855451b0-5298-48b2-a81d-84ecbcca0a85",
													},
													Name: "New employee in department X",
												},
												Incoming: []string{"_3bb98d19-0d0b-4f47-a097-30c60183e334"},
												Outgoing: []string{"_71b2c3f1-0d34-4cfb-956c-cedf83660dfb"},
											},
										},
										DataInputs: []instance.DataInput{
											{
												BaseElement: instance.BaseElement{
													ID: "_d5083b64-f0cd-479b-a773-b477a20408b5",
												},
												ItemSubjectRef: "_triso-default-bpmnItemDefinition-string_id",
												Name:           "Employee Details",
											},
										},
										DataInputAssociations: []instance.DataInputAssociation{
											{
												DataAssociation: instance.DataAssociation{
													BaseElement: instance.BaseElement{
														ID: "_7a5b52f3-ad7e-4c5c-9ac2-432c8318b821",
													},
													SourceRef: "_f00a95f6-043e-4ec1-8731-153a707e950c",
													TargetRef: "_d5083b64-f0cd-479b-a773-b477a20408b5",
												},
											},
										},
										InputSet: instance.InputSet{
											BaseElement: instance.BaseElement{
												ID: "_e304d0f8-6eaa-4a79-a6ff-b08e993d0d06",
											},
											DataInputRefs: []string{"_d5083b64-f0cd-479b-a773-b477a20408b5"},
										},
										EventDefinitions: instance.EventDefinitions{
											SignalEventDefinitions: []instance.SignalEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "_e027a3d0-02f7-40a6-b84c-3b49282070ba",
															},
														},
													},
													SignalRef: "_bf41cec3-da5e-4fc7-87f9-7520a720192b",
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
														ID: "_fe77c2f2-278f-4752-9d03-aa0c8a12af1e",
													},
													Name: "Input from Payroll ready",
												},
												Incoming: []string{"_4b186d86-fae1-4f55-a3b6-b90fa27affc1"},
												Outgoing: []string{"_e82e6ee3-24d6-419a-96c0-a147f1943c23"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											MessageEventDefinitions: []instance.MessageEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "_9e9bd4fc-0a94-4956-ac98-daf2ddeb001c",
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
														ID: "_db9147a9-7fbc-4657-a506-15e777f2cfd9",
													},
													Name: "Input from Facilities ready",
												},
												Incoming: []string{"_c1307108-5b3d-48af-ae53-11b30dd31bed"},
												Outgoing: []string{"_c8da08ab-0b59-4c21-9386-f857d624c471"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											MessageEventDefinitions: []instance.MessageEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "_3454f552-d717-4f8c-ae58-971bef888af0",
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
														ID: "_74e2cc7b-99ca-426b-ad53-ad70a56506aa",
													},
													Name: "Input from IT ready",
												},
												Incoming: []string{"_2be942a3-d921-4ad9-bce7-33e3411d98b4"},
												Outgoing: []string{"_16e4630e-39c0-4e94-ab65-a8a622245adf"},
											},
										},
										EventDefinitions: instance.EventDefinitions{
											MessageEventDefinitions: []instance.MessageEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "_940a88a5-974a-4583-afed-4510a94e5249",
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
														ID: "_36baf139-fb74-43ef-8936-d490238c2825",
													},
													Name: "End Event",
												},
												Incoming: []string{"_0fdd6987-6236-4972-a7be-3431ed6654ec"},
											},
										},
									},
								},
							},
							SequenceFlows: []instance.SequenceFlow{
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_6d1a9d5a-127d-477a-8dbe-268f125f20d8",
										},
									},
									SourceRef: "_a4220c17-364f-4a08-ae9c-757a6468b295",
									TargetRef: "_f8973a92-3d84-4672-a1a3-b0df154121e1",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_4d3df0ce-e5be-443b-a379-8dcd4aae4de9",
										},
									},
									SourceRef: "_f8973a92-3d84-4672-a1a3-b0df154121e1",
									TargetRef: "_f9e3cd76-809a-48b5-be1c-e84fc4324268",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_237c8380-5449-446e-a323-aad80181176d",
										},
										Name: "Yes",
									},
									SourceRef: "_f9e3cd76-809a-48b5-be1c-e84fc4324268",
									TargetRef: "_aa275782-c989-49ba-bf94-c58916ca7bb5",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_7e9d8b8b-faa9-4264-858b-7454702c4ec2",
										},
										Name: "No",
									},
									SourceRef: "_f9e3cd76-809a-48b5-be1c-e84fc4324268",
									TargetRef: "_987b9b74-333a-4043-a72a-daadf667acc7",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_41254af9-ca6b-49ae-a58e-c48f4c6120d2",
										},
									},
									SourceRef: "_aa275782-c989-49ba-bf94-c58916ca7bb5",
									TargetRef: "_305ddf53-49a8-4105-ad06-70272a2332aa",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_d1c906b1-8f68-4485-94ea-00e732fba73e",
										},
									},
									SourceRef: "_305ddf53-49a8-4105-ad06-70272a2332aa",
									TargetRef: "_0e71ed63-93f9-44b6-a89d-da9628652926",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_5ae80bdd-cda7-491b-b3f7-bce2c4b35a83",
										},
									},
									SourceRef: "_0e71ed63-93f9-44b6-a89d-da9628652926",
									TargetRef: "_eba690b9-34ef-49e4-b265-1411809d9302",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_15f3cde0-eeb5-4ad2-820a-114815aa782d",
										},
									},
									SourceRef: "_eba690b9-34ef-49e4-b265-1411809d9302",
									TargetRef: "_67944b4c-4950-45a2-a131-1c4679c6b433",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_2d93c5db-2e41-463b-a1a6-63a05616d912",
										},
									},
									SourceRef: "_67944b4c-4950-45a2-a131-1c4679c6b433",
									TargetRef: "_4c95f4a0-f4ec-45ed-9fdb-7b236155d6f5",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_3104b7bb-a417-4557-8c1b-d95b1c7e36a4",
										},
									},
									SourceRef: "_4c95f4a0-f4ec-45ed-9fdb-7b236155d6f5",
									TargetRef: "_82da02ca-ee9a-4403-9f3b-aad030e089b9",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_0811a8cc-1fde-4319-a718-7fb3ce64ae0b",
										},
									},
									SourceRef: "_305ddf53-49a8-4105-ad06-70272a2332aa",
									TargetRef: "_986cf801-0780-49d3-91cd-2cc6d3c1aac3",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_3bb98d19-0d0b-4f47-a097-30c60183e334",
										},
									},
									SourceRef: "_986cf801-0780-49d3-91cd-2cc6d3c1aac3",
									TargetRef: "_855451b0-5298-48b2-a81d-84ecbcca0a85",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_71b2c3f1-0d34-4cfb-956c-cedf83660dfb",
										},
									},
									SourceRef: "_855451b0-5298-48b2-a81d-84ecbcca0a85",
									TargetRef: "_82da02ca-ee9a-4403-9f3b-aad030e089b9",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_cad69ba8-8b26-41f9-87a1-372a899dcbf3",
										},
									},
									SourceRef: "_82da02ca-ee9a-4403-9f3b-aad030e089b9",
									TargetRef: "_72da5cee-0456-4c3c-ba8d-6dd085d6f52d",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_2ef25e1e-bbdd-4cba-b411-6091ec8669d6",
										},
									},
									SourceRef: "_72da5cee-0456-4c3c-ba8d-6dd085d6f52d",
									TargetRef: "_e3d3ac43-74a3-48ff-9a02-e64b1358cc34",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_240df34c-86c2-49fe-b1ec-58ea25d35c1f",
										},
									},
									SourceRef: "_e3d3ac43-74a3-48ff-9a02-e64b1358cc34",
									TargetRef: "_80f70d22-fb42-403f-8bdb-6805e9467bb7",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_2be942a3-d921-4ad9-bce7-33e3411d98b4",
										},
									},
									SourceRef: "_80f70d22-fb42-403f-8bdb-6805e9467bb7",
									TargetRef: "_74e2cc7b-99ca-426b-ad53-ad70a56506aa",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_4b186d86-fae1-4f55-a3b6-b90fa27affc1",
										},
									},
									SourceRef: "_80f70d22-fb42-403f-8bdb-6805e9467bb7",
									TargetRef: "_fe77c2f2-278f-4752-9d03-aa0c8a12af1e",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_c1307108-5b3d-48af-ae53-11b30dd31bed",
										},
									},
									SourceRef: "_80f70d22-fb42-403f-8bdb-6805e9467bb7",
									TargetRef: "_db9147a9-7fbc-4657-a506-15e777f2cfd9",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_e82e6ee3-24d6-419a-96c0-a147f1943c23",
										},
									},
									SourceRef: "_fe77c2f2-278f-4752-9d03-aa0c8a12af1e",
									TargetRef: "_19808f32-dfb5-462d-aaa6-e662f9932dba",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_16e4630e-39c0-4e94-ab65-a8a622245adf",
										},
									},
									SourceRef: "_74e2cc7b-99ca-426b-ad53-ad70a56506aa",
									TargetRef: "_19808f32-dfb5-462d-aaa6-e662f9932dba",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_c8da08ab-0b59-4c21-9386-f857d624c471",
										},
									},
									SourceRef: "_db9147a9-7fbc-4657-a506-15e777f2cfd9",
									TargetRef: "_19808f32-dfb5-462d-aaa6-e662f9932dba",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_ece7bc06-ccf7-49f5-b5f1-99e1103cf538",
										},
									},
									SourceRef: "_19808f32-dfb5-462d-aaa6-e662f9932dba",
									TargetRef: "_351b058e-c37c-4fb7-9d32-24075f53ce02",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_c8454490-8090-4cf7-a164-d0aded7bc4da",
										},
									},
									SourceRef: "_351b058e-c37c-4fb7-9d32-24075f53ce02",
									TargetRef: "_52401cbb-02b8-4eaf-84f1-1edbc0854a4a",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_0fdd6987-6236-4972-a7be-3431ed6654ec",
										},
									},
									SourceRef: "_52401cbb-02b8-4eaf-84f1-1edbc0854a4a",
									TargetRef: "_36baf139-fb74-43ef-8936-d490238c2825",
								},
								{
									FlowElement: instance.FlowElement{
										BaseElement: instance.BaseElement{
											ID: "_9ca6c9e6-4aff-4ec8-9fd3-2743c06483c3",
										},
									},
									SourceRef: "_987b9b74-333a-4043-a72a-daadf667acc7",
									TargetRef: "_f8973a92-3d84-4672-a1a3-b0df154121e1",
								},
							},
						},
					},
					{
						CallableElement: instance.CallableElement{
							RootElement: instance.RootElement{
								BaseElement: instance.BaseElement{
									ID: "_f0035388-f829-470c-b82b-0b15c3da3399",
								},
							},
							Name: "IT - Process",
						},
						IsClosed: false,
						FlowElements: instance.FlowElements{
							StartEvents: []instance.StartEvent{
								{
									CatchEvent: instance.CatchEvent{
										Event: instance.Event{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "_e9306b3f-3a77-42e1-b53e-2ed8ee45486d",
													},
													Name: "New &#10;employee&#10;hired",
												},
												Outgoing: []string{"_495f04a4-ed17-4e7b-a952-11c06c2230d1"},
											},
										},
										DataOutputs: []instance.DataOutput{
											{
												BaseElement: instance.BaseElement{
													ID: "_e96a6b5c-b24a-4221-8d35-cd3a4e68641e",
												},
												ItemSubjectRef: "_triso-default-bpmnItemDefinition-string_id",
												Name:           "Employee Details",
											},
										},
										DataOutputAssociations: []instance.DataOutputAssociation{
											{
												DataAssociation: instance.DataAssociation{
													BaseElement: instance.BaseElement{
														ID: "_ed210028-7993-4aa4-b2f7-951ab85243be",
													},
													SourceRef: "_e96a6b5c-b24a-4221-8d35-cd3a4e68641e",
													TargetRef: "_e355b511-e181-4e7b-a9b1-30cb48dd6cab",
												},
											},
										},
										OutputSet: instance.OutputSet{
											BaseElement: instance.BaseElement{
												ID: "_2d74cb72-2fcc-4869-b5f0-b867a7f2f7ec",
											},
											OptionalOutputRefs: []string{"_e96a6b5c-b24a-4221-8d35-cd3a4e68641e"},
										},
										EventDefinitions: instance.EventDefinitions{
											SignalEventDefinitions: []instance.SignalEventDefinition{
												{
													EventDefinition: instance.EventDefinition{
														RootElement: instance.RootElement{
															BaseElement: instance.BaseElement{
																ID: "_5ec0186c-9f7d-43b1-ba20-571cfe42bc5c",
															},
														},
													},
													SignalRef: "_bf41cec3-da5e-4fc7-87f9-7520a720192b",
												},
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
														ID: "_7e9d2e5a-21f7-493b-9ae4-03245aa33a5c",
													},
													Name: "Create domain account",
												},
												Incoming: []string{"_495f04a4-ed17-4e7b-a952-11c06c2230d1"},
												Outgoing: []string{"_0bccc7f6-0f56-49bc-98df-59d70433c043"},
											},
											IoSpecification: instance.IoSpecification{
												InputOutputSpecification: instance.InputOutputSpecification{
													DataInputs: []instance.DataInput{
														{
															BaseElement: instance.BaseElement{
																ID: "_b1d40770-825b-4d23-a026-db95f815d520",
															},
															ItemSubjectRef: "_triso-default-bpmnItemDefinition-string_id",
														},
													},
													DataOutputs: []instance.DataOutput{
														{
															BaseElement: instance.BaseElement{
																ID: "_18dcd171-c602-488e-b53b-0289e2fae114",
															},
															Name: "User Management",
														},
													},
													InputSets: []instance.InputSet{
														{
															BaseElement: instance.BaseElement{
																ID: "_c228336b-c661-4e3a-920f-8e13f4844656",
															},
															DataInputRefs: []string{"_b1d40770-825b-4d23-a026-db95f815d520"},
														},
													},
													OutputSets: []instance.OutputSet{
														{
															BaseElement: instance.BaseElement{
																ID: "_3c50eff0-7da3-42bf-96be-d7e22f4a2bb5",
															},
															DataOutputRefs: []string{"_18dcd171-c602-488e-b53b-0289e2fae114"},
														},
													},
												},
											},
											DataInputAssociations: []instance.DataInputAssociation{
												{
													DataAssociation: instance.DataAssociation{
														BaseElement: instance.BaseElement{
															ID: "_c4eff1b8-8817-418b-9ef9-5d3738d43659",
														},
														SourceRef: "_e355b511-e181-4e7b-a9b1-30cb48dd6cab",
														TargetRef: "_b1d40770-825b-4d23-a026-db95f815d520",
													},
												},
											},
											DataOutputAssociations: []instance.DataOutputAssociation{
												{
													DataAssociation: instance.DataAssociation{
														BaseElement: instance.BaseElement{
															ID: "_0282dbd7-a11b-4cac-84e2-0f26bf86b853",
														},
														SourceRef: "_18dcd171-c602-488e-b53b-0289e2fae114",
														TargetRef: "_b60be020-1f04-46c2-b1a9-19bfc66ead82",
													},
												},
											},
										},
									},
									Implementation: "##unspecified",
								},
							},
							ManualTasks: []instance.ManualTask{
								{
									Task: instance.Task{
										Activity: instance.Activity{
											FlowNode: instance.FlowNode{
												FlowElement: instance.FlowElement{
													BaseElement: instance.BaseElement{
														ID: "_c29af228-0768-4dfe-945a-17755e173674",
													},
													Name: "Prepare workstation",
												},
												Incoming: []string{"_0bccc7f6-0f56-49bc-98df-59d70433c043"},
												Outgoing: []string{"_dcec1142-90eb-41a8-b809-458b713ab94c"},
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
