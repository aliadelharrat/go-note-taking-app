package note

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) Note {
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}
}

func (note *Note) Display() {
	fmt.Printf("Title: %v\n", note.Title)
	fmt.Printf("Content: %v\n", note.Content)
	fmt.Printf("CreatedAt: %v\n", note.CreatedAt)
}

func (note *Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	path := filepath.Join("saved", "notes")
	path = filepath.Join(path, fileName)

	json, err := json.Marshal(note)
	if err != nil {
		return err
	}

	return os.WriteFile(path, json, 0644)
}
