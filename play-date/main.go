package main

import (
	"fmt"

	"time"

	"github.com/araddon/dateparse"
)

var a interface{}

func main() {
	str := "2019-03-08T03:50:00+07:00"
	str2 := "2019-03-06 08:23:00"
	a = true

	switch v := a.(type) {
	case bool:
		fmt.Println("type bool: ", v)
	case string:
		fmt.Println("type string: ", v)
	default:
		fmt.Println("not bool: ", v)
	}

	t, err := dateparse.ParseAny(str)
	if err != nil {
		fmt.Println(err.Error())
	}

	t2, err := dateparse.ParseAny(str2)
	if err != nil {
		fmt.Println(err.Error())
	}

	loc := t.Location()
	loc2 := t2.Location()

	fmt.Println("Location 1", loc.String())
	fmt.Println("Location 2", loc2.String())

	fmt.Println(t)
	//fmt.Println(t.UnixNano() / int64(time.Millisecond))
	fmt.Println(t.UnixNano() / int64(time.Millisecond))
	fmt.Println("---------------")
	fmt.Println(t2)
	fmt.Println(t2.Add(time.Hour*time.Duration(-7)).UnixNano() / int64(time.Millisecond))

	fmt.Println(convertCharonSellerClosedFrom("2020-05-19T17:00:00Z"))
}

func convertDateStringToTime(str string) (*time.Time, error) {
	layout := "2006-01-02"
	t, err := time.Parse(layout, str)

	if err != nil {
		return nil, err
	}
	return &t, nil
}

func convertCharonSellerClosedFrom(closedFrom string) string {
	closedFromDate, _ := time.Parse(time.RFC3339, closedFrom)
	convertedDate := closedFromDate.Format("2006-01-02 15:04:05")
	return convertedDate
}
