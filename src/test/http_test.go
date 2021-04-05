package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestPredictIntent(t *testing.T) {
	question := "电子银行服务通知"
	get, err := http.Get("http://172.16.6.66:8088/mininlp/intent/predictIntent/" + question)
	if err != nil {
		fmt.Println("ERROR Request /mininlp/intent/predictIntent/", question, "  :: ", err)
		return
	}

	body, err := ioutil.ReadAll(get.Body)
	fmt.Println("SUCCESS Response :: ", string(body))

	var response *GeneralResponse
	_ = json.Unmarshal(body, &response)
	fmt.Println("Response Data is :: ", response.Data)
	fmt.Println("Probability :: ", response.Data.Probability, ", Content :: ", response.Data.Content)
}

type GeneralResponse struct {
	Code        string          `json:"code"`
	Data        IntentPredictVO `json:"data"`
	Explanation string          `json:"explanation"`
	Meaning     string          `json:"meaning"`
}

type IntentPredictVO struct {
	Name        string  `json:"name"`
	Probability float64 `json:"probability"`
	Content     string  `json:"content"`
}
