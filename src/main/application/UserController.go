package main

import (
	"amos.wang/awesome/src/main/application/entity"
	"fmt"
)

func main() {
	user := entity.SimpleUser("amos", "18937128861")
	fmt.Println(user)
	// json
	fmt.Println(&user)
}
