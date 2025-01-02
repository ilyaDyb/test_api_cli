package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ilyaDyb/internal/functions"
	"github.com/ilyaDyb/internal/utils"
	"github.com/ilyaDyb/internal/validators"
	// "github.com/ilyaDyb/internal/router"
)

func main() {
	msgs, err := functions.GetBaseJsonMessages()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(msgs.HelloMessage)
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error when read command", err)
			continue
		}
		input = strings.TrimSpace(input)

		if input == "exit" {
			return
		}
		splitedInput := strings.Split(input, " ")
		command := splitedInput[0]
		params, err := utils.ParseCliInput(splitedInput[1:])
		if err != nil {
			fmt.Println("Error when parsing input", err)
			continue
		}
		err = validators.ValidateCliInput(command, params)
		if err != nil {
			fmt.Println("Error when validating input", err)
			continue
		}
		// message, err := router.CliCommandsRouter(command)
		// if err != nil {
		// 	fmt.Println("Unknown error: ", err.Error())
		// 	return
		// }
		// fmt.Println(message)
	}
}