package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

type User struct {
	ID   int
	Nane string
}

func main() {
	users := []User{
		//{1, "kk"},
		//{2, "libin"},
	}
	file, err := os.Open("user.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	reader := csv.NewReader(file)
	for {
		line, err := reader.Read()
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		fmt.Println(line)
		id, _ := strconv.Atoi(line[0])
		users = append(users, User{id, line[1]})
	}
	fmt.Println(users)
}
