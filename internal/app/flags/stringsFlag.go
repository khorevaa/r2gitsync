package flags

import (
	cli "github.com/jawher/mow.cli"
	"github.com/khorevaa/r2gitsync/pkg/context"
	"strings"
)

type StringsFlag struct {
	FlagType    FlagType
	Name        string
	Desc        string
	EnvVar      string
	HideValue   bool
	SetByUser   *bool
	Value       []string
	Destination *[]string
}

func (o StringsFlag) Apply(cmd command, ctx context.Context) {

	if o.Destination == nil {
		o.Destination = new([]string)
	}
	ctx.Set(o.Name, o.Destination)

	cmd.StringsPtr(o.Destination, getStringsFlag(o))

}

func (o StringsFlag) Ptr(into *[]string) StringsFlag {

	newO := o
	newO.Destination = into
	return newO
}

func (o StringsFlag) Env(envs ...string) StringsFlag {

	newO := o
	newO.EnvVar = strings.Join(envs, " ")
	return newO

}

func getStringsFlag(o StringsFlag) cli.StringsParam {

	switch o.FlagType {

	case OptType:
		return cli.StringsOpt{
			Name:      o.Name,
			Desc:      o.Desc,
			EnvVar:    o.EnvVar,
			Value:     o.Value,
			HideValue: o.HideValue,
			SetByUser: o.SetByUser,
		}
	case ArgType:
		return cli.StringsArg{
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

func StringsArg(Name string, Value []string, Desc string) StringsFlag {

	return StringsFlag{
		Name:     Name,
		Desc:     Desc,
		Value:    Value,
		FlagType: ArgType,
	}

}

func StringsOpt(Name string, Value []string, Desc string) StringsFlag {

	return StringsFlag{
		Name:     Name,
		Desc:     Desc,
		Value:    Value,
		FlagType: OptType,
	}

}
