package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("[ERROR] ::", err)
		}
	}()

	resp, _ := http.Get("https://api.github.com/repos/AmosWang0626/notes")
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)

	body := string(bytes)
	fmt.Println("原始Response: ", body)

	var project *GithubProject
	_ = json.Unmarshal(bytes, &project)
	println(project.PushedAt)
}

type GithubProject struct {
	Name          string `json:"name"`
	FullName      string `json:"full_name"`
	Url           string `json:"url"`
	HtmlUrl       string `json:"html_url"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	PushedAt      string `json:"pushed_at"`
	DefaultBranch string `json:"default_branch"`
}
