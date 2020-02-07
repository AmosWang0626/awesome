package main

import (
	"encoding/json"
	"fmt"
)

// 封装
func main() {
	user := SimpleUser("amos", "18937128861")
	fmt.Println(user)
	// json
	fmt.Println(&user)
}

type User struct {
	Id          int32  `json:"id"`
	Name        string `json:"name"`
	PhoneNo     string `json:"phone_no"`
	Age         int8   `json:"age"`
	Description string `json:"description"`
}

func (u User) Print() string {
	str, _ := json.Marshal(u)
	return string(str)
}

func (u User) SimpleUser(name, phoneNo string) User {
	return User{Name: name, PhoneNo: phoneNo}
}

func SimpleUser(name, phoneNo string) User {
	return User{Name: name, PhoneNo: phoneNo}
}

func (u *User) String() string {
	str, _ := json.Marshal(u)
	return string(str)
}
