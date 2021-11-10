package utils

import (
	"fmt"
	"os"
	"sort"
	"strings"
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
func GetSortedFileList() MyFileInfoList {
	var fileInfos MyFileInfoList
	path, _ := os.Open("data")
	files, _ := path.Readdir(-1)
	if len(files) == 0 {
		return []os.FileInfo{}

	}
	for _, file := range files {
		fileInfos = append(fileInfos, file)
	}
	sort.Sort(fileInfos)
	return fileInfos

}

//永久保存文件
func Load() {
	//创建data文件夹
	HasDataDIr()

	fileInfos := GetSortedFileList()
	file := fileInfos[len(fileInfos)-1]
	fileName := fmt.Sprintf("data/%v", file.Name())
	//判断s是否有后缀字符串suffix。
	if strings.HasSuffix(file.Name(), ".csv") {
		CsvLoad(fileName)

	} else if strings.HasSuffix(file.Name(), ".gob") {
		GobLoad(fileName)

	} else if strings.HasSuffix(file.Name(), ".json") {
		JsonLoad(fileName)

	}
}
func GetPersist() Persist {

	return Persist

}
