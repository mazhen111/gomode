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

// UserList 定义用户结构体切片类型，以及创建用户的信息
type UserList []model.User

var UsersList UserList

// Len //定义方法满足sort接口
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

//InitAdmin 初始化admin账户
func InitAdmin() {
	//判断admin账户是否存在，存在则先删除再创建
	//删除之前admin账户
	DelUser("admin")
	passwd := SetPasswd("admin")
	adminUser := model.User{
		Id:       0,
		Name:     "admin",
		Password: passwd,
	}
	UsersList = append(UsersList, adminUser)
}
func InitAdminUser() {
	_, exists := isUserExists("admin")
	if !exists {
		InitAdmin()
	} else {
	LABEL:
		for {
			fmt.Printf("已经存在的admin用户,是否重新初始化用户(y/n)?:")
			choise := ""
			fmt.Scan(&choise)
			switch choise {
			case "y", "Y":
				InitAdmin()
				break LABEL
			case "n", "N":
				break LABEL
			default:
				fmt.Println("输入错请重新输入")
			}
		}
	}
}

//定义函数map
var FuncMap = map[string]func(){
	"add":   Add,
	"de;":   Del,
	"upd":   Update,
	"help":  Help,
	"query": Query,
	"list":  List,
	"exit":  Exit,
}

//定义文件持久化接口
type Persist interface {
	Save()
}

func InitPerMod() string {
	fmt.Printf("请输入数据持久化格式(gob/csv/json)")
	var mod string
	fmt.Scan(&mod)
	if mod != "gob" && mod != "csv" && mod != "json" {
		fmt.Println("输入格式化错误，将默认使用json作为持久化格式")
		mod = "jsom"
	}
	file, _ := os.OpenFile("mod.conf", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	defer file.Close()
	file.Write([]byte(mod))
	return mod
}

//定义初始化存储设备
func InitPersist() {
	data, err := ioutil.ReadFile("mod.conf")
	if os.IsNotExist(err) || string(data) == "" {
		fmt.Printf("测到当前系统还未定义存储格式")
		InitPerMod()

	} else if err != nil {
		fmt.Println("获取存储配置失败！退出程序...")
		os.Exit(0)

	} else if string(data) == "csv" || string(data) == "json" || string(data) == "gob" {
	LABEL:
		for {
			fmt.Print("当前已经初始化过持久化配置是否重新初始化(y/n)?:")
			var choise string
			fmt.Scan(&choise)
			switch choise {
			case "y", "Y":
				InitPerMod()
				break LABEL
			case "n", "N":
				break LABEL
			default:
				fmt.Println("输入错误请重新选择")
			}
		}

	} else {
		fmt.Println("配置无法识别，重新初始化")
		InitPerMod()
	}

}
