package goToolCommon

import "time"

func GetDateStr(t time.Time) string {
	return t.Format("2006-01-02")
}

func GetDateTimeStr(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func GetMillisecond(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

func GetMicrosecond(t time.Time) int64 {
	return t.UnixNano() / 1e3
}
