package collect

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func check(err *error) {
	if *err != nil {
		panic(*err)
	}
}

func Cpu() string {
	file, err := os.Open("/proc/stat")
	check(&err)

	reader := bufio.NewReader(file)
	str, err := reader.ReadString('\n')
	check(&err)
	file.Close()

	var total_time1 float64
	var idle_time1 float64
	for k, v := range(strings.Fields(str)) {
		if k == 0 {
			continue
		}

		if k == 4 {
			idle_time1, err = strconv.ParseFloat(v, 64)
			check(&err)
		}
		vuint, err := strconv.ParseFloat(v, 64)
		check(&err)
		total_time1 += vuint
	}

	time.Sleep(time.Millisecond * 990)

	file, err = os.Open("/proc/stat")
	check(&err)

	reader = bufio.NewReader(file)
	str, err = reader.ReadString('\n')
	check(&err)
	file.Close()

	var total_time2 float64
	var idle_time2 float64
	for k, v := range(strings.Fields(str)) {
		if k == 0 {
			continue
		}

		if k == 4 {
			idle_time2, err = strconv.ParseFloat(v, 64)
			check(&err)
		}
		vuint, err := strconv.ParseFloat(v, 64)
		check(&err)
		total_time2 += vuint
	}

	return fmt.Sprintf("%.1f%%", (((total_time2 - total_time1) - (idle_time2 - idle_time1)) / (total_time2 - total_time1)) * 100)
}
