package main

import "fmt"


type BadUser struct {
	name string
	age int
}

func (u *BadUser) GetName() string { // not idiomatic in go
	return u.name
}

func (u *BadUser) SetName(name string) { // no validation, no logic -> useless
	u.name = name
}

func (u *BadUser) GetAge() int { // also unnecessary
	return u.age
}
func (u *BadUser) SetAge(age int) { // no validation, no logic -> useless
	u.age = age
}

func badExample() {
	u := &BadUser{}
	u.SetName("Tanishq")
	u.SetAge(20)

	fmt.Println("Name:" , u.GetName())
	fmt.Println("age:", u.GetAge())
}

type User struct {
	Name string 
	Age  int
}

func goodExample() {
	u := User{Name: "Tanishq", Age: 20}
	u.Age++ // simple, readable, idiomatic

	fmt.Println("Name:", u.Name)
	fmt.Println("Age:", u.Age)
}

func main() {
	badExample()
	goodExample()
}

