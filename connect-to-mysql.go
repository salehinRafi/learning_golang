package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToMySql() {
	fmt.Println("------Connect To MySql Sections------")
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/mrsm")

	if err != nil {
		panic(err.Error())

	}

	fmt.Println("Connected")
	defer db.Close()
}
