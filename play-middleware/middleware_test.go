package main

import "testing"

func TestMiddlewarex(t *testing.T) {
	m := NewTestMiddleware()
	m.Pre(middlewareTrimString())
	err := m.Test("F A D L Y M U N A N D A R", PrintValue)
	if err != nil {
		t.Error(err)
	}
}
