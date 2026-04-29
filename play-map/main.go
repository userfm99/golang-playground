package main

import (
	"log"
	"strconv"
)

var arrayMap []map[string]string
var mapvalue map[string]string
var counter = 0

func init() {
	arrayMap = make([]map[string]string, 0)
	mapvalue = make(map[string]string)
	arrayMap = append(arrayMap, mapvalue)
}

func main() {
	players := map[int]string{1: "Keeper", 2: "Pivot", 3: "Wing 1", 4: "Wing 2", 5: "Target"}
	log.Println(players[3])

	v, ok := players[110]
	if ok {
		log.Println("value ", v)
	} else {
		log.Println("value not found")
	}

	i := 0
	for i < 10000 {
		addField("key" + strconv.Itoa(i), "value" + strconv.Itoa(i))
		i++
	}
	log.Println("LENGTH OF arrayMap", len(arrayMap))
	mv := arrayMap[9]
	if v, exist := mv["key8600"]; exist {
		log.Println("random fetch mapvalue", v)
	}
}


func addField(key string, value string) {
	mapvalue[key] = value

	if counter > 1000 {
		log.Println(arrayMap)
		mapvalue = make(map[string]string)
		arrayMap = append(arrayMap, mapvalue)
		counter = 0
	}

	counter ++
}