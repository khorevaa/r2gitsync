package context

import (
	"context"
	"strings"
	"time"
)

type flagSet map[string]interface{}

type ctx struct {
	context.Context
	flagSet       flagSet //*flag.FlagSet
	parentContext *Context
}

type Context interface {
	context.Context

	Set(name string, value interface{})
	Duration(name string) time.Duration
	Bool(name string) bool
	Int(name string) int
	String(name string) string
}

func NewContext() Context {

	return &ctx{
		flagSet: make(flagSet),
	}

}

func (c *ctx) Set(name string, value interface{}) {

	for _, n := range strings.Fields(name) {
		c.flagSet[n] = value
	}

}
