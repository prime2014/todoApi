package todo

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	Title       string     `json:"title" gorm:"unique;index"`
	Slug        *string    `json:"slug"`
	Description *string    `json:"description"`
	CreatedAt   *time.Time `json:"created_at" gorm:"index;autoCreateTime"`
	UpdatedAt   *time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (todo *Todo) Create(tx *gorm.DB) (Todo, error) {
	result := tx.Create(&todo)

	if result.Error != nil {
		return Todo{}, result.Error
	}

	return *todo, nil
}

func (todo *Todo) FetchAll(tx *gorm.DB) ([]Todo, error) {
	todoData := []Todo{}
	result := tx.Order("created_at DESC").Find(&todoData)

	fmt.Println(result)
	if result.Error != nil {
		return []Todo{}, result.Error
	}
	return todoData, nil
}
