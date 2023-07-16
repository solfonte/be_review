package utils

import (
	"fmt"
	"io/ioutil"
	"os"
)

type FileReader struct{}

func (fileReader *FileReader) ReadFile(filepath string) ([]byte, error) {
	file, err := os.Open(filepath)
	var byteValue []byte
	if err != nil {
		fmt.Println("An error occured while reading the match file with path ", filepath, ": ", err)
		return nil, err
	}
	byteValue, _ = ioutil.ReadAll(file)
	file.Close()
	return byteValue, err
}

func (fileReader *FileReader) WriteFile(filepath string, jsonData []byte) {
	file, err := os.Create(filepath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	_, err = file.Write(jsonData)
	file.Close()
}