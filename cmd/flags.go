package cmd

import (
	cli "github.com/jawher/mow.cli"
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

type StringOpt struct {
	Name        string
	Desc        string
	EnvVar      string
	HideValue   bool
	SetByUser   *bool
	Value       string
	Destination *string
}

type Flag interface {
	Apply(cmd command)
}

func (o StringOpt) Apply(cmd command) {

	cmd.StringPtr(o.Destination, cli.StringOpt{
		Name:      o.Name,
		Desc:      o.Desc,
		EnvVar:    o.EnvVar,
		Value:     o.Value,
		HideValue: o.HideValue,
		SetByUser: o.SetByUser,
	})

}

//func StringOpt(name, value, desc string) StringOpt {
//	return StringOpt{
//		StringOpt: cli.StringOpt{
//			Name:  name,
//			Value: value,
//			Desc:  desc,
//		},
//	}
//}
