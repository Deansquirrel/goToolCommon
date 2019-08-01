package goToolCommon

import (
	"fmt"
	"time"
)

const (
	TimeFormat = "2006-01-02 15:04:05.000"
)

func GetDateStr(t time.Time) string {
	return t.Format(TimeFormat[0:10])
}

func GetDateTimeStr(t time.Time) string {
	return t.Format(TimeFormat[0:19])
}

func GetDateTimeStrWithMillisecond(t time.Time) string {
	return t.Format(TimeFormat[0:23])
}

func ParseDateStr(s string) (time.Time, error) {
	return time.Parse(TimeFormat[0:10], fmt.Sprintf("parse err,reqire format %s", TimeFormat[0:10]))
}

func ParseDateTimeStr(s string) (time.Time, error) {
	return time.Parse(TimeFormat[0:19], fmt.Sprintf("parse err,reqire format %s", TimeFormat[0:19]))
}

func ParseDateTimeStrWithMillisecond(s string) (time.Time, error) {
	return time.Parse(TimeFormat[0:23], fmt.Sprintf("parse err,reqire format %s", TimeFormat[0:23]))
}

func GetMillisecond(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

func GetMicrosecond(t time.Time) int64 {
	return t.UnixNano() / 1e3
}

func GetDurationBySecond(l int) time.Duration {
	return time.Duration(int64(1000 * 1000 * 1000 * int64(l)))
}

func GetDurationByMinute(l int) time.Duration {
	return time.Duration(1000 * 1000 * 1000 * 60 * int64(l))
}

func GetDurationByHour(l int) time.Duration {
	return time.Duration(int64(1000 * 1000 * 1000 * 60 * 60 * int64(l)))
}

func GetDurationByDay(l int) time.Duration {
	return time.Duration(int64(1000 * 1000 * 1000 * 60 * 60 * 24 * int64(l)))
}
