package boot

import (
	"github.com/SliverHorn/sliver/global"
	"github.com/gogf/gf/i18n/gi18n"
)

func InitI18n() {
	global.I18n = gi18n.New()
}
