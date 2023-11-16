package bpmn

import (
	"encoding/xml"
	"io"
	"os"

	"github.com/sparrow-community/bpm/model/bpmn/instance"
)

type BpmnModelInstance struct {
	Definitions *instance.Definitions
}

// BpmnModelInstanceFromBytes reads a BPMN model instance from a byte array
func BpmnModelInstanceFromBytes(data []byte) (*BpmnModelInstance, error) {
	definitions := &instance.Definitions{}
	err := xml.Unmarshal(data, definitions)
	if err != nil {
		return nil, err
	}

	return &BpmnModelInstance{
		Definitions: definitions,
	}, nil
}

// BpmnModelInstanceFromFile reads a BPMN model instance from a file
func BpmnModelInstanceFromFile(filePath string) (*BpmnModelInstance, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return BpmnModelInstanceFromBytes(data)
}

// BpmnModelInstanceFromStream reads a BPMN model instance from a stream
func BpmnModelInstanceFromStream(stream io.Reader) (*BpmnModelInstance, error) {
	data, err := io.ReadAll(stream)
	if err != nil {
		return nil, err
	}
	return BpmnModelInstanceFromBytes(data)
}
