package apis

import (
	"net/http"
	. "models"
	. "utils"
	"github.com/gin-gonic/gin"
	"strconv"
)



func IndexApi(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Welcome": "Hugo"})
	//c.HTML(http.StatusOK, "index.html", gin.H{
	//	"title": "hello  hugo",
	//})
	//c.Header("Content-Type", "text/html; charset=utf-8")
	//c.String(200, `<p>hello hugo</p>`)
}

type ApiPerson struct {
	Id  int
	Username string
	Passwd   string
}

func AddUserApi(c *gin.Context) {

	var PInfo ApiPerson
	if err := c.BindJSON(&PInfo); err != nil {
		Error.Println("取不到参数")
		c.JSON(http.StatusOK, gin.H{"retcode": 1, "stderr": "取不到参数"})
	} else {
		p := User{Username: PInfo.Username, Passwd: PInfo.Passwd}
		r, err := p.AddUser()
		if err != nil {
			Error.Println(err)
			c.JSON(http.StatusOK, gin.H{"retcode": 1, "stderr": err})

		}else {
			Info.Println("用户注册成功")
			c.JSON(http.StatusOK, gin.H{"retcode": 0, "stdout": r})
		}
	}
}

func GetUserApi(c *gin.Context) {

	r, err := GetUser()
	if err != nil {
		Error.Println(err)
		c.JSON(http.StatusOK, gin.H{"retcode": 1, "stderr": err})

	}else {
		Info.Println("获取用户信息成功")
		c.JSON(http.StatusOK, gin.H{"retcode": 0, "stdout": r})
	}
}

func QueryUserApi(c *gin.Context) {
	id := c.Query("id")
	nid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"retcode": 1, "stderr": err})
	}
	r, err := QueryUser(nid)
	if err != nil {
		Error.Println(err)
		c.JSON(http.StatusOK, gin.H{"retcode": 1, "stderr": err})

	}else {
		Info.Println("根据id查询用户信息成功")
		c.JSON(http.StatusOK, gin.H{"retcode": 0, "stdout": r})
	}
}

func ModUserApi(c *gin.Context) {

	var PInfo ApiPerson
	if err := c.BindJSON(&PInfo); err != nil {
		Error.Println("取不到参数")
		c.JSON(http.StatusOK, gin.H{"retcode": 1, "stderr": "取不到参数"})
	} else {
		p := User{Id: PInfo.Id, Username: PInfo.Username, Passwd: PInfo.Passwd}
		r, err := p.ModUser()
		if err != nil {
			Error.Println(err)
			c.JSON(http.StatusOK, gin.H{"retcode": 1, "stderr": err})

		}else {
			Info.Println("修改用户信息成功")
			c.JSON(http.StatusOK, gin.H{"retcode": 0, "stdout": r})
		}
	}
}


func DelUserApi(c *gin.Context) {
	id := c.Query("id")
	nid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"retcode": 1, "stderr": err})
	}
	r, err := DelUser(nid)
	if err != nil {
		Error.Println(err)
		c.JSON(http.StatusOK, gin.H{"retcode": 1, "stderr": err})

	}else {
		Info.Println("根据id删除用户成功")
		c.JSON(http.StatusOK, gin.H{"retcode": 0, "stdout": r})
	}
}

type ParamIndex struct {
	Instanc  string
	Sql  string
}

func IndexQuery(c *gin.Context) {

	var p ParamIndex
	if err := c.BindJSON(&p); err != nil {
		Error.Println("取不到参数")
		c.JSON(http.StatusOK, gin.H{"retcode": 1, "stderr": "取不到参数"})
	} else {
		if p.Instanc == "Advert" {
			r, err := QueryAdvertIndex(p.Instanc, p.Sql)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"retcode": 1, "stderr": err})
			} else {
				c.JSON(http.StatusOK, gin.H{"retcode": 0, "stdout": r})
			}
		} else if p.Instanc == "Eshop" {
			r, err := QueryEshopIndex(p.Instanc, p.Sql)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"retcode": 1, "stderr": err})
			} else {
				c.JSON(http.StatusOK, gin.H{"retcode": 0, "stdout": r})
			}
		} else {
			c.JSON(http.StatusOK, gin.H{"retcode": 1, "stderr": "找不到对应的数据库实例"})
		}
	}
}