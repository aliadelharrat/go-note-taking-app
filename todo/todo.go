package todo

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Todo struct {
	Text string `json:"text"`
}

func New(content string) Todo {
	return Todo{
		Text: content,
	}
}

func (todo *Todo) Display() {
	fmt.Printf("Text: %v\n", todo.Text)
}

func (todo *Todo) Save() error {
	fileName := "todo.json"
	path := filepath.Join("saved", "todos")
	path = filepath.Join(path, fileName)

	json, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	return os.WriteFile(path, json, 0644)
}
