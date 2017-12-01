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
		c.JSON(http.StatusUnauthorized, gin.H{"retcode": 1, "stderr": "取不到参数"})
	} else {
		p := User{Username: PInfo.Username, Passwd: PInfo.Passwd}
		r, err := p.AddUser()
		if err != nil {
			Error.Println(err)
			c.JSON(http.StatusOK, gin.H{"retcode": 1, "stderr": err})

		}else {
			Info.Println("新增客户成功")
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
		Info.Println("根据id查询客户成功")
		c.JSON(http.StatusOK, gin.H{"retcode": 0, "stdout": r})
	}
}

func ModUserApi(c *gin.Context) {

	var PInfo ApiPerson
	if err := c.BindJSON(&PInfo); err != nil {
		Error.Println("取不到参数")
		c.JSON(http.StatusUnauthorized, gin.H{"retcode": 1, "stderr": "取不到参数"})
	} else {
		p := User{Id: PInfo.Id, Username: PInfo.Username, Passwd: PInfo.Passwd}
		r, err := p.ModUser()
		if err != nil {
			Error.Println(err)
			c.JSON(http.StatusOK, gin.H{"retcode": 1, "stderr": err})

		}else {
			Info.Println("修改客户信息成功")
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