package utils

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
	"usermanage/model"
)

//定义用户结构体切片类型，以及创建用户的信息
type UserList []model.User

var UsersList UserList

////定义方法满足sort接口
func (this UserList) Len() int {
	return len(this)
}
func (this UserList) Less(i, j int) bool {
	return this[i].Id < this[j].Id
}

func (this UserList) Swap(i, j int) {
	this[i], this[j] = this[i], this[j]
}

//定义用户结构体切片打印
func (this UserList) String() string {
	//对用户列表按id排序
	//ort 操作的对象通常是一个 slice，需要满足三个基本的接口，并且能够使用整数来索引
	sort.Sort(this)
	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)
	table.SetHeader([]string{"Id", "Name", "Addr", "Tel", "Birthday", "Password"})
	for _, v := range this {
		table.Append([]string{strconv.Itoa(v.Id), v.Name, v.Addr, v.Tel, v.Birthday.Format("2006-01-02"), v.Password})
	}
	table.Render()
	return tableString.String()

}
