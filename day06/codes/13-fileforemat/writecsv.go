package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

//写入csv
type User struct {
	ID   int
	Name string
}

func main() {

	users := []User{
		{1, "kk"},
		{2, "mazhen"},
	}
	file, err := os.Create("user.csv")
	if err != nil {
		return

	}
	defer file.Close()
	writer := csv.NewWriter(file)
	for _, user := range users {
		writer.Write([]string{strconv.Itoa(user.ID), user.Name})
	}
	writer.Flush()

}
