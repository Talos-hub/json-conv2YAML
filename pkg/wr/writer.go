package wr

import (
	"fmt"
	"os"
	"strings"
)

// WriteToFile Write data to file
func WriteToFile(b []byte, path string) (int, error) {

	path = strings.TrimSpace(path)

	//Creating a file where we can write a data
	file, err := os.Create(path)

	if err != nil {
		if os.IsExist(err) {
			return 0, fmt.Errorf("file already exists at %q: %w", path, err)
		}
		return 0, fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	//write data to file
	n, err := file.Write(b)

	if err != nil {
		return n, fmt.Errorf("error writing file: %w", err)
	}

	return n, nil
}

// WriteToStdout write data to a file.
// By default we write to stdout
func WriteToDestination(b []byte) (int, error) {
	n, err := os.Stdout.Write(b)

	if err != nil {
		return n, fmt.Errorf("error writing file: %w", err)
	}

	return n, nil
}
