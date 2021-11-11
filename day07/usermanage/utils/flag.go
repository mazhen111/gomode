package utils

import (
	"flag"
	"fmt"
	"os"
)

func Flag() {
	var init bool
	flag.BoolVar(&init, "init", false, "初始化")
	flag.Parse()
	if init {
		//永久保存文件
		Load()
		fmt.Println("初始化程序...")
		//定义初始化存储设备
		InitPersist()
		//注册admin管理员
		InitAdminUser()
		persist := GetPersist()
		persist.Save()
		os.Exit(0)

	}

}
