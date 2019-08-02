package goToolCommon

import (
	"math/rand"
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
