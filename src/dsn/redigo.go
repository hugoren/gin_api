package dsn

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
	. "utils"
)

func RedisConn()  {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		Error.Println(err)
	}
	defer c.Close()

	ret, _ := c.Do("SET","hugo", "boss")
	fmt.Printf("%s\n", ret)

	ret, _ = c.Do("GET","hu")
	fmt.Printf("%s\n", ret)
}
