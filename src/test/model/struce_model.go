package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name  string   `json:"name"`
	Age   int      `json:"age"`
	Skill []string `json:"skill"`
}

var fileName = "E:\\project\\golang\\awesome\\src\\main\\res\\text.txt"

func (m *Monster) Format() (string, bool) {
	bytes, err := json.Marshal(m)
	fmt.Printf("\tFormat &m point address is %p\n", m)

	err = ioutil.WriteFile(fileName, bytes, 0777)

	return string(bytes), err == nil
}

func (m *Monster) Parse() (*Monster, bool) {
	var m2 Monster

	bytes, err := ioutil.ReadFile(fileName)
	err = json.Unmarshal(bytes, &m2)
	fmt.Printf("\tParse &m2 point address is %p\n", &m2)

	return &m2, err == nil
}
