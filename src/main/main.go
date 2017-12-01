package main

import (
	db "dsn"
)

func main() {

	defer db.SqlDB.Close()
	router := initRouter()
	router.Run(":30000")

}
