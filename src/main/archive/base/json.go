package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	fmt.Println("[user to json string] >>>>>>>>>>>>>")
	user := User{Id: int32(111), Name: "AMOS", PhoneNo: "18937128861", Age: 18, Description: "this is description"}
	userStr, _ := json.Marshal(user)
	fmt.Println("\tuser", user)
	fmt.Println("\t&user", &user)

	fmt.Println("..........................")

	fmt.Println("[json string to user] >>>>>>>>>>>>>")
	var user2 User
	_ = json.Unmarshal(userStr, &user2)
	fmt.Println("\tuser2", user2)
	fmt.Println("\t&user2", &user2)

}
