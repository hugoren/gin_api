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
	FirstName string
	LastName   string
}

func AddPersonApi(c *gin.Context) {

	var PInfo ApiPerson
	if err := c.BindJSON(&PInfo); err != nil {
		Error.Println("取不到参数")
		c.JSON(http.StatusUnauthorized, gin.H{"retcode": 1, "stderr": "取不到参数"})
	} else {
		p := Person{FirstName: PInfo.FirstName, LastName: PInfo.LastName}
		r, err := p.AddPerson()
		if err != nil {
			Error.Println(err)
			c.JSON(http.StatusOK, gin.H{"retcode": 1, "stderr": err})

		}else {
			Info.Println("新增客户成功")
			c.JSON(http.StatusOK, gin.H{"retcode": 0, "stdout": r})
		}
	}
}

func GetPersonApi(c *gin.Context) {

	r, err := GetPerson()
	if err != nil {
		Error.Println(err)
		c.JSON(http.StatusOK, gin.H{"retcode": 1, "stderr": err})

	}else {
		Info.Println("获取客户成功")
		c.JSON(http.StatusOK, gin.H{"retcode": 0, "stdout": r})
	}
}

func QueryPersonApi(c *gin.Context) {
	id := c.Query("id")
	nid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"retcode": 1, "stderr": err})
	}
	r, err := QueryPerson(nid)
	if err != nil {
		Error.Println(err)
		c.JSON(http.StatusOK, gin.H{"retcode": 1, "stderr": err})

	}else {
		Info.Println("查询客户成功")
		c.JSON(http.StatusOK, gin.H{"retcode": 0, "stdout": r})
	}
}
