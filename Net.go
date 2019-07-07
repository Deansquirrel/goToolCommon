package goToolCommon

import (
	"errors"
	"os/exec"
	"runtime"
)

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
	cmd := exec.Command("ping", "www.BaiDu.com", "-w", "30", "-n", "1")
	err := cmd.Run()
	if err != nil {
		return false
	} else {
		return true
	}
}

func netWorkStatusLinux() bool {
	cmd := exec.Command("ping", "www.BaiDu.com", "-W", "30", "-c", "1")
	err := cmd.Run()
	if err != nil {
		return false
	} else {
		return true
	}
}
