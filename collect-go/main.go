package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	BYTES = 0
	KILOBYTES = 1
	MEGABYTES = 2
	GIGABYTES = 3
)

func convert(num uint64, from int32, to int32) uint64 {
	for ; to > from; from ++ {
		num /= 1024
	}
	return num
}

func check(err *error) {
	if *err != nil {
		panic(*err)
	}
}

func value(str *string) uint64 {
	v, err := strconv.ParseUint(strings.ReplaceAll(strings.ReplaceAll(*str, " ", ""), "kB\n", ""), 10, 64)
	check(&err)
	return v
}

func main() {
	var total_mem uint64
	var used_mem uint64
	f, err := os.Open("/proc/meminfo")
	check(&err)
	defer f.Close()

	reader := bufio.NewReader(f)
	var str string
	for {
		str, err = reader.ReadString('\n')
		if err != nil {
			break
		}
		// fmt.Println(str)
		// fmt.Println(strings.Index(str, ":"))
		strs := strings.Split(str, ":")
		switch strs[0] {
		case "MemTotal":
			total_mem = value(&strs[1])
		case "MemFree":
			used_mem = total_mem - value(&strs[1])
		case "Shmem":
			used_mem += value(&strs[1])
		case "Buffers", "Cached", "SReclaimable":
			used_mem -= value(&strs[1])
		}
	}
	fmt.Println(convert(used_mem, KILOBYTES, MEGABYTES))
}
