package note

import (
	"encoding/json"
	"errors"
	"fmt"

	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string
	Content   string
	CreatedAt time.Time
}

func (note Note) Display() {
	fmt.Printf("Your note title: %v has the following content: \n\n%v", note.Title, note.Content)
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Invalid input.")
	}
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(note)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}
