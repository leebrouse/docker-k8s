package models

import (
	"time"
)

func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

// 日期转换为时间戳
func DateToUnix(timeStr string) int64 {
	template := "2006-01-02 15:04:05"
	t, err := time.ParseInLocation(template, timeStr, time.Local)
	if err != nil {
		return 0
	}

	return t.Unix()
}

// 获取当前时间戳
func GetUnix() int64 {
	return time.Now().Unix()
}

// 获取当前时间
func GetDate() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 获取当前年月日
func GetDay() string {
	template := "20060102"
	return time.Now().Format(template)
}
