package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nsg3355/ph-cafe-manager/services/api/v1/product"
	"github.com/nsg3355/ph-cafe-manager/services/api/v1/user"
)

// 메인 라우터
func InitRouter() *gin.Engine {
	r := gin.Default()

	subRg := r.Group("cafe-mgr/api")
	{
		subRg.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
	InitV1Router(subRg)

	return r
}

func InitV1Router(rg *gin.RouterGroup) {
	rg.Use()
	{
		rgV1 := rg.Group("/v1")
		InitUserRouter(rgV1)
		InitProductRouter(rgV1)
	}
}

func InitUserRouter(rg *gin.RouterGroup) {
	rg.Use()
	{
		rgSub := rg.Group("/user")
		{
			rgSub.POST("/signup", user.PostSignup)
			rgSub.POST("/login", user.PostLogin)
			rgSub.POST("/logout", user.PostLogout)
			rgSub.GET("/verification", user.GetVerification)
		}
	}
}

func InitProductRouter(rg *gin.RouterGroup) {
	rg.Use()
	{
		rgSub := rg.Group("/product")
		{
			rgSub.GET("/list", product.GetList)
			rgSub.GET("/byid", product.GetByid)
			rgSub.POST("/item", product.PostItem)
			rgSub.PUT("/item", product.PutItem)
			rgSub.DELETE("/item", product.DeleteItem)
		}
	}
}
