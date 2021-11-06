package main

import "fmt"

/*
// 产生告警 => 发送给相关的人
// email, sms, weixin, dingding
// 发送的信息 => 发送的方式
*/

type Sender interface {
	Send(string) error
}

type AllSender interface {
	Send(string) error
	SendAll([]string) error
}

type EmailSender struct {
}

type DingDingSender struct {
}

func (f *EmailSender) Send(msg string) error {

	fmt.Println("email sender", msg)
	return nil
}
func (f *DingDingSender) Send(msg string) error {
	fmt.Println("dingding sendr", msg)
	return nil
}
func (f *EmailSender) SendAll(msg []string) error {

	fmt.Println("dingding sender all:", msg)
	return nil
}

func main() {
	fmt.Println("11")
	//var sender *EmailSender=new(EmailSender)
	//sender.Send("XXXX")
	//var sender *DingDingSender=new(DingDingSender)
	//sender.Send("didi")
	var sender Sender = new(EmailSender)
	sender.Send("111")
	var allsender AllSender = new(EmailSender)
	err := allsender.SendAll([]string{"aaaaa"})
	if err != nil {
		return
	}
	allsender.Send("mazhen")
	sender = allsender
	sender.Send("xxxx")
	fmt.Println(allsender, sender)

}
