package reader

import (
	"os"
	"testing"
)

func TestReaderBytes(t *testing.T) {
	//arange
	tempFile, err := os.CreateTemp("", "tempsfile")
	if err != nil {
		t.Error(err)
	}

	defer tempFile.Close()

	t.Cleanup(func() { os.Remove(tempFile.Name()) })

	var hello []byte = []byte("hello")

	//write to tempFile
	_, err = tempFile.Write(hello)
	if err != nil {
		t.Error(err)
	}

	//act
	//read from temp
	data, err := ReadFile(tempFile.Name())
	if err != nil {
		t.Error(err)
	}

	//assert
	if string(data) != string(hello) {
		t.Error()
	}

}
