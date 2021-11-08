package utils

import (
	"os"
	"sort"
)

//FileInfo用来描述一个文件对象。
type MyFileInfoList []os.FileInfo

func (this MyFileInfoList) Len() int {
	return len(this)
}

func (this MyFileInfoList) Less(i, j int) bool {
	return this[j].ModTime().Sub(this[i].ModTime()) > 0
}

func (this MyFileInfoList) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
func GetMaxId() int {
	if len(UsersList) > 0 {
		sort.Sort(UsersList)
		return UsersList[len(UsersList)-1].Id
	}
	return 0
}
