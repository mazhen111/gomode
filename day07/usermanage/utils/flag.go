package utils

import "flag"

func Flag() {
	var init bool
	flag.BoolVar(&init, "init", false, "初始化")
	flag.Parse()
	if init {

	}

}
