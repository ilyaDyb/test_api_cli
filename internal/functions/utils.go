package functions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Message struct {
	HelloMessage string `json:"hello_message"`
	HelpMessage  string `json:"help_message"`
}

func GetBaseJsonMessages() (*Message, error) {
	var err error
	jsonFile, err := os.Open("internal/otherfiles/instructions.json")
	if err != nil {
		err = fmt.Errorf("error opening file: %v", err)
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		err = fmt.Errorf("error reading file: %v", err.Error())
		return nil, err
	}

	var msg Message
	err = json.Unmarshal(byteValue, &msg)
	if err != nil {
		err = fmt.Errorf("error unmarshaling JSON: %v", err.Error())
		return nil, err
	}
	return &msg, nil
}