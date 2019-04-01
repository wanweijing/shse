package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("ddd")
	r := gin.Default()
	r.GET("test", func(c *gin.Context) {
		c.String(200, "fff")
	})
	r.Run(":1234")
}
