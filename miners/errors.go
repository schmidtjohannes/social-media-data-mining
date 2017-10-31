package miners

import "fmt"

type EmptyResultError struct {
	Name string
}

func (err EmptyResultError) Error() string {
	return fmt.Sprintf("The group returns empty result: %s", err.Name)
}

type HttpStatusNokError struct {
	Message string
	Code    int
}

func (err HttpStatusNokError) Error() string {
	return fmt.Sprintf("Response returned code %d with message: %s", err.Code, err.Message)
}
