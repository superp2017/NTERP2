/*
@Time : 2019/9/18 17:53
@Author : mp
@File : main.go
@Software: GoLand
*/
package main

import (
	"github.com/superp2017/golibs/Daemon"
	"runtime"
)


func init()  {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main()  {
	Daemon.StartDaemon()
	startDB()
	startTcp()
	startHttp()
	//startTimerTask()
}