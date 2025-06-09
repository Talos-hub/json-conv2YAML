package reader

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/Talos-hub/json-conv2YAML/pkg/flags"
)

// ReadFile reads of file and returns data
func ReadFile(flag *flags.Flag) ([]byte, error) {

	//by default we read data from os.Stdin
	if flag.Input == "" {
		fmt.Println("Type json and press Enter")

		//crete reader
		reader := bufio.NewReader(os.Stdin)

		data, err := reader.ReadBytes('\n')

		if err != nil {
			return nil, fmt.Errorf("error reading os.stdin: %w", err)
		}

		return data, err
	}

	//Open file
	file, err := os.Open(flag.Input)

	//catch a error
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("file is not exist: %w", err)
		}
		return nil, fmt.Errorf("error opening file: %w", err)
	}

	defer file.Close()

	//reads
	data, err := io.ReadAll(file)

	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return data, nil
}
