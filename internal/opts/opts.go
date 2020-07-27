package opts

import (
	cli "github.com/jawher/mow.cli"
	"strings"
)

type cmd interface {
	String(p cli.StringParam) *string
	Bool(p cli.BoolParam) *bool
	Int(p cli.IntParam) *int
	Float64(p cli.Float64Param) *float64
	Strings(p cli.StringsParam) *[]string
}

type stringOpt struct {
	cli.StringOpt
}

func (o stringOpt) Env(envs ...string) stringOpt {

	newO := o
	newO.EnvVar = strings.Join(envs, " ")
	return newO

}

func (o stringOpt) Default(value string) stringOpt {

	newO := o
	newO.Value = value
	return newO

}

func (o stringOpt) HideValue(hide bool) stringOpt {

	newO := o
	newO.StringOpt.HideValue = hide
	return newO

}

func (o stringOpt) Desc(desc string) stringOpt {

	newO := o
	newO.StringOpt.Desc = desc
	return newO

}

func (o stringOpt) Opt(cmd cmd) *string {

	return cmd.String(o)

}

type boolOpt struct {
	cli.BoolOpt
}

func (o boolOpt) Env(envs ...string) boolOpt {

	newO := o
	newO.EnvVar = strings.Join(envs, " ")
	return newO

}

func (o boolOpt) Default(value bool) boolOpt {

	newO := o
	newO.Value = value
	return newO

}

func (o boolOpt) HideValue(hide bool) boolOpt {

	newO := o
	newO.BoolOpt.HideValue = hide
	return newO

}

func (o boolOpt) Desc(desc string) boolOpt {

	newO := o
	newO.BoolOpt.Desc = desc
	return newO

}

func (o boolOpt) Opt(cmd cmd) *bool {

	return cmd.Bool(o)

}

func StringOpt(name, value, desc string) stringOpt {
	return stringOpt{
		StringOpt: cli.StringOpt{
			Name:  name,
			Value: value,
			Desc:  desc,
		},
	}
}

func BoolOpt(name string, value bool, desc string) boolOpt {
	return boolOpt{
		BoolOpt: cli.BoolOpt{
			Name:  name,
			Value: value,
			Desc:  desc,
		},
	}
}
