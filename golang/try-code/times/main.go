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

func (d *dateTime) FormatDate() string {
	str := fmt.Sprintf("%02d %s %d", d.Date.Day(), d.Date.Month().String()[:3], d.Date.Year())
	return str
}

func (d *dateTime) DateTimes() (year int, month time.Month, day int, hour int, min int, sec int, nsec int, loc *time.Location) {
	// loc, _ = time.LoadLocation("Asia/Jakarta")

	// log.Println("DATE-T", d.Date.Day(), d.Date.Month().String()[:3], d.Date.Year(), d.Date.Hour(), d.Date.Minute(), d.Date.Second(), d.Date.Nanosecond())
	// t, _ := time.ParseInLocation(time.RFC3339, d.Date.String(), loc)
	// log.Println("op", t.Month())

	return d.Date.Year(), d.Date.Month(), d.Date.Day(), d.Date.Hour(), d.Date.Minute(), d.Date.Second(), d.Date.Nanosecond(), d.Date.UTC().Location()

}

func CompareDate(dateOne, dateTwo string) bool {
	dateTime1 := NewDate(dateOne)
	dateTime2 := NewDate(dateTwo)
	// year1, month1, day1, hour1, min1, sec1, nsec1, loc1 := dateTime1.DateTimes()
	// year2, month2, day2, hour2, min2, sec2, nsec2, loc2 := dateTime2.DateTimes()
	// log.Println(dateTime1.DateTimes())
	date1 := time.Date(dateTime1.DateTimes())
	date2 := time.Date(dateTime2.DateTimes())
	// log.Println("LOG-D-TIME", year1, month1, day1, hour1, min1, sec1, nsec1, loc1)
	return date1.After(date2)
	// d1 > d2 = true
	// d1 < d2 = false
	// d1 == d2 = false

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

	// date := "2018-03-19T23:59:59.000+07:00"
	// tdm, _ := time.Parse(layoutISO, date)
	// now := time.Now()
	// fmt.Println(tdm.Format(tdm.Day()))
	// loc, _ := time.LoadLocation("Asia/Jakarta")
	// t, _ := time.ParseInLocation(time.RFC3339, date, loc)
	// sd := fmt.Sprintf("%02d-%02d-%02d", tdm.Day(), tdm.Month(), tdm.Year())
	// formatted := fmt.Sprintf("%Y%m%d", t.Day(), t.Month(), t.Year())
	// str := fmt.Sprintf("%02d %s %d", t.Day(), t.Month().String()[:3], t.Year())
	// log.Println(t)
	// log.Println(str)
	// log.Println(sd)

	d1 := "2022-11-14T23:59:59.000+07:00"
	d2 := "2018-03-19T23:59:59.000+07:00"
	date1 := NewDate(d1)
	// date2 := NewDate(d2)
	// log.Println("LOP-", date1.Date)
	// log.Println(date1.DateTimes())
	log.Println("dd-Formatted", date1.FormatDate())
	// log.Println(date2.DateTimes())

	log.Println("T", CompareDate(d1, d2))
	log.Println("F", CompareDate(d2, d1))
	log.Println("F", CompareDate(d2, d2))
}
