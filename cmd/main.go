package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	
	"github.com/ilyaDyb/internal/functions"
	"github.com/ilyaDyb/internal/utils"
	"github.com/ilyaDyb/internal/validators"
	"github.com/ilyaDyb/internal/router"
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
			fmt.Println("Exiting...")
			return
		} else if input == "help" {
			fmt.Println(msgs.HelpMessage)
			continue
		} else if input == "mytests" {
			// in future
		}
		splitedInput := strings.Split(input, " ")
		command := splitedInput[0]
		args := splitedInput[1:]
		requester, err := utils.ParseCliInput(args)
		if err != nil {
			fmt.Println("Error when parsing input", err)
			continue
		}
		err = validators.ValidateCliInput(command, requester)
		if err != nil {
			fmt.Println("Error when validating input", err)
			continue
		}
		err = router.CliCommandsRouter(command, requester)
		if err != nil {
			fmt.Println("Unknown error: ", err.Error())
			return
		}
	}
}