package main

import (
	"fmt"
	"time"
)

func main() {
	startDate := "2015-10-01"
	endDate := "2016-03-01"

	layout := "2006-01-02"

	var months = []string{
		"jan",
		"feb",
		"mar",
		"apr",
		"may",
		"jun",
		"jul",
		"aug",
		"sep",
		"oct",
		"nov",
		"dec",
	}

	startTime, err := time.Parse(layout, startDate)
	if err != nil {
		panic(err)
	}
	endTime, err := time.Parse(layout, endDate)
	if err != nil {
		panic(err)
	}

	sYear, sMonth, sDate := startTime.Date()
	eYear, eMonth, eDate := endTime.Date()
	if sYear > eYear {
		panic(fmt.Errorf("start date is greater than end date"))
	}

	if sYear == eYear && sMonth > eMonth {
		panic(fmt.Errorf("start date is greater than end date"))
	}

	if sYear == eYear && sMonth == eMonth && sDate >= eDate {
		panic(fmt.Errorf("start date is greater than end date"))
	}

	fmt.Println(startTime, endTime, months)
}
