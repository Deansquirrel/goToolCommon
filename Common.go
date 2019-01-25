package goToolCommon

import (
	"crypto/md5"
	"fmt"
	"github.com/satori/go.uuid"
	"math/rand"
	"runtime"
	"strings"
	"sync"
	"time"
)

//生成随机数
func RandInt(min int, max int) int {
	if min == max {
		return min
	}
	if min > max {
		t := min
		min = max
		max = t
	}
	rand.Seed(getRandSeek())
	return min + rand.Intn(max-min)
}

var randSeek = int64(1)
var randMax = int64(1000000)
var l sync.Mutex

//获取随机数种子值
func getRandSeek() int64 {
	l.Lock()
	if randSeek >= randMax {
		randSeek = 0
	}
	randSeek++
	l.Unlock()
	return time.Now().UnixNano() + randSeek
}

//生成GUID
func Guid() string {
	id := uuid.NewV4()
	return strings.ToUpper(id.String())
}

//获取字符串MD5
func Md5(s string) string {
	data := []byte(s)
	has := md5.Sum(data)
	md5Str := fmt.Sprintf("%X", has)
	return md5Str
}

//获取当前操作系统下的换行
func GetWrapStr() string {
	switch runtime.GOOS {
	case "windows":
		return " \r\n"
	case "linux":
		return "\n"
	case "darwin":
		return "\r"
	default:
		return ""
	}
}
