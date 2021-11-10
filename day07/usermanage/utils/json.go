package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
	"usermanage/model"
)

type JsonPersister struct {
}

//定义保存函数
func (this JsonPersister) Save() {
	timeNow := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("data/%v.gob", timeNow)
	file, _ := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	defer file.Close()
	//开启编码工具
	ctx, err := json.Marshal(UsersList)
	if err != nil {
		fmt.Printf("json编码失败%v", err)
		return

	}
	file.Write(ctx)
	DelFile()
}

//定义读取函数
func JsonLoad(name string) {
	ctx, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Printf("文件读取失败%v", err)
		return
	}
	UsersList = []model.User{}
	err = json.Unmarshal(ctx, &UsersList)
	if err != nil {
		fmt.Printf("json解析失败:%v", err)
		return
	}

}
