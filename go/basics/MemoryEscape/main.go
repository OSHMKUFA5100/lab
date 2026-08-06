package main

import (
	"fmt"
)
type User struct {
	Name string
}

func NewUser() *User {
	u := User{}
	return &u
}

func ClosureFunc() func() int {
    a := 1
    return func() int {
        return a
    }
}
func main() {
	fmt.Println(NewUser())
	cf:= ClosureFunc()
	fmt.Println(cf())
}
