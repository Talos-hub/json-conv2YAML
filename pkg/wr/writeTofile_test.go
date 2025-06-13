package wr

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestWriteToFile(t *testing.T) {

	//arrange
	var hello []byte = []byte("Hello")

	nameTemp := "tempfile"
	tempFile, err := os.CreateTemp("", nameTemp)
	if err != nil {
		t.Error(err)
	}

	defer tempFile.Close()

	t.Cleanup(func() {
		os.Remove(tempFile.Name()) // Cleanup after test
	})

	//act
	_, err = WriteToFile(hello, tempFile.Name())
	if err != nil {
		t.Error(err)
	}

	data, err := Read(tempFile.Name())
	if err != nil {
		t.Error(err)
	}
	//assert
	if string(data) != string(hello) {
		t.Error(err)
	}

}

func Read(path string) ([]byte, error) {

	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("file is not exist: %w", err)
		}
		return nil, err
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return data, nil
}
