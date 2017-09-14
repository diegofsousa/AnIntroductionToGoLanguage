package main

import "fmt"

func main() {
	u := User{
		"Diego",
		"diegofsousa",
		true,
	}

	adm := Admin{
		User{
			"Diogo",
			"fernando",
			false,
		},
		21,
	}
	ShowUserInfo(u)
	ShowUserInfo(adm)
}

func ShowUserInfo(u UsersInterface) {
	fmt.Println(u.Show())
}

// Saudosa interface
type UsersInterface interface {
	Show() string
}

type User struct {
	name     string
	username string
	online   bool
}

func (u User) Show() string {
	return fmt.Sprintf("Hello, my name in %s, and my username in %s.", u.name, u.username)
}

func (a Admin) Show() string {
	return fmt.Sprintf("Hello, my name in %s, my username in %s and I have %d years old.", a.name, a.username, a.age)
}

type Admin struct {
	User
	age int
}
