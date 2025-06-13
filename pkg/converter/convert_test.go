package converter

import (
	"encoding/json"
	"testing"

	"gopkg.in/yaml.v3"
)

type worker struct {
	name string
	age  int
}

func TestConvertYaml(t *testing.T) {
	//arange
	//obj
	person := worker{name: "some", age: 12}

	//json
	dataJson, err := json.Marshal(person)
	if err != nil {
		t.Error(err)
	}

	//yaml
	dataYaml, err := yaml.Marshal(person)
	if err != nil {
		t.Error(err)
	}

	//act
	data, err := ConvertJsonToYaml(dataJson)
	if err != nil {
		t.Error(err)
	}

	//asssert
	if string(dataYaml) != string(data) {
		t.Errorf("data: %s, want: %s", string(data), string(dataYaml))
	}

}
