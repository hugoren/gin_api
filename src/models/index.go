package models

import (
	db "dsn"
	"log"
)

type Index struct {
	Table string
	Non_unique int
	Key_name string
	Seq_in_index int
	Column_name string
	Collation string
	Cardinality int
	Sub_part string
	Packed string
	Null string
	Index_type string
	Comment string
	Index_comment string
}

func QueryDbIndex(instance string, sql string) (indexs []Index, err error) {
	indexs = make([]Index, 0)
	i := instance
	s := sql
	log.Println(i, s)
	rows, err := db.SqlDB.Query("show index from gin_api.user")
	defer rows.Close()
	if err != nil {
		return indexs, err
	}
	for rows.Next() {
		var i Index
		rows.Scan(&i.Table, &i.Non_unique, &i.Key_name,
			&i.Seq_in_index, &i.Column_name, &i.Collation,
				&i.Cardinality, &i.Sub_part, &i.Packed, &i.Null,
					&i.Index_type, &i.Comment, &i.Index_comment)
		indexs = append(indexs, i)
	}
	if err = rows.Err(); err != nil {
		return indexs, err
	}
	return indexs, err
}