/*
@Time : 2019/9/18 17:53
@Author : mp
@File : main.go
@Software: GoLand
*/
package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"syscall"
)

func daemon() {
	if syscall.Getppid() == 1 { // already daemon
		f, err := os.OpenFile("/dev/null", os.O_RDWR, 0)
		if err != nil {
			fmt.Println("open /dev/null failed")
			os.Exit(-1)
		}

		fd := f.Fd()
		syscall.Dup2(int(fd), int(os.Stdin.Fd()))
		syscall.Dup2(int(fd), int(os.Stdout.Fd()))
		syscall.Dup2(int(fd), int(os.Stderr.Fd()))

		return
	}

	args := append([]string{os.Args[0]}, os.Args[1:]...)
	_, err := os.StartProcess(os.Args[0], args,
		&os.ProcAttr{Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}})

	if err != nil {
		log.Println("Daemon start failed:", err)
	}

	os.Exit(0)
}

func init()  {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main()  {
	daemon()
}