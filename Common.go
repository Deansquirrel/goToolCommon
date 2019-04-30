package goToolCommon

import (
	"crypto/md5"
	"fmt"
	"github.com/satori/go.uuid"
	"golang.org/x/text/encoding/simplifiedchinese"
	"math/rand"
	"runtime"
	"strconv"
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
func Md5(data []byte) string {
	has := md5.Sum(data)
	md5Str := fmt.Sprintf("%X", has)
	return md5Str
}

//获取当前操作系统下的换行
func GetWrapStr() string {
	switch runtime.GOOS {
	case "windows":
		return "\r\n"
	case "linux":
		return "\n"
	case "darwin":
		return "\r"
	default:
		return "\r\n"
	}
}

//获取当前操作系统的路径分割符
func GetFolderSplitStr() string {
	switch runtime.GOOS {
	case "windows":
		return "\\"
	case "linux":
		return "/"
	case "darwin":
		return "\\"
	default:
		return ""
	}
}

//检查并删除字符串的第一个字符
func CheckAndDeleteFirstChar(s string, deleteChar string) string {
	if strings.HasPrefix(s, deleteChar) {
		s = s[len(deleteChar):len(s)]
		return CheckAndDeleteFirstChar(s, deleteChar)
	} else {
		return s
	}
}

//检查并删除字符串的最后一个字符
func CheckAndDeleteLastChar(s string, deleteChar string) string {
	if strings.HasSuffix(s, deleteChar) {
		s = s[:len(s)-len(deleteChar)]
		return CheckAndDeleteLastChar(s, deleteChar)
	} else {
		return s
	}
}

//清除字符串数组中的空白项
func ClearBlock(list []string) []string {
	r := make([]string, 0)
	for _, s := range list {
		if strings.Trim(s, " ") != "" {
			r = append(r, s)
		}
	}
	return r
}

//清除字符串数组中的重复项
func ClearRepeat(list []string) []string {
	r := make([]string, 0)
	for _, s := range list {
		if IsExist(r, s) {
			continue
		} else {
			r = append(r, s)
		}
	}
	return r
}

//检查字符串列表中是否包含特定的项
func IsExist(list []string, item string) bool {
	for _, s := range list {
		if s == item {
			return true
		}
	}
	return false
}

//比较两个字符串数组
func CheckDiff(listA []string, listB []string) (onlyA []string, onlyB []string, existAB []string) {
	for _, s := range listA {
		if IsExist(listB, s) {
			existAB = append(existAB, s)
		} else {
			onlyA = append(onlyA, s)
		}
	}
	for _, s := range listB {
		if !IsExist(listA, s) {
			onlyB = append(onlyB, s)
		}
	}
	return
}

//转换为字符串
func ConvertToString(arg interface{}) string {
	switch f := arg.(type) {
	case time.Time:
		return GetDateTimeStr(f)
	case bool:
		return strconv.FormatBool(f)
	case float32:
		return strconv.FormatFloat(float64(f), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(f, 'f', -1, 64)
	case int:
		return strconv.FormatInt(int64(f), 10)
	case int8:
		return strconv.FormatInt(int64(f), 10)
	case int16:
		return strconv.FormatInt(int64(f), 10)
	case int32:
		return strconv.FormatInt(int64(f), 10)
	case int64:
		return strconv.FormatInt(int64(f), 10)
	case uint:
		return strconv.FormatUint(uint64(f), 10)
	case uint8:
		return strconv.FormatUint(uint64(f), 10)
	case uint16:
		return strconv.FormatUint(uint64(f), 10)
	case uint32:
		return strconv.FormatUint(uint64(f), 10)
	case uint64:
		return strconv.FormatUint(uint64(f), 10)
	case uintptr:
		return strconv.FormatUint(uint64(f), 10)
	case string:
		r, err := simplifiedchinese.GB18030.NewDecoder().String(f)
		if err != nil {
			return f
		} else {
			return r
		}
	case []byte:
		r, err := simplifiedchinese.GB18030.NewDecoder().Bytes(f)
		if err != nil {
			return string(f)
		} else {
			return string(r)
		}
	default:
		return fmt.Sprint(arg)
	}
}
