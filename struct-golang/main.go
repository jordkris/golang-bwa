package main

import (
	"fmt"
	"struct-golang/management"
)

func displayUser(user management.User) string {
	return fmt.Sprintf("Name : %s %s, Email : %s", user.FirstName, user.LastName, user.Email)
}

func displayGroup(group management.Group) {
	fmt.Printf("Name : %s\n", group.Name)
	fmt.Printf("Member count : %d\n", len(group.Users))
	fmt.Printf("Member Name : \n")
	for _, user := range group.Users {
		fmt.Println(user.FirstName)
	}
}

func main() {
	user := management.User{}
	user.ID = 1
	user.FirstName = "Zelda"
	user.LastName = "Safitri"
	user.Email = "zelda@nintendo.com"
	user.IsActive = true

	user2 := management.User{
		ID:        2,
		FirstName: "Link",
		LastName:  "Awakening",
		Email:     "link@nintendo.com",
		IsActive:  true,
	}

	user3 := management.User{3, "Sonic", "The Hedgehog", "sonic@nintendo.com", true}
	fmt.Println(user)
	fmt.Println(user2)
	fmt.Println(user3)

	// struct as parameter
	displayUser := displayUser(user)
	fmt.Println(displayUser)

	// embedded struct
	users := []management.User{user, user2}
	group := management.Group{"Gamer", user, users, true}
	displayGroup(group)

	//method
	result := user.Display()
	fmt.Println(result)
	group.Display()
}
