package metadata

import (
	. "github.com/khorevaa/r2gitsync/pkg/plugin/types"
)

// Plugin основной интерфейс плагинов
type Plugin interface {
	Subscribe() Subscriber
}
