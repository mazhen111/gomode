package main

import (
	"fmt"
	"net"
)

func main() {
	//ipv4 ipv6 主机名/域名
	fmt.Println(net.JoinHostPort("127.0.0.1", "8080")) //127.0.0.1:8080
	fmt.Println(net.SplitHostPort("127.0.0.1:8080"))   //127.0.0.1 8080 <nil>
	//ip->主机名
	fmt.Println(net.LookupHost("localhost")) //[::1 127.0.0.1] <nil>
	// ip ipnet
	for _, ipstr := range []string{"127.0.0.1", ":::1"} {
		ip := net.ParseIP(ipstr)
		fmt.Printf("%v\n", ip)
	}

	ip, ipnet, err := net.ParseCIDR("192.168.0.0/24")
	fmt.Printf("%#v %#v %#v\n", ip, ipnet, err)
	//判断ip是否在范围之内
	fmt.Println(ipnet.Contains(net.ParseIP("192.168.0.10"))) //true
	//获取本地ip
	addrs, err := net.InterfaceAddrs()
	fmt.Println(addrs, err)
	//获取网卡信息
	intfs, err := net.Interfaces()
	//获取网卡详细参数
	fmt.Println(intfs, err)
	for _, intf := range intfs {
		fmt.Println(intf.Index, intf.MTU, intf.Name, intf.HardwareAddr, intf.Flags)
		fmt.Println("--------------------------")
		fmt.Println(intf.Addrs())

	}

}
