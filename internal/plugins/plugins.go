package plugins

import (
	"github.com/khorevaa/r2gitsync/internal/plugins/limit"
	"github.com/khorevaa/r2gitsync/pkg/plugin"
)

func init() {

	plugin.Register(limit.Symbol)

}
