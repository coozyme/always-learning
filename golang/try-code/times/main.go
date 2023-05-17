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
	t, _ := time.Parse(time.RFC3339, date)
	if date == "" {
		t = time.Now()
	}
	return dateTime{
		Date: t,
	}
}

// Format Date: DD/MM/YY
func (d dateTime) FormatedDateSlash() (format string) {
	lasttwoDigits := d.Date.Year() % 1e2
	format = fmt.Sprintf("%02d/%02d/%02d", d.Date.Day(), d.Date.Month(), lasttwoDigits)
	return
}

func main() {
	d2 := "2023-01-06T03:37:47Z"
	log.Println("dates", NewDate(d2).FormatedDateSlash())
	// tm := time.Now().Format()
	// tmi := time.Parse("2006-01-02 15:04:05", )
	tmi := time.Now().Add(time.Minute * 10).Format("2006-01-02 15:04:05")
	// log.Println("LOG-TM", tm.Format("2006-01-02 15:04:05"))
	// lf := time.Minute * 10
	log.Printf("log-ro %v", tmi)
	// log.Println("log-r", lf.Minutes())
	log.Println("LOGSDS", time.Now().Format(time.RFC3339))

}
