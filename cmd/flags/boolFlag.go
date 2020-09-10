package flags

import (
	cli "github.com/jawher/mow.cli"
	"github.com/khorevaa/r2gitsync/context"
	"strings"
)

type BoolFlag struct {
	FlagType    FlagType
	Name        string
	Desc        string
	EnvVar      string
	HideValue   bool
	SetByUser   *bool
	Value       bool
	Destination *bool
}

func (o BoolFlag) Apply(cmd command, ctx context.Context) {

	if o.Destination != nil {
		o.Destination = new(bool)
	}
	ctx.Set(o.Name, o.Destination)

	cmd.BoolPtr(o.Destination, getBoolFlag(o))

}

func (o BoolFlag) Ptr(into *bool) BoolFlag {

	newO := o
	newO.Destination = into
	return newO
}

func (o BoolFlag) Env(envs ...string) BoolFlag {

	newO := o
	newO.EnvVar = strings.Join(envs, " ")
	return newO

}

func getBoolFlag(o BoolFlag) cli.BoolParam {

	switch o.FlagType {

	case OptType:
		return cli.BoolOpt{
			Name:      o.Name,
			Desc:      o.Desc,
			EnvVar:    o.EnvVar,
			Value:     o.Value,
			HideValue: o.HideValue,
			SetByUser: o.SetByUser,
		}
	case ArgType:
		return cli.BoolArg{
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

func BoolArg(Name string, Value bool, Desc string) BoolFlag {

	return BoolFlag{
		Name:     Name,
		Desc:     Desc,
		Value:    Value,
		FlagType: ArgType,
	}

}

func BoolOpt(Name string, Value bool, Desc string) BoolFlag {

	return BoolFlag{
		Name:     Name,
		Desc:     Desc,
		Value:    Value,
		FlagType: OptType,
	}

}
