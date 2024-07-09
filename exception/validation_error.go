package exception

import "strings"

type ValidationError struct {
	ListOfError []string
}

func (validationError ValidationError) Error() string {
	return strings.Join(validationError.ListOfError, ",")
}
