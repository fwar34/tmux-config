package collect

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func Net(device string) string {
	file, err := os.Open("/proc/net/dev")
	Check(&err)

	reader := bufio.NewReader(file)
	var (
		last_send uint64 = 0
		last_recv uint64 = 0
		now_send  uint64 = 0
		now_recv  uint64 = 0
	)
	for {
		str, err := reader.ReadString('\n')
		Check(&err)
		fields := strings.Fields(str)
		if len(fields) > 0 && fields[0] == device+":" {
			last_recv, err = strconv.ParseUint(fields[1], 10, 64)
			Check(&err)
			last_send, err = strconv.ParseUint(fields[9], 10, 64)
			Check(&err)
			break
		}
	}
	file.Close()

	time.Sleep(time.Second)

	file, err = os.Open("/proc/net/dev")
	Check(&err)
	defer file.Close()
	reader = bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		Check(&err)
		fields := strings.Fields(str)
		if fields[0] == device+":" {
			now_recv, err = strconv.ParseUint(fields[1], 10, 64)
			Check(&err)
			now_send, err = strconv.ParseUint(fields[9], 10, 64)
			Check(&err)
			break
		}
	}

	send := now_send - last_send
	recv := now_recv - last_recv

	var sendStr, recvStr string
	switch {
	case send < 1024:
		sendStr = fmt.Sprintf("UP: %dB", send)
	case send >= 1024 && send < 1024*1024:
		sendStr = fmt.Sprintf("UP: %dKB", send/1024)
	case send >= 1024*1024:
		sendStr = fmt.Sprintf("UP: %dMB", send/1024/1024)
	}
	switch {
	case recv < 1024:
		recvStr = fmt.Sprintf("DOWN: %dB", recv)
	case recv >= 1024 && recv < 1024*1024:
		recvStr = fmt.Sprintf("DOWN: %dKB", recv/1024)
	case recv >= 1024*1024:
		recvStr = fmt.Sprintf("DOWN: %dMB", recv/1024/1024)
	}

	return recvStr + "/" + sendStr
}
