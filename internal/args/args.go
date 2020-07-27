package args

import (
	cli "github.com/jawher/mow.cli"
	"strings"
)

type stringArg struct {
	cli.StringArg
}

type cmd interface {
	String(p cli.StringParam) *string
	Bool(p cli.BoolParam) *bool
	Int(p cli.IntParam) *int
	Float64(p cli.Float64Param) *float64
	Strings(p cli.StringsParam) *[]string
}

func (o stringArg) Env(envs ...string) stringArg {

	newO := o
	newO.EnvVar = strings.Join(envs, " ")
	return newO

}

func (o stringArg) Default(value string) stringArg {

	newO := o
	newO.Value = value
	return newO

}

func (o stringArg) HideValue(hide bool) stringArg {

	newO := o
	newO.StringArg.HideValue = hide
	return newO

}

func (o stringArg) Desc(desc string) stringArg {

	newO := o
	newO.StringArg.Desc = desc
	return newO

}

func (o stringArg) Arg(cmd cmd) *string {

	return cmd.String(o)

}

func StringArg(name, value, desc string) stringArg {
	return stringArg{
		StringArg: cli.StringArg{
			Name:  name,
			Value: value,
			Desc:  desc,
		}}
}
