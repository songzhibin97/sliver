package boot

import (
	"fmt"
	"github.com/SliverHorn/sliver/global"
	"github.com/SliverHorn/sliver/integration/server"
	"go.uber.org/zap"
)

func init() {
	Router := Routers()
	Router.Static("/form-generator", "./static/page")

	address := fmt.Sprintf(":%d", global.Config.System.Addr)
	s := server.Init(address, Router)
	global.Zap.Debug("server run success on ", zap.String("address", address))

	fmt.Printf(`欢迎使用 github.com/SliverHorn/sliver
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:8080
`, address)
	global.Zap.Error(s.ListenAndServe().Error())
}