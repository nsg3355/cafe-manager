package router

import (
	"github.com/gin-gonic/gin"
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
		// InitProductRouter(rgV1)
	}
}

func InitUserRouter(rg *gin.RouterGroup) {
	rg.Use()
	{
		rgSub := rg.Group("/user")
		{
			rgSub.POST("/post/signup", user.PostSignup)
			// rgSub.POST("/post/login", user.PostSignup)
			// rgSub.POST("/post/logout", user.PostSignup)
		}
	}
}

// func InitProductRouter(rg *gin.RouterGroup) {
// 	rg.Use()
// 	{
// 		rgSub := rg.Group("/product")
// 		{
// 			rgSub.POST("/get/list", product.GetUserList)
// 			rgSub.POST("/get/byid", product.GetUserList)
// 			rgSub.POST("/post/item", product.GetUserList)
// 			rgSub.POST("/put/item", product.GetUserList)
// 			rgSub.POST("/delete/item", product.GetUserList)
// 		}
// 	}
// }
