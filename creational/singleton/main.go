package main

import (
	"design-patterns/creational/singleton/service"
	"log"
)
func main(){
	/*var singleton service.IdServiceSingleton
	s1 := singleton.GetService()
	s2 := singleton.GetService()
	buildCar(s1.Next())
	buildCar(s1.Next())
	buildCar(s2.Next())*/
	singleton := service.GetInstance()
	buildCar(singleton.AddOne())
	buildCar(singleton.AddOne())
	buildCar(singleton.AddOne())
}

func buildCar(id int){
	log.Print("car: ", id)
}