package main

import (
	"fmt"
	"time"
)

func main() {
	endDateStr  := "2020-05-31"
	startDateStr := "2019-11-16"

	promoStartDate, _ := convertDateStringToTime(startDateStr)
	promoEndDate, _ := convertDateStringToTime(endDateStr)

	fmt.Println("Promo Start date: ", promoStartDate)
	fmt.Println("Promo End date: ", promoEndDate)

	newDate := promoStartDate.Add(-24 * time.Hour)
	fmt.Println("promo start date minus 1 day: ", newDate)
	
}

func convertDateStringToTime(str string) (*time.Time, error) {
	layout := "2006-01-02"
	t, err := time.Parse(layout, str)

	if err != nil {
		return nil, err
	}
	return &t, nil
}