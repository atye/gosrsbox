package gosrsbox

import "fmt"

// Error contains a code and message returned from the API if something went wrong
type serverError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err *serverError) Error() string {
	return fmt.Sprintf("code: %d, message: %s", err.Code, err.Message)
}
