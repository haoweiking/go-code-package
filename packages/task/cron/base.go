package cron

import "time"

// func runs at specified time
func Cron(f func(), hour int, minute int, second int) {
	go func() {
		for {
			f()
			now := time.Now()

			next := now

			nextHour := hour
			if hour == -1 {
				if minute == -1 {
					next = next.Add(time.Minute * 1)
				} else {
					next = next.Add(time.Hour * 1)
				}
				nextHour = next.Hour()

			} else {
				next = next.Add(time.Hour * 24)
			}

			next = time.Date(next.Year(), next.Month(), next.Day(), nextHour, next.Minute(), second, 0, next.Location())

			t := time.NewTimer(next.Sub(now))
			<-t.C
		}
	}()
}
