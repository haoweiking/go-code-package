package main

import (
	"fmt"
	"github.com/go-code-package/packages/task/cron"
	"time"
)

func main() {
	go cron.TaskHour()
	fmt.Println("1111")
	go cron.TaskMinute()
	fmt.Println("2222")
	go cron.TaskSecond()
	fmt.Println("3333")
	time.Sleep(70 * time.Minute)
}
