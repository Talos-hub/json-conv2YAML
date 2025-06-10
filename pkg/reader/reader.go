package reader

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// ReadFile reads of file and returns data
func ReadFile(path string) ([]byte, error) {

	path = strings.TrimSpace(path)

	//Open file
	file, err := os.Open(path)

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

// ReadFromStdin reads data from os.Stdin
func ReadFromStdin() ([]byte, error) {
	fmt.Println("Type json and press Enter")

	//crete reader
	reader := bufio.NewReader(os.Stdin)

	data, err := reader.ReadBytes('\n')

	if err != nil {
		return nil, fmt.Errorf("error reading os.stdin: %w", err)
	}

	return data, nil
}
