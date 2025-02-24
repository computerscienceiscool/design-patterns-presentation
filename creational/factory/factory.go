package factory

import "fmt"

type User interface {
	GetRole() string
}

type Admin struct{}

func (a Admin) GetRole() string { return "Admin" }

type Guest struct{}

func (g Guest) GetRole() string { return "Guest" }

func UserFactory(userType string) User {
	if userType == "Admin" {
		return Admin{}
	}
	return Guest{}
}

func DemoFactory() {
	user := UserFactory("Admin")
	fmt.Println(user.GetRole()) // Admin
}
