package utils

import (
	"strconv"
	"strings"

	"github.com/ilyaDyb/internal/requester"
)

func ParseCliInput(args []string) (*requester.Requester, error) {
	var (
		url string
		method string
		concurrency int
		requests int
	)
	for index, arg := range args {
		switch arg {
		case "--url":
			url = args[index+1]
		case "--method":
			method = strings.ToUpper(args[index+1])
		case "--concurency":
			intVal, err := strconv.Atoi(args[index+1])
			if err != nil {
				return nil, err
			}
			concurrency = intVal
		case "--requests":
			intVal, err := strconv.Atoi(args[index+1])
			if err != nil {
				return nil, err
			}
			concurrency = intVal
		}
	}
	requester := requester.Requester{Url: url, Method: method, Concurrency: concurrency, Requests: requests}

	return &requester, nil
}

// func ParseReportCommand(args []string) ()