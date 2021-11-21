package main

import (
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"pgd/controller/notice"
	"pgd/controller/order"
	"pgd/controller/user"
)

type login struct {
	ID       string `form:"id" json:"id" binding:"required"`
	Password string `form:"pwd" json:"pwd" binding:"required"`
}

var identityKey = "id"

func main() {
	gin.ForceConsoleColor()

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	authMiddleware, err := middleware()
	if err != nil {
		log.Fatal("JWT 错误:" + err.Error())
	}

	r.POST("/login", authMiddleware.LoginHandler)
	r.POST("/register", user.Register)
	r.GET("/refresh_token", authMiddleware.RefreshHandler)

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	r.Use(authMiddleware.MiddlewareFunc())

	noticeGroup := r.Group("/notice") //通知
	{
		noticeGroup.POST("", notice.New)
		noticeGroup.DELETE("", notice.Delete)
		noticeGroup.GET("", notice.List)
	}

	orderGroup := r.Group("/order")

	orderGroup.Use(order.Handler)
	{
		orderGroup.GET("", order.Detail)            //订单详情
		orderGroup.GET("/can_receive", order.List)  //可接订单列表
		orderGroup.GET("/receive_list", order.List) //我接的单
		orderGroup.GET("/list", order.List)         //我下的单
		orderGroup.POST("", order.New)              //下单
		orderGroup.PUT("/receive", order.Receive)   //接单
		orderGroup.DELETE("/", order.Delete)        //删除订单
		orderGroup.PUT("/complete", order.Complete) //确认完成
		orderGroup.PUT("/cancel", order.Cancel)     //取消接单
	}

	if err := http.ListenAndServe(":9090", r); err != nil {
		log.Fatal(err)
	}
}
