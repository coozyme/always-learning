package main

import "fmt"

type note struct {
	Title string
	Body  string
}

func main() {
	// m := map[string]string{"key1": "val1", "key2": "val2"}
	// delete(m, "key1")

	// m := make(map[string]int)
	// m["numberOne"] = 1
	// m["numberTwo"] = 2

	// notes := *note[{}]
	// notes := make(map[string]string)

	// notes["Title"] = "Booke"
	// notes["Body"] = "Booke body"

	notes := []note{
		{Title: "Booke", Body: "Bookei-body"},
		{Title: "Booke2", Body: "Bookei-body2"},
	}

	fmt.Println("not", notes)
	for _, note := range notes {
		fmt.Println("ip", note.Title)
		// thisNote := map[string]string{
		// 	"Title": note.Title,
		// 	"Body":  note.Body,
		// }

		// content["notes"] = append(content["notes"], thisNote)
	}
}
