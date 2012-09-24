package main

import (
	"fmt"
	"time"
)

var week time.Duration

func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())
	t = time.Now().UTC()
	fmt.Println(t)
	// calculating times:
	week = 60 * 60 * 24 * 7 * 1e9		// must be in nanosec
	week_from_now := t.Add(week)
	fmt.Println(week_from_now)
	t = time.Now()
	// formatting times:
	// These are predefined layouts for use in Time.Format.
  // The standard time used in the layouts is:
  //      Mon Jan 2 15:04:05 MST 2006
	fmt.Println(t.Format(time.RFC822))
	fmt.Println(t.Format(time.ANSIC))
	s := t.Format("20060102")
	fmt.Println(t, "==>", s)
	fmt.Println(t.Format("02 Jan 2006 15:04"))
}
