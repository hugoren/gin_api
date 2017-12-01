package models

import (
	db "dsn"
	"fmt"
	. "utils"
)

type User struct {
	Id        int
	Username string
	Passwd  string
}

func (p *User) AddUser() (result string, err error) {
	rs, err := db.SqlDB.Exec("INSERT INTO user(username, passwd) VALUES (?, ?)", p.Username, p.Passwd)
	fmt.Println(rs)
	if err != nil {
		return "用户注册失败", err
	}
	return "用户注册成功", err
}

func GetUser() (persons []User, err error) {
	persons = make([]User, 0)
	rows, err := db.SqlDB.Query("SELECT * FROM user")
	defer rows.Close()
	if err != nil {
		return persons, err
	}
	for rows.Next() {
		var person User
		rows.Scan(&person.Id, &person.Username, &person.Passwd)
		persons = append(persons, person)
	}
	if err = rows.Err(); err != nil {
		return persons, err
	}
	return persons, err
}

func QueryUser(id int) (p User, err error) {
	var person User
	db.SqlDB.QueryRow("SELECT * FROM person where id=?", id).Scan(
		&person.Id, &person.Username, &person.Passwd,
	)
	return  person, err
}


func (p *User) ModUser() (result string, err error) {
	//var person Person
	stmt, err := db.SqlDB.Prepare("UPDATE user SET username=?, passwd=? WHERE id=?")
	defer stmt.Close()
	if err != nil {
		return "预修改客户信息失败", err
	}
	rs, err := stmt.Exec(p.Username, p.Passwd, p.Id)
	if err != nil {
		return "修改客户信息失败", err
	}
	r, err := rs.RowsAffected()
	if err != nil {
		Info.Println(r)
		return "执行修改客户信息失败", err
	}

	return "修改客户信息成功" , err
}