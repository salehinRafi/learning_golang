package main

import (
	"fmt"
	"net"
)

func GetIpAddress() {
	fmt.Println("------Get IP Address Sections------")
	addr, err := net.LookupIP("www.youtube.com")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Ip of website youtube:", addr)
	}
}
