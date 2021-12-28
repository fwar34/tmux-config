package collect

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	BYTES     = 0
	KILOBYTES = 1
	MEGABYTES = 2
	GIGABYTES = 3
)

func value(str *string) uint64 {
	num, err := strconv.ParseUint(strings.ReplaceAll(strings.ReplaceAll(*str, " ", ""), "kB\n", ""), 10, 64)
	if err != nil {
		panic(err)
	}
	return num
}

func convert(num uint64, from int32, to int32) uint64 {
	for ; from < to; from++ {
		num /= 1024
	}
	return num
}

func Memory() string {
	var total_mem uint64
	var used_mem uint64
	file, err := os.Open("/proc/meminfo")
	Check(&err)
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			break
		}
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

	return fmt.Sprintf("%dMB/%dMB", convert(used_mem, KILOBYTES, MEGABYTES), convert(total_mem, KILOBYTES, MEGABYTES))
}
