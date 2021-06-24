package main

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"time"
)

func main() {
	gdb.SetConfig(gdb.Config{
		"default": gdb.ConfigGroup{
			gdb.ConfigNode{
				Host: "192.168.230.129",
				Port: "3306",
				User: "root",
				Pass: "123",
				Name: "gf-demo",
				Type: "mysql",
			},
		},
	})
	db := g.DB()
	// 开启调试模式，以便于记录所有执行的SQL
	db.SetDebug(true)

	// 执行2次查询并将查询结果缓存1小时，并可执行缓存名称(可选) ,输出结果只有一条sql说明缓冲生效
	for i := 0; i < 2; i++ {
		r, _ := db.Table("user").Cache(time.Hour, "vip-user").Where("id", 1).One()
		g.Log().Print(r.Map())
	}

	// 执行更新操作，并清理指定名称的查询缓存
	_, err := db.Table("user").Cache(-1, "vip-user").Data(gdb.Map{"nickname": "smith"}).Where("id", 1).Update()
	if err != nil {
		g.Log().Fatal(err)
	}

	// 再次执行查询，启用查询缓存特性
	r, _ := db.Table("user").Cache(time.Hour, "vip-user").Where("id", 1).One()
	g.Log().Print(r.Map())
}
