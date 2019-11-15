package main

import (
	"log"
	"time"
)



func main() {
	t := int64(1546926630)
	t1 := "2019-01-08 13:50:30"

	timeTemplate1 := "2006-01-02 15:04:05"
	timeTemplate2 := "2006/01/02 15:04:05"
	timeTemplate3 := "2006-01-02"
	timeTemplate4 := "15:04:05"


	log.Println(time.Unix(t, 0).Format(timeTemplate1))
	log.Println(time.Unix(t, 0).Format(timeTemplate2))
	log.Println(time.Unix(t, 0).Format(timeTemplate3))
	log.Println(time.Unix(t, 0).Format(timeTemplate4))


	stamp, _ := time.ParseInLocation(timeTemplate1, t1, time.Local)
	log.Println(stamp.Unix())
}

