package todo

import (
	"fmt"

	"github.com/UmetsuJunya/todo-app-backend/backend/lib"
)

type Todo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

func New() *Todo {
	return &Todo{}
}

func (r *Todo) Add(data Todo) bool {
	db := lib.GetDBConn().DB
	if err := db.Create(data).Error; err != nil {
		fmt.Println("err!")
		return false
	}
	return true
}

func (r *Todo) GetAll() []Todo {
	db := lib.GetDBConn().DB
	var todos []Todo
	if err := db.Find(&todos).Error; err != nil {
		return nil
	}
	return todos
}

func (r *Todo) UpdateTodo(id int, data Todo) bool {
	db := lib.GetDBConn().DB
	var todo Todo
	if err := db.Where("ID = ?", id).First(&todo).Updates(data).Error; err != nil {
		fmt.Println("err!")
		return false
	}
	return true
}

func (r *Todo) UpdateStatus(id int) bool {
	db := lib.GetDBConn().DB

	var todo Todo
	if err := db.Where("id = ?", id).First(&todo).Error; err != nil {
		return false
	}

	// statusを切り替えて更新
	if err := db.Model(&todo).Where("id = ?", id).Update("STATUS", !todo.Status).Error; err != nil {
		return false
	}
	return true
}

func (r *Todo) Delete(id int) bool {
	db := lib.GetDBConn().DB
	var todo Todo
	if err := db.Where("id = ?", id).Delete(&todo).Error; err != nil {
		return false
	}
	return true
}
