package flags

import (
	cli "github.com/jawher/mow.cli"
	"github.com/khorevaa/r2gitsync/context"
)

type FlagType int

const (
	ArgType FlagType = iota
	OptType
)

type command interface {
	String(p cli.StringParam) *string
	StringPtr(into *string, p cli.StringParam)

	Bool(p cli.BoolParam) *bool
	BoolPtr(into *bool, p cli.BoolParam)

	Int(p cli.IntParam) *int
	IntPtr(into *int, p cli.IntParam)

	Float64(p cli.Float64Param) *float64
	Float64Ptr(into *float64, p cli.Float64Param)

	Strings(p cli.StringsParam) *[]string
	StringsPtr(into *[]string, p cli.StringsParam)
}

type Flag interface {
	Apply(cmd command, ctx context.Context)
}

//func IntFlag(name, value, desc string) IntFlag {
//	return IntFlag{
//		IntFlag: cli.IntFlag{
//			Name:  name,
//			Value: value,
//			Desc:  desc,
//		},
//	}
//}
