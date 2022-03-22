package models

import (
	"backlog/dao"
)

type Todo struct {
	//gorm.Model
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"` // 表示待办事情是否完成
}

/*
 所有和表相关的正删改查
*/

// CreateTodo 创建todo
func CreateTodo(todo *Todo) (err error) {
	if err = dao.DB.Create(&todo).Error; err != nil {
		return err
	}
	return
}

// GetTodoList 获取todolist
func GetTodoList() (todoList []Todo, err error) {
	if err = dao.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}

// GetTodoById 根据ID获取todo
func GetTodoById(id string) (todo *Todo, err error) {
	todo = new(Todo)
	if err = dao.DB.Where("id=?", id).First(todo).Error; err != nil {
		return nil, err
	}
	return
}

// UpdateTodo 更新todo
func UpdateTodo(todo *Todo) (err error) {
	err = dao.DB.Save(todo).Error
	return
}

// DeleteTodoById 根据ID删除todo
func DeleteTodoById(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}
