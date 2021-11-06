package main

import (
	"fmt"
)

type Sender interface {
	Send(string) error
}

type AllSender interface {
	Send(string) error
	SendAll([]string) error
}

type EmailSender struct {
	Addr string
	To   string
}

type DingDingSender struct {
	To string
}

func (f *EmailSender) Send(msgs string) error {
	fmt.Println("email sender:", msgs)
	return nil
}
func (f *EmailSender) SendAll(msge []string) error {
	fmt.Println("email sender:", msge)
	return nil
}

func (f *DingDingSender) Send(msg string) error {
	fmt.Println("dingding sender:" + msg)
	return nil
}

func (s *EmailSender) Test() {
	fmt.Println("test")
}
func test(sender Sender) {
	if obj, ok := sender.(*EmailSender); ok {
		fmt.Println("email sender", obj.Addr)
	} else if obj, ok := sender.(*DingDingSender); ok {
		fmt.Println("dingding sender", obj.To)
	} else {

		fmt.Println("error")
	}
}

func query(sender Sender) {
	//类型查询
	// switch case + 接口变量.(type)
	switch obj := sender.(type) {
	case *EmailSender:
		fmt.Println("email sender", obj.Addr, obj.To)
	case *DingDingSender:
		fmt.Println("dingding sender", obj.To)
	default:
		fmt.Println("error")
	}

}

func main() {
	var sender Sender = &EmailSender{"xxxx", "yyyyy"}
	sender.Send("xxxx")
	fmt.Println("11")
	//断言
	obj, ok := sender.(*EmailSender)
	fmt.Printf("%T, %#v, %#v\n", obj, obj, ok)
	if ok {
		fmt.Println(obj.To)
		fmt.Println(obj.Addr)
		obj.Test()
	}
	dingding, ok := sender.(*DingDingSender)
	fmt.Println(dingding)
	test(sender)              //email sender xxxx
	test(new(DingDingSender)) //dingding sender
	query(sender)
	query(new(DingDingSender))
	var allsender AllSender = &EmailSender{"111", "222"}
	//allsender.SendAll([]string{"xxxxx"})
	//allsender.Send("xxxx")
	sender = allsender
	//allsender=sender
	sender.Send("xxxx")
	esender, ok := sender.(*EmailSender)
	fmt.Println(esender, ok)
	fmt.Println(esender.Addr, esender.To)
	esender.Test()
	switch sender.(type) {
	case *EmailSender:
		fmt.Println("emailstnder")
	case AllSender:
		fmt.Println("allsender")
	default:
		fmt.Println("default")
	}
}
