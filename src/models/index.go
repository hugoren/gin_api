package models

import (
	db "dsn"
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

func QueryAdvertIndex(instance string, sql string) (indexs []Index, err error) {
	indexs = make([]Index, 0)
	rows, err := db.AdvertSqlDB.Query(sql)
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

func QueryEshopIndex(instance string, sql string) (indexs []Index, err error) {
	indexs = make([]Index, 0)
	rows, err := db.EshopSqlDB.Query(sql)
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