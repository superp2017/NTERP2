/*
@Time : 2019/9/19 16:48
@Author : mp
@File : timertask
@Software: GoLand
*/
package main

import "time"


////启动定时任务
func startTimerTask()  {
	go timeTask(doTask, 2)
}


func doTask()  {

}

////定时任务
func timeTask(task func(), hour int) {
	for {
		//////////定时任务//////////////////
		now := time.Now()
		// 计算下一个零点
		next := now.Add(time.Hour * 24)
		next = time.Date(next.Year(), next.Month(), next.Day(), hour, 0, 0, 0, next.Location())
		t := time.NewTimer(next.Sub(now))
		<-t.C
		task()
	}
}