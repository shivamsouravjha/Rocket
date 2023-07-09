package db

import (
	"database/sql"
	"fmt"
	"log"
	"rocket/config"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

var Dbmap = initDb()

func initDb() *gorp.DbMap {
	connection := config.Get().DBUserName + ":" + config.Get().DBPassword + "@tcp(" + config.Get().DBHostReader + ":" + config.Get().DBPort + ")/" + config.Get().DBName
	db, err := sql.Open("mysql", connection)
	checkErr(err, "sql.Open failed")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")
	fmt.Println("connected")

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		fmt.Println(err)
		log.Fatalln(msg, err)
	}
}
