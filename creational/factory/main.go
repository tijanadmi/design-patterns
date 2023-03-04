package main

import "log"

func main() {
	farm(60)
}

//interface
type Animal interface {
	Sound() string
}

//cat
type Cat struct {
}

func NewCat() Cat{
	return Cat{}
}

func (c Cat) Sound() string {
	return "mjau"
}

type Dog struct {
	name string
}

func NewDog() *Dog{
	return &Dog{name: "Dino"}
}
func (d Dog) Sound() string {
	return "vuf"
}

func farm(x int){
	var a Animal
	if x>42 {
		a=NewCat()
	} else{
		a=NewDog()
	}

	log.Print(a.Sound())
}

