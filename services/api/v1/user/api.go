package user

import (
	"github.com/gin-gonic/gin"
)

func PostSignup(c *gin.Context) {
	result := postSignup(c)
	c.JSON(result.Meta.Code, result)
}

func PostLogin(c *gin.Context) {
	result := postLogin(c)
	c.JSON(result.Meta.Code, result)
}

func PostLogout(c *gin.Context) {
	result := postLogout(c)
	c.JSON(result.Meta.Code, result)
}

func GetVerification(c *gin.Context) {
	result := getVerification(c)
	c.JSON(result.Meta.Code, result)
}
