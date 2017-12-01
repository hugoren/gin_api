package models

import (
	db "dsn"
	"fmt"
)

type Person struct {
	Id        int    `json:"id" form:"id"`
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
}

func (p *Person) AddPerson() (result string, err error) {
	rs, err := db.SqlDB.Exec("INSERT INTO person(firstname, lastname) VALUES (?, ?)", p.FirstName, p.LastName)
	fmt.Println(rs)
	if err != nil {
		return "新增客户失败", err
	}
	return "新增客户成功", err
}

func GetPerson() (persons []Person, err error) {
	persons = make([]Person, 0)
	rows, err := db.SqlDB.Query("SELECT id, firstname, lastname FROM person")
	defer rows.Close()
	if err != nil {
		return persons, err
	}
	for rows.Next() {
		var person Person
		rows.Scan(&person.Id, &person.FirstName, &person.LastName)
		persons = append(persons, person)
	}
	if err = rows.Err(); err != nil {
		return persons, err
	}
	return persons, err
}

func QueryPerson(id int) (p Person, err error) {
	var person Person
	db.SqlDB.QueryRow("SELECT * FROM person where id=?", id).Scan(
		&person.Id, &person.FirstName, &person.LastName,
	)
	return  person, err
}