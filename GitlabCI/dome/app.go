package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main()  {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		for s := range c {
			switch s {
			case syscall.SIGINT, syscall.SIGTERM:
				fmt.Println("退出", s)
				ExitFunc()
			default:
				fmt.Println("other", s)
			}
		}
	}()

	fmt.Println("进程启动...")
	time.Sleep(time.Duration(200000)*time.Second)
}

func ExitFunc()  {
	fmt.Println("正在退出...")
	fmt.Println("执行清理...")
	fmt.Println("退出完成...")
	os.Exit(0)
}
