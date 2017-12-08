package dsn

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	. "utils"
)

var SqlDB *sql.DB

func init() {
	var err error
	SqlDB, err = sql.Open("mysql", "root:hugo9091@tcp(127.0.0.1:3306)/gin_api?parseTime=true")
	if err != nil {
		Error.Fatal(err.Error())
	}
	err = SqlDB.Ping()
	if err != nil {
		Error.Fatal(err.Error())
	}
}
