package validators

import (
	"fmt"
	"slices"
	"strings"

	"github.com/ilyaDyb/internal/requester"
)

func ValidateCliInput(command string, requester *requester.Requester) error {
	if requester.Body != nil && requester.Method == "GET"{
		return fmt.Errorf("you can't use body with GET request")
	}
	if !strings.HasPrefix(requester.Url, "http://") || !strings.HasPrefix(requester.Url, "https://") {
		return fmt.Errorf("invalid url need 'http' or 'https'")
	}
	if requester.Body == nil && requester.Method != "GET" {
		return fmt.Errorf("you can't use POST, PUTCH and etc. without body")
	}
	if !slices.Contains([]string{"test"}, command) {
		return fmt.Errorf("unknown command")
	}
	return nil
}