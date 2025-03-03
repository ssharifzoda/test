package utils

import "fmt"

var data map[int]string

type Person struct {
	Name    string
	Surname string
	Age     int
}

type Note struct {
	Id   int
	Text map[int]Person
}

func (n Note) AddNote() {

}

func (n Note) DeleteNote() {
	delete(data, n.Id)
}

func (n Note) List() {
	fmt.Println(data)
}

func (n Note) UpdateNote(newText string) {
	data[n.Id] = newText
}
