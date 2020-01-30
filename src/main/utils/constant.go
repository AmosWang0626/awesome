package utils

import "fmt"

var DefaultPassword string
var ProjectName string
var Version string

func init() {
	fmt.Println("constant.go init")
	DefaultPassword = "888888"
	ProjectName = "awesome"
	Version = "1.0.0"
}
