package todo

import "fmt"

type Todo struct {
	Text string `json:"text"`
}

func New(text string) Todo {
	return Todo{
		Text: text,
	}
}

func (todo Todo) Display() {
	fmt.Println(todo.Text)
}
