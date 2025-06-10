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

	//data from a file
	var data []byte
	var err error

	//reads from a file or a stdin
	if fl.Input == "" {
		data, err = reader.ReadFromStdin()
	} else {
		data, err = reader.ReadFile(fl.Input)
	}

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
