package cron

import (
	"fmt"
	"time"
)

func taskSecond () {
	fmt.Println("@@@...second function, ", time.Now())

}

func taskMinute() {
	fmt.Println("###...minute function, ", time.Now())

}

func taskHour() {
	fmt.Println("$$$...hour function, ", time.Now())

}

func TaskSecond() {
	Cron(taskSecond, -1, -1, 10)
}

func TaskMinute() {
	Cron(taskMinute, -1, 0, 0)
}

func TaskHour() {
	Cron(taskHour, 1, 0, 0)
}

