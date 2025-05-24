package main

import (
	"fmt"
	"log"

	"github.com/aliadelharrat/note-taking-app/input"
	"github.com/aliadelharrat/note-taking-app/note"
	"github.com/aliadelharrat/note-taking-app/todo"
)

func main() {
	title := input.GetUserInput("Note title: ")
	content := input.GetUserInput("Note content: ")
	new_note := note.New(title, content)
	OutputData(&new_note, "note")

	todo_content := input.GetUserInput("Todo content: ")
	new_todo := todo.New(todo_content)
	OutputData(&new_todo, "todo")
}

type Saver interface {
	Save() error
	Display()
}

func saveData(model Saver) error {
	err := model.Save()
	if err != nil {
		return err
	}
	return nil
}

func OutputData(model Saver, text string) {
	err := saveData(model)
	if err != nil {
		log.Printf("Saving %v failed\n", text)
	} else {
		fmt.Printf("%s saved successfully\n", text)
	}

	model.Display()
}
