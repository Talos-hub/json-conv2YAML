package converter

import (
	"encoding/json"
	"fmt"

	conv "gopkg.in/yaml.v3"
)

// JsonToYaml converte json to yaml
func JsonToYaml(j []byte) ([]byte, error) {
	var obj interface{}

	if err := json.Unmarshal(j, &obj); err != nil {
		return nil, fmt.Errorf("error parsing json: %w", err)
	}

	//marshaling to yaml
	dataYaml, err := conv.Marshal(obj)

	if err != nil {
		return nil, err
	}

	return dataYaml, nil

}
