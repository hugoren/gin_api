package utils

import (
	"database/sql"
	"fmt"
	"github.com/kshvakov/clickhouse"
	"log"
)

var db = &sql.DB{}

func  init()  {
	var url = "tcp://192.168.0.23:9000?debug=true"
	db, _ = sql.Open("clickhouse", url)

}

func CreateTable(){
	var url = "tcp://192.168.0.23:9000?debug=true&database=db_hystrix"
	connect, err := sql.Open("clickhouse", url)
	if err != nil {
		log.Fatal(err)
	}
	if err := connect.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			fmt.Println(err)
		}
		return
	}

	_, err = connect.Exec(`
		CREATE TABLE IF NOT EXISTS hystrix (
			uuid                String,
			appName             String,
			interfaceName       String,
			shortCircuited      String,
			rollingCountTimeout Float64,
			rollingCountFailure Float64,
			th99                Float64,
			note                String,
			insertDate          String
		) engine=Memory
	`)

	if err != nil {
		log.Fatal(err)
	}

}

func InsertData(uuid string, appName string, interfaceName string, shortCircuited string, rollingCountTimeout float64, rollingCountFailure float64, th99 float64, note string, insertDate string) {
	var url= "tcp://192.168.0.23:9000?database=db_hystrix"
	connect, err := sql.Open("clickhouse", url)
	if err != nil {
		log.Fatal(err)
	}
	if err := connect.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			fmt.Println(err)
		}
		return
	}

	var (
		tx, _= connect.Begin()
		stmt, _= tx.Prepare("INSERT INTO hystrix (uuid, appName, interfaceName, shortCircuited, rollingCountTimeout, rollingCountFailure, th99, note, insertDate) VALUES (?,?,?,?,?,?,?,?,?)")
	)
	defer stmt.Close()

	if _, err := stmt.Exec(
		uuid,
		appName,
		interfaceName,
		shortCircuited,
		rollingCountTimeout,
		rollingCountFailure,
		th99,
		note,
		insertDate,
	); err != nil {
		log.Fatal(err)
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)

	}

}
