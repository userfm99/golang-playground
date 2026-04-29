package main

import (
	"errors"
	"strings"
)

// StringService test
type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

type stringService struct{}

func (stringService) Uppercase(v string) (string, error) {
	if v == "" {
		return "", errors.New("empty string")
	}
	return strings.ToUpper(v), nil
}

func (stringService) Count(v string) int {
	return len(v)
}

type upperCaseRequest struct {
	S string `json:"s"`
}
type upperCaseResponse struct {
	S   string `json:"s"`
	Err error  `json:"err, omitempty"`
}
type countRequest struct {
	S string
}
type countResponse struct {
	V int
}

func main() {

}
