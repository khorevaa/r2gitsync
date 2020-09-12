package plugins

import (
	"fmt"
	"github.com/khorevaa/r2gitsync/plugin"
	"github.com/khorevaa/r2gitsync/plugins/limit"
)

func init() {

	err := plugin.Register(limit.NewPlugin)
	if err != nil {
		fmt.Println(err)
	}
	err = plugin.Register(plugin.NewPlugin("test", "1.0.0+f8sd8fa", "test plugins", func() plugin.Plugin {
		return &limit.LimitPlugin{}
	}))

	if err != nil {
		fmt.Println(err)
	}

}
