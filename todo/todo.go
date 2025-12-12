package todo

// todo
// add & remove task feature
// more useful data (priority, dueTime, users?)

import (
	"fmt"
	"time"
)

// todo:
// - timestamp
// - prazo
// - lógica para reproduzir a tarefa em um período especificado
type Todo struct {
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
}

func New(text string) (Todo, error) {
	return Todo{
		Text:      text,
		CreatedAt: time.Now(),
	}, nil
}

func (todo Todo) Display() {
	fmt.Println(todo.Text)
}
