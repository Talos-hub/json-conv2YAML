package wr

import (
	"fmt"
	"os"

	"github.com/Talos-hub/json-conv2YAML/pkg/flags"
)

func Write(b []byte, flags *flags.Flag) (int, error) {

	//by default we write to stdout
	if flags.Output == "" {
		n, err := os.Stdout.Write(b)

		if err != nil {
			return n, fmt.Errorf("error writing file: %w", err)
		}

		return n, nil
	}

	//Creating a file where we can write a data
	file, err := os.Create(flags.Output)

	if err != nil {
		return 0, fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	//write data to file
	n, err := file.Write(b)

	if err != nil {
		return n, fmt.Errorf("error wirting file: %w", err)
	}

	return n, nil
}
