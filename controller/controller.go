package controller

import (
	"backlog/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
url  --> controller --> logic/service --> model/dao             redis/mysql等连接操作放在rpc目录下面
请求  -->   控制器   -->     业务逻辑    -->  模型层的增删改查
*/

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateTodo(c *gin.Context) {
	// 前端填写待办事项，点击提交请求
	// 1、获取请求中的数据
	var todo models.Todo
	c.BindJSON(&todo)
	// 2、将数据写入数据库
	if err := models.CreateTodo(&todo); err != nil {
		// 3、返回响应
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func GetTodoList(c *gin.Context) {
	//查询 todos 表里面所有的数据
	if todoList, err := models.GetTodoList(); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func UpdateTodoById(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}

	todo, err := models.GetTodoById(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.BindJSON(&todo) // 将请求传入的数据填充如todo

	if err = models.UpdateTodo(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, todo)
	}

}
func DeleteTodoById(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	if err := models.DeleteTodoById(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}
