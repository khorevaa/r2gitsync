package flags

import (
	cli "github.com/jawher/mow.cli"
	"github.com/khorevaa/r2gitsync/context"
	"strings"
)

type StringFlag struct {
	FlagType    FlagType
	Name        string
	Desc        string
	EnvVar      string
	HideValue   bool
	SetByUser   *bool
	Value       string
	Destination *string
}

func (o StringFlag) Apply(cmd command, ctx context.Context) {

	if o.Destination == nil {
		o.Destination = new(string)
	}
	ctx.Set(o.Name, o.Destination)

	cmd.StringPtr(o.Destination, getStringFlag(o))

}

func (o StringFlag) Ptr(into *string) StringFlag {

	newO := o
	newO.Destination = into
	return newO
}

func (o StringFlag) Env(envs ...string) StringFlag {

	newO := o
	newO.EnvVar = strings.Join(envs, " ")
	return newO

}

func getStringFlag(o StringFlag) cli.StringParam {

	switch o.FlagType {

	case OptType:
		return cli.StringOpt{
			Name:      o.Name,
			Desc:      o.Desc,
			EnvVar:    o.EnvVar,
			Value:     o.Value,
			HideValue: o.HideValue,
			SetByUser: o.SetByUser,
		}
	case ArgType:
		return cli.StringArg{
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

func StringArg(Name, Value, Desc string) StringFlag {

	return StringFlag{
		Name:     Name,
		Desc:     Desc,
		Value:    Value,
		FlagType: ArgType,
	}

}

func StringOpt(Name, Value, Desc string) StringFlag {

	return StringFlag{
		Name:     Name,
		Desc:     Desc,
		Value:    Value,
		FlagType: OptType,
	}

}
