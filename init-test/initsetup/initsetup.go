package initsetup

import "log"

func init() {
	log.Println("Message from init")
}

func Construct() {
	log.Println("construct from init")
}
