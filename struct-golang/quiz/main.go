package main

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvailable bool
}

func displayUser(user User) string {
	return fmt.Sprintf("Name : %s %s, Email : %s", user.FirstName, user.LastName, user.Email)
}

func displayGroup(group Group) {
	fmt.Printf("Name : %s\n", group.Name)
	fmt.Printf("Member count : %d\n", len(group.Users))
	fmt.Printf("Member Name : \n")
	for _, user := range group.Users {
		fmt.Println(user.FirstName)
	}
}

func (user User) display() string {
	return fmt.Sprintf("Name : %s %s, Email : %s", user.FirstName, user.LastName, user.Email)
}

func (group Group) display() {
	fmt.Printf("Name : %s\n", group.Name)
	fmt.Printf("Member count : %d\n", len(group.Users))
	fmt.Printf("Member Name : \n")
	for _, user := range group.Users {
		fmt.Println(user.FirstName)
	}
}

func main() {
	user := User{}
	user.ID = 1
	user.FirstName = "Zelda"
	user.LastName = "Safitri"
	user.Email = "zelda@nintendo.com"
	user.IsActive = true

	user2 := User{
		ID:        2,
		FirstName: "Link",
		LastName:  "Awakening",
		Email:     "link@nintendo.com",
		IsActive:  true,
	}

	user3 := User{3, "Sonic", "The Hedgehog", "sonic@nintendo.com", true}
	fmt.Println(user)
	fmt.Println(user2)
	fmt.Println(user3)

	// struct as parameter
	displayUser := displayUser(user)
	fmt.Println(displayUser)

	// embedded struct
	users := []User{user, user2}
	group := Group{"Gamer", user, users, true}
	displayGroup(group)

	//method
	result := user.display()
	fmt.Println(result)
	group.display()
}
