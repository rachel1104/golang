package models

import (
	"fmt"
	"time"
)

// 可在所有控制器下使用Models封装公共方法
// 时间戳转换成日期
func UnixToTime(timestamp int) string {
	fmt.Println(timestamp)
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

// 日期转换为时间戳 2024-12-12 12:02:02
func DataToUnix(str string) int64 {
	template := "2024-12-12 12:02:02"
	t, err := time.ParseInLocation(template, str, time.Local)
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
	template := "2006-01-02 15:04:05"
	return time.Now().Format(template)
}

// 获取年月日
func GetDay() string {
	template := "2006-01-02"
	return time.Now().Format(template)
}
