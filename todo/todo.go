package todo

import "fmt"

type Todo struct {
	Text string `json:"text"`
}

func New(text string) (Todo, error) {
	return Todo{
		Text: text,
	}, nil
}

func (todo Todo) Display() {
	fmt.Println(todo.Text)
}
