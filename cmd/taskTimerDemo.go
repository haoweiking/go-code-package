package main

import (
	"fmt"
	"github.com/go-code-package/packages/task/timer"
	"time"
)

func main() {
	go timer.TaskHour()
	fmt.Println("1111")
	go timer.TaskMinute()
	fmt.Println("2222")
	go timer.TaskSecond()
	fmt.Println("3333")
	time.Sleep(70 * time.Minute)
}
