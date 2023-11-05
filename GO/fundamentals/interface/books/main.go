package main

import "fmt"

type Book struct {
	Author string
	Title  string
}

type Magazine struct {
	Title string
	Issue int
}

func (b *Book) Assign(n, t string) {
	b.Author = n
	b.Title = t
}

// recomendación usar puntero para casos de modificar y tambien para leer/mostrar
func (b *Book) Print() {
	fmt.Printf("Author: %s, Title: %s\n", b.Author, b.Title)
}

func (m *Magazine) Assign(t string, i int) {
	m.Title = t
	m.Issue = i
}

func (m *Magazine) Print() {
	fmt.Printf("Title: %s, Issue: %d\n", m.Title, m.Issue)
}

type Printer interface {
	Print()
}

func main() {
	var (
		b Book
		m Magazine
		i Printer
	)
	b.Assign("Jack Rabbit", "Book of Rabbits") // Assign values to b via method
	b.Assign("ab", "hola")
	m.Assign("Rabbit Weekly", 26) // Assign values to m via method

	fmt.Println("Call interface")
	i = &b    // Method has pointer receiver, interface does not
	i.Print() // Show book values via the interface
	i = &m    // Magazine also satisfies shower interface
	i.Print() // Show magazine values via the interface

}
