package main

import (
	"github.com/go-code-package/packages/task/timer"
)

func main() {
	go timer.TaskHour()
	go timer.TaskMinute()
	go timer.TaskSecond()
}
