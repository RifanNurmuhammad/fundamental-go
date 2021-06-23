package main

import (
	"fmt"
	"sort"
)

// User struct
type User struct {
	Name string
	Age  int
}

// UserSlice type
type UserSlice []User

func (user UserSlice) Len() int {
	return len(user)
}

func (user UserSlice) Less(i, j int) bool {
	return user[i].Age < user[j].Age
}

func (user UserSlice) Swap(i, j int) {
	user[i], user[j] = user[j], user[i]
}

func main() {
	users := []User{
		{"Rifan", 30},
		{"Nur", 35},
		{"Muhammad", 31},
		{"Mahveen", 25},
	}

	fmt.Println(users)
	sort.Sort(UserSlice(users))

	fmt.Println(users)
}
