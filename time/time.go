package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	utc := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println(utc)

	layout := "2006-01-02"
	parse, _ := time.Parse(layout, "1994-10-11")
	fmt.Println(parse)
}

func timezone() {

	nowere := time.Now()
	fmt.Println(nowere.Format(time.RFC3339))
	fmt.Println(time.Now().Format(time.RFC3339))

	fmt.Println(time.Now())
}
