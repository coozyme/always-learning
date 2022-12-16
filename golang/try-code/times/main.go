package main

import (
	"fmt"
	"log"
	"time"
)

type dateTime struct {
	Date time.Time
}

func NewDate(date string) dateTime {
	t, _ := time.Parse("2006-01-02T15:04:05.999999-07:00", date)
	return dateTime{
		Date: t,
	}
}

func (d *dateTime) Formated() (format string) {
	format = fmt.Sprintf("%02d %s %d", d.Date.Day(), d.Date.Month().String()[:3], d.Date.Year())
	return
}

func (d *dateTime) DateTimes() (year int, month time.Month, day int, hour int, min int, sec int, nsec int, loc *time.Location) {

	year = d.Date.Year()
	month = d.Date.Month()
	day = d.Date.Day()
	hour = d.Date.Hour()
	min = d.Date.Minute()
	sec = d.Date.Second()
	nsec = d.Date.Nanosecond()
	loc = d.Date.UTC().Location()

	return
}

func CompareDate(dateOne, dateTwo string) bool {
	dateTime1 := NewDate(dateOne)
	dateTime2 := NewDate(dateTwo)
	date1 := time.Date(dateTime1.DateTimes())
	date2 := time.Date(dateTime2.DateTimes())
	return date1.After(date2)

	/***
		d1 > d2 = true
		d1 < d2 = false
		d1 == d2 = false

		currentDate := ""
		if true {
			replace current date
		}
	***/
}

func main() {

	d1 := "2022-11-14T23:59:59.000+07:00"
	d2 := "2018-03-19T23:59:59.000+07:00"
	date1 := NewDate(d1)
	date2 := NewDate(d2)

	log.Println(date1.Formated())
	log.Println(date2.Formated())

	log.Println("T", CompareDate(d1, d2))
	log.Println("F", CompareDate(d2, d1))
	log.Println("F", CompareDate(d2, d2))
}
