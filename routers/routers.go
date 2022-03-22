package routers

import (
	"backlog/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// 获取 gin 引擎
	r := gin.Default()
	// 加载模板文件地址
	r.LoadHTMLGlob("templates/*")
	// 加载静态文件地址
	r.Static("/static", "static")

	r.GET("/", controller.IndexHandler)

	// v1: 处理代办事项
	v1Group := r.Group("v1")
	{
		// 添加
		v1Group.POST("/todo", controller.CreateTodo)
		// 查看所有待办事项
		v1Group.GET("/todo", controller.GetTodoList)
		// 查看某一个待办事项
		v1Group.GET("/todo/:id", func(c *gin.Context) {
			// todo
		})
		// 修改
		v1Group.PUT("/todo/:id", controller.UpdateTodoById)
		// 删除
		v1Group.DELETE("/todo/:id", controller.DeleteTodoById)
	}
	return r
}
