package flags

import (
	cli "github.com/jawher/mow.cli"
	"github.com/khorevaa/r2gitsync/pkg/context"
	"strings"
)

type IntFlag struct {
	FlagType    FlagType
	Name        string
	Desc        string
	EnvVar      string
	HideValue   bool
	SetByUser   *bool
	Value       int
	Destination *int
}

func (o IntFlag) Apply(cmd command, ctx context.Context) {

	if o.Destination == nil {
		o.Destination = new(int)
	}
	ctx.Set(o.Name, o.Destination)
	cmd.IntPtr(o.Destination, getIntFlag(o))

}

func (o IntFlag) Ptr(into *int) IntFlag {

	newO := o
	newO.Destination = into
	return newO
}

func (o IntFlag) Env(envs ...string) IntFlag {

	newO := o
	newO.EnvVar = strings.Join(envs, " ")
	return newO

}

func getIntFlag(o IntFlag) cli.IntParam {

	switch o.FlagType {

	case OptType:
		return cli.IntOpt{
			Name:      o.Name,
			Desc:      o.Desc,
			EnvVar:    o.EnvVar,
			Value:     o.Value,
			HideValue: o.HideValue,
			SetByUser: o.SetByUser,
		}
	case ArgType:
		return cli.IntArg{
			Name:      o.Name,
			Desc:      o.Desc,
			EnvVar:    o.EnvVar,
			Value:     o.Value,
			HideValue: o.HideValue,
			SetByUser: o.SetByUser,
		}
	default:
		panic("unknown flag type")
	}

}

func IntArg(Name string, Value int, Desc string) IntFlag {

	return IntFlag{
		Name:     Name,
		Desc:     Desc,
		Value:    Value,
		FlagType: ArgType,
	}

}

func IntOpt(Name string, Value int, Desc string) IntFlag {

	return IntFlag{
		Name:     Name,
		Desc:     Desc,
		Value:    Value,
		FlagType: OptType,
	}

}
