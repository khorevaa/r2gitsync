package plugin

import (
	cli "github.com/jawher/mow.cli"
)

type command interface {
	String(p cli.StringParam) *string
	StringPtr(into *string, p cli.StringParam)
	Bool(p cli.BoolParam) *bool
	BoolPtr(into *bool, p cli.BoolParam)
	Int(p cli.IntParam) *int
	Float64(p cli.Float64Param) *float64
	Strings(p cli.StringsParam) *[]string
	IntPtr(into *int, p cli.IntParam)
	Float64Ptr(into *float64, p cli.Float64Param)
	StringsPtr(into *[]string, p cli.StringsParam)
}

type subscriber struct {
	handlers []interface{}
}

func (s subscriber) Handlers() []interface{} {
	return s.handlers
}
