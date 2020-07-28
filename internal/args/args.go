package args

import (
	cli "github.com/jawher/mow.cli"
	"strings"
)

type stringArg struct {
	cmd command
	cli.StringArg
}

type command interface {
	String(p cli.StringParam) *string
	StringPtr(into *string, p cli.StringParam)
	Bool(p cli.BoolParam) *bool
	BoolPtr(into *bool, p cli.BoolParam)
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

func (o stringArg) Arg() *string {

	return o.cmd.String(o.StringArg)

}

func (o stringArg) Ptr(into *string) {

	o.cmd.StringPtr(into, o.StringArg)

}

func StringArg(cmd command, name, value, desc string) stringArg {
	return stringArg{
		cmd: cmd,
		StringArg: cli.StringArg{
			Name:  name,
			Value: value,
			Desc:  desc,
		}}
}
