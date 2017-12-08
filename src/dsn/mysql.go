package dsn

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	. "utils"
	"os"
)

var AdvertSqlDB *sql.DB
var EshopSqlDB *sql.DB


func init() {
	// 根据不同的环境读取数据库连接
	var env string
    env = os.Getenv("GOENV")
    Info.Println(env)
	var err error

	if env == "prod" {
		AdvertSqlDB, err = sql.Open("mysql", "root:hugo9091@tcp(127.0.0.1:3306)/gin_api?parseTime=true")
		if err != nil {
			Error.Fatal(err.Error())
		}
		err = AdvertSqlDB.Ping()
		if err != nil {
			Error.Fatal(err.Error())
		}

		EshopSqlDB, err = sql.Open("mysql", "root:hugo9091@tcp(127.0.0.1:3306)/gin_api?parseTime=true")
		if err != nil {
			Error.Fatal(err.Error())
		}
		err = EshopSqlDB.Ping()
		if err != nil {
			Error.Fatal(err.Error())
		}
	} else if  env == "test" {
		AdvertSqlDB, err = sql.Open("mysql", "root:hugo9091@tcp(127.0.0.1:3306)/gin_api?parseTime=true")
		if err != nil {
			Error.Fatal(err.Error())
		}
		err = AdvertSqlDB.Ping()
		if err != nil {
			Error.Fatal(err.Error())
		}

		EshopSqlDB, err = sql.Open("mysql", "root:hugo9091@tcp(127.0.0.1:3306)/gin_api?parseTime=true")
		if err != nil {
			Error.Fatal(err.Error())
		}
		err = EshopSqlDB.Ping()
		if err != nil {
			Error.Fatal(err.Error())
		}
	} else {
		AdvertSqlDB, err = sql.Open("mysql", "root:hugo9091@tcp(127.0.0.1:3306)/gin_api?parseTime=true")
		if err != nil {
			Error.Fatal(err.Error())
		}
		err = AdvertSqlDB.Ping()
		if err != nil {
			Error.Fatal(err.Error())
		}

		EshopSqlDB, err = sql.Open("mysql", "root:hugo9091@tcp(127.0.0.1:3306)/gin_api?parseTime=true")
		if err != nil {
			Error.Fatal(err.Error())
		}
		err = EshopSqlDB.Ping()
		if err != nil {
			Error.Fatal(err.Error())
		}
	}
}
