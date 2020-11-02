package timer

import "time"

// 定时器
func startTimer(f func(), hours int, minutes int, seconds int) {
	go func() {
		for {
			f()
			now := time.Now()
			next := now
			if hours > 0 {
				next = next.Add(time.Hour * time.Duration(hours))
			}
			if minutes > 0 {
				next = next.Add(time.Minute * time.Duration(minutes))
			}
			if seconds > 0 {
				next = next.Add(time.Second * time.Duration(seconds))
			}

			next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), next.Minute(), next.Second(), 0, next.Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
		}
	}()
}

