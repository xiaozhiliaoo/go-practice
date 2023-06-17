package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	c := cron.New()

	// 添加定时任务
	_, err := c.AddFunc("* * * * *", func() {
		fmt.Println("Task executed at:", time.Now())
	})

	if err != nil {
		fmt.Println("Failed to start cron:", err)
		return
	}

	// 启动定时器
	c.Start()

	// 等待定时器执行
	time.Sleep(5 * time.Minute)

	// 停止定时器
	c.Stop()
}
