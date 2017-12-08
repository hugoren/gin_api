package main

import (
	db "dsn"
	"net/http"
	"time"
)

func main() {

	db.RedisConn()
	defer db.AdvertSqlDB.Close()
	defer db.EshopSqlDB.Close()
	router := initRouter()
	//router.LoadHTMLFiles("template/index.html")
	//router.LoadHTMLGlob("/Users/admin/devops_go/gin_api/gopath/src/templates/html/*")
	//router.Run(":30000")
	s := &http.Server{
		Addr:           ":30000",
		Handler:        router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 30,
	}
	s.ListenAndServe()



	}
