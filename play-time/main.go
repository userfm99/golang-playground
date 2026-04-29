package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func main() {
	tsStart := "05:55"
	tsEnd := "20:02"

	shInt, smInt, err := splitTimeSlot(tsStart)
	if err != nil {
		log.Fatal(err)
	}
	ehInt, emInt, err := splitTimeSlot(tsEnd)
	if err != nil {
		log.Fatal(err)
	}

	t := time.Now()
	startSlot := time.Date(t.Year(), t.Month(), t.Day(), shInt, smInt, 0, 0, t.Location())
	endSlot := time.Date(t.Year(), t.Month(), t.Day(), ehInt, emInt, 0, 0, t.Location())

	fmt.Println("time slot start", startSlot)
	fmt.Println("time slot end", endSlot)
}

func splitTimeSlot(slot string) (int, int, error) {
	hm := strings.Split(slot, ":")
	if len(hm) != 2 {
		return 0, 0, errors.New("time slot invalid")
	}
	hInt, err := strconv.Atoi(hm[0])
	if err != nil {
		return 0, 0, errors.New("error parsing hour. " + err.Error())
	}
	mInt, err := strconv.Atoi(hm[1])
	if err != nil {
		return 0, 0, errors.New("error parsing minute. " + err.Error())
	}

	return hInt, mInt, nil

}
