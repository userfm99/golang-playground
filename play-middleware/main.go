package main

import (
	"errors"
	"log"
)

type (
	// HandlerFunc ---
	HandlerFunc func(string) error

	// MiddlewareFunc ---
	MiddlewareFunc func(HandlerFunc) HandlerFunc
)

// TestMiddleware ---
type TestMiddleware struct {
	preMiddleWare []MiddlewareFunc
}

// NewTestMiddleware ---
func NewTestMiddleware() *TestMiddleware {
	return &TestMiddleware{}
}

// Pre ---
func (t *TestMiddleware) Pre(middlewareFuncs ...MiddlewareFunc) {
	t.preMiddleWare = append(t.preMiddleWare, middlewareFuncs...)
}

func (t *TestMiddleware) Test(value string, handler HandlerFunc) error {
	h := handler
	for i := len(t.preMiddleWare) - 1; i >= 0; i-- {
		h = t.preMiddleWare[i](h)
	}
	return h(value)
}

// PrintValue ---
func PrintValue(value string) error {
	log.Println(value)
	return nil
}

func main() {
	t := NewTestMiddleware()
	t.Pre(middlewareTrimString())
	t.Test("F A D L Y M U N A N D A R", PrintValue)
}

func middlewareTrimString() MiddlewareFunc {
	return func(next HandlerFunc) HandlerFunc {
		return func(v string) error {
			if len(v) == 0 {
				return errors.New("empty string")
			}
			return next("xxxxxx")
		}
	}
}
