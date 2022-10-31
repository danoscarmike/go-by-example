package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	t := time.Now()
	p(t.Format(time.RFC3339))

	t1, err := time.Parse(time.RFC3339, "2022-10-24T15:52:52-07:00")
	if err != nil {
		panic(err)
	}
	p(t1)

	p(t.Format("3:04PM"))
	p(t.Format(time.Kitchen))
	p(t.Format(time.StampMicro))
	p(t.Format("2006-01-02 03:04:05PM"))

	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 AM")
	if e != nil {
		panic(e)
	}
	p(t2)

}
