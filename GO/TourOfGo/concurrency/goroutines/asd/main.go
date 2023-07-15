package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/lorelval/goroutiner/book"
)

func main() {

	var cache = map[int]Book{}
	var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 5; i++ {
		id := rnd.Intn(5) + 1
		if b, ok := queryCache(id); ok {
			fmt.Println("from cache")
			fmt.Println(b)
			continue
		}
		if b, ok := queryDatabase(id); ok {
			fmt.Println("from database")
			fmt.Println(b)
			continue
		}
		fmt.Printf("Book not found with ID: '%v'", id)
		time.Sleep(150 * time.Millisecond)
	}
}

func queryCache(id int) (book.Book, bool) {
	b, ok := cache[id]
	return b, ok
}

func queryDatabase(id int) (book.Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range books {
		if b.ID == id {
			cache[id] = b
			return b, true
		}

	}
	return Book{}, false
}
