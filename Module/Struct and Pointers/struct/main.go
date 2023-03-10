package main

import (
	"encoding/json"
	"fmt"
	"regexp"
)

type Person struct {
	ID       string
	Name     string
	Phone    string
	Email    string
	Password string
}

func (p *Person) Register() {
	fmt.Println("Registering person")
	fmt.Print("ENTER ID: ")
	fmt.Scanln(&p.ID)
	fmt.Print("ENTER NAME: ")
	fmt.Scanln(&p.Name)
	fmt.Print("ENTER PHONE: ")
	fmt.Scanln(&p.Phone)
	fmt.Print("ENTER EMAIL: ")
	// check if email is correct format email like @gmail.
	fmt.Scanln(&p.Email)

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	isValidEmail := emailRegex.MatchString(p.Email)

	if !isValidEmail {
		fmt.Println("Invalid Email Format")
	}

	// print the person as json
	response, err := json.Marshal(p)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(response))

}

func (p *Person) Login() {
	fmt.Print("Enter Email:")
	fmt.Scanln(&p.Email)
	if p.Email != "faiz@gmail.com" {
		fmt.Println("Error Login")
	} else {
		fmt.Println("Success")
	}

}

func main() {
	var p Person
	p.Login()
}
