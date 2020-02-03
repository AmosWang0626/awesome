package test

import (
	"amos.wang/awesome/src/test/model"
	"encoding/json"
	"testing"
)

func TestModel(t *testing.T) {
	m := &model.Monster{Name: "amos", Age: 18, Skill: []string{"111", "222"}}
	t.Logf("\tTestModel &m point address is %p\n", m)
	content, save := m.Format()
	t.Logf("save content to file, save=%t, content is [%s]\n", save, content)

	m2, read := m.Parse()
	t.Logf("\tTestModel &m2 point address is %p\n", m2)

	bytes, _ := json.Marshal(m2)
	t.Logf("read content to file, read=%t, content is [%s]\n", read, string(bytes))
}
