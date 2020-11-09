package main

import (
	_ "github.com/SliverHorn/sliver/boot"
	"github.com/SliverHorn/sliver/global"
)

func main() {
	db, _ := global.Db.DB()
	defer db.Close() // 程序结束前关闭数据库链接
}