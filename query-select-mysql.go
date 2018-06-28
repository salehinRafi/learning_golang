package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func QuerySelectResult() {
	fmt.Println("------Query Select Connect To MySql Sections------")

	// Connect to Db
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/mrsm")

	if err != nil {
		panic(err.Error())

	} else {
		fmt.Println("Connected")
	}

	rows, err := db.Query("SELECT id,role_id,name FROM users")

	checkErr(err)

	for rows.Next() {
		var id int
		var roleID int
		var name string

		err = rows.Scan(&id, &roleID, &name)
		checkErr(err)
		fmt.Println(id)
		fmt.Println(roleID)
		fmt.Println(name)
	}
	defer db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
