package plugins

import (
	"github.com/khorevaa/r2gitsync/plugin"
	"github.com/khorevaa/r2gitsync/plugins/limit"
)

func init() {

	//log.SetDebug()

	//err := plugin.Register(limit.NewPlugin)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//err = plugin.Register(plugin.NewPlugin("test", "1.0.0+f8sd8fa", "test plugins", func() plugin.Plugin {
	//	return &limit.LimitPlugin{}
	//}, plugin.WithModule("init", "sync")))
	//
	//if err != nil {
	//	fmt.Println(err)
	//}

}

var Plugins = []plugin.Symbol{
	limit.NewPlugin,
}
