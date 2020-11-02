package timer

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
	startTimer(taskSecond, 0, 0, 10)
}

func TaskMinute() {
	startTimer(taskMinute, 0, 10, 0)
}

func TaskHour() {
	startTimer(taskHour, 1, 0, 0)
}

