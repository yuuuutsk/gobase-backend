package model

import (
	"github.com/yuuuutsk/gobase-backend/app/domain"
)

type Todo struct {
	ID   domain.TodoID
	Text string
	Done bool
}

func NewTodo(
	text string,
	done bool,
) *Todo {
	return &Todo{
		Text: text,
		Done: done,
	}
}

func RestoreTodo(
	ID domain.TodoID,
	text string,
	done bool,
) *Todo {
	return &Todo{
		ID:   ID,
		Text: text,
		Done: done,
	}
}
