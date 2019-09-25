/*
@Time : 2019/9/20 10:44
@Author : mp
@File : util
@Software: GoLand
*/
package main

import (
	"fmt"
	"strconv"
	"time"
)

//返回当前时间：例如 2017-02-17 16:33
func CurTime() string {
	return time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
}

//返回当前日期：例如：20170217
func CurDateEx() string {
	return time.Unix(time.Now().Unix(), 0).Format("20060102")
}

//返回当前月份
func CurMonth() string {
	return time.Unix(time.Now().Unix(), 0).Format("200601")
}

//返回当前日期：例如：2017-02-17
func CurDate() string {
	return time.Unix(time.Now().Unix(), 0).Format("2006-01-02")
}

///当前时间的时间戳
func CurStamp() int64 {
	return time.Now().Unix()
}

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}
