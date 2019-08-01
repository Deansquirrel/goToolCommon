package goToolCommon

import (
	"errors"
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
	t, err := time.Parse(TimeFormat[0:10], s)
	if err != nil {
		return time.Now(), errors.New(fmt.Sprintf("time parse err,require format: %s", TimeFormat[0:10]))
	}
	return t, nil
}

func ParseDateTimeStr(s string) (time.Time, error) {
	t, err := time.Parse(TimeFormat[0:19], s)
	if err != nil {
		return time.Now(), errors.New(fmt.Sprintf("time parse err,require format: %s", TimeFormat[0:19]))
	}
	return t, nil
}

func ParseDateTimeStrWithMillisecond(s string) (time.Time, error) {
	t, err := time.Parse(TimeFormat[0:23], s)
	if err != nil {
		return time.Now(), errors.New(fmt.Sprintf("time parse err,require format: %s", TimeFormat[0:23]))
	}
	return t, nil
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
