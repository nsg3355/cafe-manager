package user

import (
	"github.com/gin-gonic/gin"
)

func PostSignup(c *gin.Context) {
	result := postSignup(c)
	c.JSON(result.Meta.Code, result)
}
