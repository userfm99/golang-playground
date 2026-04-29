package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"log"
	"github.com/pkg/errors"
)


func main() {
	zlog, _ := zap.NewProduction()

	defer zlog.Sync()

	sugar := zlog.Sugar()

	db, err := sql.Open("mysql", "root:r00t@tcp(127.0.0.1:3306)/employees")
	if err != nil {
		log.Fatal("err")
		sugar.Error("Error connecting to database", "db","employees", "user", "root")
		return
	}

	errPing := db.Ping()
	if errPing != nil {
		zlog.Error("error pinging database", zap.Error(errPing))
		errors.Wrap(errPing, "Error pinging database")
		return
	}

	zlog.Info("successfully connected to db")

	rows, err := db.Query("Select emp_no, first_name from employees")
	if err != nil {
		zlog.Error("error getting employees", zap.Error(err))
		return
	}
	defer rows.Close()

	var empNo int
	var firstName string
	for rows.Next() {
		rows.Scan(&empNo, &firstName)
		zap.L().Info("Scanning result", zap.Int("Emp no", empNo), zap.String("name", firstName))
	}
	/*// defer db.Close()
	sugar.Infof("Successfully connected to db %v", db.Stats())

	rows, err := db.Query("SELECT * FROM employees")
	if err != nil {
		sugar.Error("error connecting to db")
	}
	defer rows.Close()*/
}
