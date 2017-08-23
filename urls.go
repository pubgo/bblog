package main

import (
//"github.com/gin-gonic/gin"
//"github.com/googollee/go-socket.io"
//"log"
)
import "github.com/gin-gonic/gin"

func InitUrls(r *gin.Engine) {

	//server, err := socketio.NewServer(nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//server.On("connection", socketio_conn)
	//server.On("error", socketio_error)

	r.GET("/", ping)
	r.GET("/ping", ping)
	//r.GET("/status", ping)

	// 服务资源操作
	programs := r.Group("/api/blog")
	{
		// 添加多个服务资源
		programs.POST("", programs_post)

		// 修改多个服务的信息
		programs.PUT("", programs_put)

		// 修改单个服务的信息
		programs.PUT(":name", programs_put)

		// 获取多个服务信息
		programs.GET("", programs_get)

		// 根据服务名称获取单个服务信息
		programs.GET(":name", programs_get)

		// 删除多个服务资源
		programs.DELETE("", programs_delete)

		// 删除单个服务资源
		programs.DELETE(":name", programs_delete)

	}

	//r.GET("/socket.io/", func(c *gin.Context) {
	//	server.ServeHTTP(c.Writer, c.Request)
	//})
}
