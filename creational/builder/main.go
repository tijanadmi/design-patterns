package main

import (
	"log"
)



func main() {
	me := newHuman().withEyeColor("brown").withAge(3).withHeight(164)
	you := newHumanWithFiels(30,176,"blue")
	log.Printf("%v",me)
	
	log.Printf("%v",me.reset())
	log.Printf("%v",you)

	youngGiant := giant().withAge(10)
	oldGiant := giant().withAge(100)
	log.Printf("%v",youngGiant)
	log.Printf("%v",oldGiant)


	manufacturer := Manufacturer{}

	homePCBuilder := &HomePCBuilder{}
	manufacturer.SetBuilder(homePCBuilder)
	manufacturer.ConstructPC()
	homePC := homePCBuilder.GetPC()
	log.Printf("%v",homePC)
}

type human struct {
	age      int
	height   int
	eyeColor string
}

func newHuman() human {
	return human{}
}

func newHumanWithFiels(age int, height int, eyeColor string) human{
	return human{
		age: age,
		height: height,
		eyeColor:  eyeColor,
	}
}

// Builders
func (h human) withAge(age int) human{
	h.age = age
	return h
}


func (h human) withHeight(height int) human{
	h.height = height
	return h
}

func (h human) withEyeColor(colour string) human{
	h.eyeColor = colour
	return h
}

//Reset
func (h human) reset() human {
	h= newHuman()
	return human{}
}

//Pre-defined builders
func giant() human {
	return newHuman().withHeight(180).withEyeColor("green")
}


// Product to build
type PersonalComputer struct {
	ramCap int
	hddCap int
	cpu    string
}

// Each builder should implement this interface 
type PCBuilder interface {
	SetRAM() PCBuilder
	SetHDD() PCBuilder
	SetCPU() PCBuilder
	GetPC() PersonalComputer
}

type HomePCBuilder struct {
	pc PersonalComputer
}

func (b *HomePCBuilder) SetRAM() PCBuilder {
	b.pc.ramCap = 4
	return b
}

func (b *HomePCBuilder) SetHDD() PCBuilder {
	b.pc.hddCap = 500
	return b
}

func (b *HomePCBuilder) SetCPU() PCBuilder {
	b.pc.cpu = "i3"
	return b
}

func (b *HomePCBuilder) GetPC() PersonalComputer {
	return b.pc
}

//Manufacturer object which aware of build process for builder type
type Manufacturer struct {
	b PCBuilder
}

func (m *Manufacturer) SetBuilder(builder PCBuilder) {
	m.b = builder
}

func (m *Manufacturer) ConstructPC() {
	m.b.SetCPU().SetHDD().SetRAM()
}