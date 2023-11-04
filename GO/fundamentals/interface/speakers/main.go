package main

import "fmt"

// EXAMPLE --> NATIVE OR FOREIGN SPEAKER
// Speaker interface : is the global type that assembles the others structs below
type Speaker interface {
	Arabic()
	English()
	French()
}
type NativeSpeaker struct {
	Name string
}
type Foreignpeaker struct {
	Name        string
	Nationality string
}

// methods for NativeSpeaker struct
func (n NativeSpeaker) Arabic() {

}
func (n NativeSpeaker) English() {

}
func (n NativeSpeaker) French() {

}

// methods for ForeingSpeaker struct
func (f Foreignpeaker) Arabic() {

}
func (f Foreignpeaker) English() {

}
func (f Foreignpeaker) French() {

}

// all 3 methods must be defined in order to belong to the Speaker type.
func main() {
	listSpeakers := make([]Speaker, 0)
	Sp1 := NativeSpeaker{Name: "Su"}
	Sp2 := Foreignpeaker{Name: "Lorena"}
	listSpeakers = append(listSpeakers, Sp1, Sp2)
	fmt.Println(listSpeakers)
}
