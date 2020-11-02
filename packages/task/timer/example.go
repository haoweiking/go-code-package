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
	Timer(taskSecond, 0, 0, 10)
}

func TaskMinute() {
	Timer(taskMinute, 0, 10, 0)
}

func TaskHour() {
	Timer(taskHour, 1, 0, 0)
}

