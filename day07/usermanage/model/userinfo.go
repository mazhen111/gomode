package model

import (
	"github.com/olekukonko/tablewriter"
	"strconv"
	"strings"
	"time"
)

type User struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Addr     string    `json:"addr"`
	Tel      string    `json:"tel"`
	Birthday time.Time `json:"birthday"`
	Password string    `json:"password"`
}

//以表格的形式打印单条用户信息
func (this User) String() string {
	tebleString := &strings.Builder{} //追加字串用bytes.Buffer
	table := tablewriter.NewWriter(tebleString)
	table.Append([]string{strconv.Itoa(this.Id), this.Name, this.Addr, this.Tel,
		this.Birthday.Format("2006-01-02"), this.Password})
	table.Render()
	return tebleString.String()
}
