package main

import (
	"fmt"
	"strconv"
	"time"
)

func main () {
	t1 := time.Now()
	t2 := time.Date(2022, time.January, 05, 23, 10, 52, 211, time.UTC)

	time := diffForHuman(t1, t2)
	fmt.Println(time)
}

func diffForHuman (t1 time.Time, t2 time.Time) string {
	diff := t1.Sub(t2)
	diffit := int(diff)
	timestr := ""

	if(diffit < 60000000000){
		timestr = "Less than 1 minute ago"
	} else if(diffit < 3600000000000){
		intval := int(diff.Minutes())
		strval := strconv.Itoa(intval)
		timestr = strval + " minute(s) ago"
	} else if(diffit < 86400000000000){
		intval := int(diff.Hours())
		strval := strconv.Itoa(intval)
		timestr = strval + " hour(s) ago"
	} else if(diffit < 2592000000000000){
		intval := int(diff.Hours() / 24)
		strval := strconv.Itoa(intval)
		timestr = strval + " day(s) ago"
	} else if(diffit < 31104000000000000){
		intval := int((diff.Hours() / 30) / 24)
		strval := strconv.Itoa(intval)
		timestr = strval + " month(s) ago"
	} else {
		intval := int(((diff.Hours() / 12) / 30) / 24)
		strval := strconv.Itoa(intval)
		timestr = strval + " year(s) ago"
	}
	return timestr
}