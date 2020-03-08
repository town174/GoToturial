package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var dbName = "mysql"
var dbConn = "root:123456@tcp(127.0.0.1:3306)/SmartBookShelf"

//func Init()  {
//	sql.Register(dbName, &mysql.MySQLDriver{})
//}

func MysqlTest() {
	//连接数据库
	db, err := sql.Open(dbName, dbConn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var (
		BookName string
		Author   string
	)

	//Fetching data from database
	rows, err := db.Query("SELECT BookName,Author FROM syncbookinfoes LIMIT 10;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&BookName, &Author)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("get book ", BookName, " author ", Author)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	//Preparing Queries
	stmt, err := db.Prepare("SELECT BookName,Author FROM syncbookinfoes where BookName like ?;")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err = stmt.Query("%溪水%")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&BookName, &Author)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("get book ", BookName, " author ", Author)
	}
}
