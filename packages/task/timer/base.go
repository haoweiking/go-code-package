package timer

import "time"

// Timer function is a method that it controls the hour, minute and second of the function run.
// The f is a function param, and others are time duration param
func Timer(f func(), hours int, minutes int, seconds int) {
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
