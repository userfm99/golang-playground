package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgres://postgres:test123@localhost/postgres"
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	id := generateMemberID(1)
	fmt.Println(id)

	clientID := uuid.New().String()
	fmt.Println("client id", clientID)
}

func generateMemberID(id int64) string {
	prefix := "HM0000000"
	lenPrefix := len(prefix)
	idStr := strconv.Itoa(int(id))
	lenIDStr := len(idStr)
	runes := []rune(prefix)

	return string(runes[0:lenPrefix-lenIDStr]) + idStr
}
