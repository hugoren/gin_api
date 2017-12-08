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

func GetUser() (users []User, err error) {
	users = make([]User, 0)
	rows, err := db.SqlDB.Query("SELECT * FROM user")
	defer rows.Close()
	if err != nil {
		return users, err
	}
	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.Username, &user.Passwd)
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return users, err
	}
	return users, err
}

func QueryUser(id int) (p User, err error) {
	var u User
	db.SqlDB.QueryRow("SELECT * FROM user where id=?", id).Scan(
		&u.Id, &u.Username, &u.Passwd,
	)
	return  u, err
}


func (p *User) ModUser() (result string, err error) {
	//var person Person
	stmt, err := db.SqlDB.Prepare("UPDATE user SET username=?, passwd=? WHERE id=?")
	defer stmt.Close()
	if err != nil {
		return "预修改用户信息失败", err
	}
	rs, err := stmt.Exec(p.Username, p.Passwd, p.Id)
	if err != nil {
		return "修改客户信息失败", err
	}
	r, err := rs.RowsAffected()
	if err != nil {
		Error.Println(r)
		return "执行修改用户信息失败", err
	}

	return "修改用户信息成功" , err
}


func DelUser(id int) (s string, err error) {
	rs, err := db.SqlDB.Exec("DELETE FROM user WHERE id=?", id)
	if err != nil {
		Error.Println(err)
		return "删除用户信息失败", err
	}
	r, err := rs.RowsAffected()
	if err != nil {
		Error.Println(r)
		return "删除用户信息失败", err
	}
	return  "删除用户信息成功", err
}
