package instance

import (
	"encoding/xml"
	"os"
	"testing"
)

func TestDefinitions(t *testing.T) {
	xmlData, err := os.ReadFile("./BpmnDiTest.xml")
	if nil != err {
		t.Fatalf("err: %v", err)
	}
	definitions := &Definitions{}
	if err := xml.Unmarshal(xmlData, definitions); err != nil {
		t.Fatalf("err: %v", err)
	}
	t.Logf("definitions: %v", definitions)
}
