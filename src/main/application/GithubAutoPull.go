package main

import (
	"amos.wang/awesome/src/main/utils/date_utils"
	"amos.wang/awesome/src/main/utils/exec_utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	// 缓存最后一次提交时间
	var lastPushedAt string

	var ch chan int

	ticker := time.NewTicker(time.Second * 10)
	go func() {
		for range ticker.C {
			repos, err := getPushedAt()
			if err != nil {
				fmt.Println("[ERROR] ::", err)
				continue
			}

			currentPushedAt := repos.PushedAt
			if lastPushedAt == currentPushedAt {
				fmt.Println(date_utils.NowStr(), "无需更新")
				continue
			}

			exec_utils.ExecCmd("xxx", "git pull")
			lastPushedAt = currentPushedAt
			fmt.Println(date_utils.NowStr(), "更新项目, 最后一次提交时间", lastPushedAt)
		}
		ch <- 1
	}()
	<-ch

}

/*
获取项目信息
*/
func getPushedAt() (repos *GithubRepos, err error) {
	resp, err := http.Get("https://api.github.com/repos/AmosWang0626/notes")
	if err != nil {
		return
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(bytes, &repos)
	fmt.Println("RES", repos)

	return
}

type GithubRepos struct {
	FullName      string `json:"full_name"`
	DefaultBranch string `json:"default_branch"`
	PushedAt      string `json:"pushed_at"`
	UpdatedAt     string `json:"updated_at"`
	HtmlUrl       string `json:"html_url"`
}
