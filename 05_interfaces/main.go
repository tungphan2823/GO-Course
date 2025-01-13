package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"

	"example.com/note/todo"
)

type saver interface {
	Save() error
}
type outputtable interface {
	saver
	Display()
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Error saving data")
		return err
	}
	fmt.Println("Success")
	return nil
}
func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")
	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = outputData(todo)
	if err != nil {
		return
	}
	fmt.Print("Saving the todo succeeded!")

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userNote)

	if err != nil {

		return
	}

	fmt.Println("Saving the note succeeded!")
}
func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}
func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
