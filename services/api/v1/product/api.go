package product

import (
	"github.com/gin-gonic/gin"
)

func GetList(c *gin.Context) {
	result := getList(c)
	c.JSON(result.Meta.Code, result)
}

func GetByid(c *gin.Context) {
	result := getByid(c)
	c.JSON(result.Meta.Code, result)
}

func PostItem(c *gin.Context) {
	result := postItem(c)
	c.JSON(result.Meta.Code, result)
}

func PutItem(c *gin.Context) {
	result := putItem(c)
	c.JSON(result.Meta.Code, result)
}

func DeleteItem(c *gin.Context) {
	result := deleteItem(c)
	c.JSON(result.Meta.Code, result)
}
