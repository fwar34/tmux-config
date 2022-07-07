package main

import (
	"fmt"
	// "flag"
	"collect-go/collect"
)

func main() {
	// fmt.Println(collect.Memory(), collect.Cpu(), collect.Net(*flag.String("device", "eth0", "network interface name")))
	fmt.Println(collect.Memory(), collect.Cpu())
}
