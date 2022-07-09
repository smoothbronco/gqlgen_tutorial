package model

import (
	"fmt"

	"github.com/smoothbronco/gqlgen_tutorial/entity"
)

type Todo struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserID string `json:"user"`
}

func NewTodoFromEntity(e *entity.Todo) *Todo {
	return &Todo{
		ID:     fmt.Sprintf("%d", e.ID),
		Text:   e.Text,
		Done:   e.Done,
		UserID: e.UserID,
	}
}
