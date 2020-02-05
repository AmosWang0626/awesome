package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	var user string
	var pwd string
	var host string
	var port string

	flag.StringVar(&user, "u", "root", "user name, default null")
	flag.StringVar(&pwd, "pwd", "root", "user password, default null")
	flag.StringVar(&host, "h", "localhost", "shell host, default null")
	flag.StringVar(&port, "p", "3306", "shell port, default null")
	flag.Parse()

	// go run main_args.go -u amos -pwd 1433233 -h 172.0.0.1 -p 8080
	fmt.Printf("user=%v, password=%v, host=%v, port=%v\n", user, pwd, host, port)

	fmt.Println("len(os.Args)", len(os.Args))
	for index, value := range os.Args {
		fmt.Println("\t", index, ":", value)
	}

}
