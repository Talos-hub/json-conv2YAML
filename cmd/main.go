package main

import (
	"fmt"
	"os"

	"github.com/Talos-hub/json-conv2YAML/pkg/converter"
	"github.com/Talos-hub/json-conv2YAML/pkg/flags"
	"github.com/Talos-hub/json-conv2YAML/pkg/reader"
	"github.com/Talos-hub/json-conv2YAML/pkg/wr"
)

func main() {

	//parse falgs
	fl := flags.Parse()

	//reads data
	data, err := reader.ReadFile(fl)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//convert
	yaml, err := converter.JsonToYaml(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, err = wr.Write(yaml, fl)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
