package collect

import (
	"bufio"
	"fmt"
	"os"
	// "strconv"
	"strings"
)

const (
	BYTES     = 0
	KILOBYTES = 1
	MEGABYTES = 2
	GIGABYTES = 3
)

func value(str *string) uint64 {
	fmt.Println(*str)
	// num, err := strconv.ParseUint(strings.ReplaceAll(strings.ReplaceAll(*str, " ", ""), "kB\n", ""), 10, 64)
// println(strings.ReplaceAll(strings.ReplaceAll(*str, " ", ""), "kB\n", ""))
	// if err != nil {
	// 	panic(err)
	// }
	// return num
	return 0
}

func convert(num uint64, from int32, to int32) uint64 {
	for ; from < to; from++ {
		num /= 1024
	}
	return num
}

func Memory() (uint64, error) {
	var total_mem uint64
	var used_mem uint64
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return 0, err
	}
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
		case "FreeMem":
			used_mem = total_mem - value(&strs[1])
		case "Shmem":
			used_mem += value(&strs[1])
		case "Buffers", "Cached", "SReclaimable":
			used_mem -= value(&strs[1])
		}
	}

	return convert(used_mem, KILOBYTES, MEGABYTES), nil
}
