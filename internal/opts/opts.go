package opts

import (
	cli "github.com/jawher/mow.cli"
	"strings"
)

type command interface {
	String(p cli.StringParam) *string
	StringPtr(into *string, p cli.StringParam)
	Bool(p cli.BoolParam) *bool
	BoolPtr(into *bool, p cli.BoolParam)
	Int(p cli.IntParam) *int
	IntPtr(into *int, p cli.IntParam)
	Float64(p cli.Float64Param) *float64
	Strings(p cli.StringsParam) *[]string
}

type stringOpt struct {
	cmd command
	cli.StringOpt
}

//type stringOpt cli.StringOpt //{
////	cli.StringOpt
////}

func (o stringOpt) Cmd(cmd command) stringOpt {

	newO := o
	newO.cmd = cmd
	return newO

}

func (o stringOpt) Env(envs ...string) stringOpt {

	newO := o
	newO.EnvVar = strings.Join(envs, " ")
	return newO

}

func (o stringOpt) Default(value string) stringOpt {

	newO := o
	newO.StringOpt.Value = value
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

func (o stringOpt) Opt() *string {

	return o.cmd.String(o.StringOpt)

}

func (o stringOpt) Ptr(into *string) {

	o.cmd.StringPtr(into, o.StringOpt)

}

func StringOpt(cmd command, name, value, desc string) stringOpt {
	return stringOpt{
		cmd: cmd,
		StringOpt: cli.StringOpt{
			Name:  name,
			Value: value,
			Desc:  desc,
		},
	}
}

type boolOpt struct {
	cmd command
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

func (o boolOpt) Opt() *bool {

	return o.cmd.Bool(o.BoolOpt)

}

func (o boolOpt) Ptr(into *bool) {

	o.cmd.BoolPtr(into, o.BoolOpt)

}

func BoolOpt(cmd command, name string, value bool, desc string) boolOpt {
	return boolOpt{
		cmd: cmd,
		BoolOpt: cli.BoolOpt{
			Name:  name,
			Value: value,
			Desc:  desc,
		},
	}
}

type intOpt struct {
	cmd command
	cli.IntOpt
}

func (o intOpt) Env(envs ...string) intOpt {

	newO := o
	newO.EnvVar = strings.Join(envs, " ")
	return newO

}

func (o intOpt) Default(value int) intOpt {

	newO := o
	newO.Value = value
	return newO

}

func (o intOpt) HideValue(hide bool) intOpt {

	newO := o
	newO.IntOpt.HideValue = hide
	return newO

}

func (o intOpt) Desc(desc string) intOpt {

	newO := o
	newO.IntOpt.Desc = desc
	return newO

}

func (o intOpt) Opt() *int {

	return o.cmd.Int(o.IntOpt)

}

func (o intOpt) Ptr(into *int) {

	o.cmd.IntPtr(into, o.IntOpt)

}

func IntOpt(cmd command, name string, value int, desc string) intOpt {
	return intOpt{
		cmd: cmd,
		IntOpt: cli.IntOpt{
			Name:  name,
			Value: value,
			Desc:  desc,
		},
	}
}
