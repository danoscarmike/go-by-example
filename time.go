package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now)
	p(now.Location().String())

	then := time.Date(2016, time.November, 18, 12, 0, 0, 0, time.UTC)
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Weekday())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then)
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Microseconds())
	p(diff.Nanoseconds())

	isItNow := then.Add(diff)
	isItThen := now.Add(-diff)
	p(isItNow.Equal(now))
	p(isItThen.Equal(then))
}
