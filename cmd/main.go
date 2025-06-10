package main

import (
	"log"
	"os"

	"github.com/Talos-hub/json-conv2YAML/pkg/converter"
	"github.com/Talos-hub/json-conv2YAML/pkg/flags"
	"github.com/Talos-hub/json-conv2YAML/pkg/reader"
	"github.com/Talos-hub/json-conv2YAML/pkg/wr"
)

func main() {

	//parse
	fl := flags.Parse()

	//for json
	var data []byte
	var err error

	//geted json
	if fl.Input == "" {
		data, err = reader.ReadFromStdin()
	} else {
		data, err = reader.ReadFile(fl.Input)
	}
	handleError(err)

	//marshal
	yamls, err := converter.ConvertJsonToYaml(data)

	handleError(err)

	//write
	if fl.Output == "" {
		_, err = wr.Write(yamls, os.Stdout)
	} else {
		_, err = wr.WriteToFile(yamls, fl.Output)
	}

	handleError(err)

}

// handleError cathes error
func handleError(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
