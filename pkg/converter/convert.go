package converter

import (
	"encoding/json"
	"fmt"

	conv "gopkg.in/yaml.v3"
)

// JsonToYaml converte json to yaml
func ConvertJsonToYaml(j []byte) ([]byte, error) {

	if len(j) == 0 {
		return nil, fmt.Errorf("empty JSON input")
	}
	//for json unmarshaling
	var obj interface{}

	if err := json.Unmarshal(j, &obj); err != nil {
		return nil, fmt.Errorf("error parsing json: %w", err)
	}

	//marshaling to yaml
	dataYaml, err := conv.Marshal(obj)

	if err != nil {
		return nil, fmt.Errorf("error marshaling to yaml: %w", err)
	}

	return dataYaml, nil

}
