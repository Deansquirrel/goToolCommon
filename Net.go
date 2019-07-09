package goToolCommon

import (
	"errors"
	"os/exec"
	"runtime"
	"strconv"
)

var NetWorkTestAddress string
var NetWorkTestTimeout int

func init() {
	NetWorkTestAddress = "www.baidu.com"
	NetWorkTestTimeout = 30
}

func NetWorkStatus() bool {
	switch runtime.GOOS {
	case "windows":
		return netWorkStatusWindows()
	case "linux":
		return netWorkStatusLinux()
	case "darwin":
		panic(errors.New("does not support IOS"))
	default:
		panic(errors.New("unknown operating platform"))
	}
}

func netWorkStatusWindows() bool {
	cmd := exec.Command("ping", NetWorkTestAddress, "-w", strconv.Itoa(NetWorkTestTimeout), "-n", "1")
	err := cmd.Run()
	if err != nil {
		return false
	} else {
		return true
	}
}

func netWorkStatusLinux() bool {
	cmd := exec.Command("ping", NetWorkTestAddress, "-W", strconv.Itoa(NetWorkTestTimeout), "-c", "1")
	err := cmd.Run()
	if err != nil {
		return false
	} else {
		return true
	}
}
