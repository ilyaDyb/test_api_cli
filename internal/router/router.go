package router

import "github.com/ilyaDyb/internal/functions"

func CliCommandsRouter(value string) (string, error) {
	msgs, err := functions.GetBaseJsonMessages()
	if err != nil {
		return "", err
	}
	switch value {
	case "help":
		return msgs.HelpMessage, nil
	default:
		return "Unknown command", nil
	}
}