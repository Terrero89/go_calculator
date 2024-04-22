//file that manages the file system operation such as open and close files while reading it

package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// NewFileManager struct
//takes an input file

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

// ReadLines responsible for reading the content of text and returning the text
func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	//opening file
	if err != nil {
		fmt.Println("Could not  open:", err)
		return nil, errors.New("Could not open file")
	}
	//reading line by line content with scanner
	scanner := bufio.NewScanner(file) //scanner line by line

	var lines []string //slice for each line
	//for loop to append each line
	for scanner.Scan() {
		lines = append(lines, scanner.Text())

	}
	//store error in already initialized err variable
	err = scanner.Err()

	if err != nil {
		file.Close()
		return nil, errors.New("Failed to read lines in file path:")
	}

	file.Close()
	return lines, nil
}

// WriteJSON will receive any type as a parameter for data,and string for the path

func (fm FileManager) WriteResult(data interface{}) error {
	//if not found, create a new file, if exists then overwrite it
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {

		return errors.New("Failed to create file")
	}

	//create a new file in JSON format

	encoder := json.NewEncoder(file) //encodes the valuess in file in json .
	err = encoder.Encode(data)       // store value of file and encode them
	if err != nil {

		return errors.New("Failed to convert data to json")
	}

	file.Close()
	return nil
}

func New(inputFile, outputFile string) FileManager {
	return FileManager{
		InputFilePath:  inputFile,
		OutputFilePath: outputFile,
	}
}
