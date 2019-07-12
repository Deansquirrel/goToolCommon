package goToolCommon

import "time"

func GetDateStr(t time.Time) string {
	return t.Format("2006-01-02")
}

func GetDateTimeStr(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func GetDateTimeStrWithMillisecond(t time.Time) string {
	return t.Format("2006-01-02 15:04:05.000")
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
