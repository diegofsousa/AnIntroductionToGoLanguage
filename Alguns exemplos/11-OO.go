package main

import "fmt"

func main() {
	var adm Admin
	var su simpleUser

	adm.Login()

	fmt.Println(adm.online)

	adm.Logout()

	fmt.Println(adm.online)

	su.Message()

}

type User struct {
	name   string
	online bool
}

func (u *User) Login() {
	u.online = true
}

func (u *User) Logout() {
	u.online = false
}

type Admin struct {
	User    // estendendo usuários
	ranking int
}

type simpleUser User

func (s simpleUser) Message() {
	fmt.Println("Olá")
}
