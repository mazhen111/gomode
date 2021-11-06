package main

import "fmt"

//匿名接口
type EmailSender struct {
}

func (s *EmailSender) Send(msg string) error {
	fmt.Println("msg", msg)
	return nil

}

func main() {
	var sender interface {
		Send(msg string) error
	}
	sender = new(EmailSender)
	sender.Send("xxx")

}
