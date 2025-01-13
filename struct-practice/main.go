package main

import (
	"bufio"
	"example/practice/note"
	"fmt"
	"os"
	"strings"
)

func getNoteData() (title, content string) {
	title = getUserInput("Title: ")

	content = getUserInput("Content: ")

	return title, content
}
func main() {
	title, content := getNoteData()
	note, err := note.New(title, content)
	if err != nil {
		fmt.Println("Error creating note:", err)
		return
	}
	note.Display()

}
func getUserInput(prompt string) string {

	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "/n")
	text = strings.TrimSuffix(text, "/r")
	return text
}
