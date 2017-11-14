package main

import "fmt"

func main() {
	u := user{
		Age:  19,
		Name: "Inveja",
		Live: true,
	}
	fmt.Println(u)
	u.Morrer()

	var a age
	a = 19
	a.Envelhecer()
	fmt.Println(a)
}

type user struct {
	Name string
	Age  int
	Live bool
}

type age int

func (a *age) Envelhecer() {
	*a += 1
}

func (u *user) Morrer() {
	fmt.Println("Faliceu")
	u.Live = false
}
