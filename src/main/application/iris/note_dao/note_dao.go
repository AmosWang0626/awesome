package note_dao

import (
	"amos.wang/awesome/src/main/application/im/common/imerror"
	"amos.wang/awesome/src/main/application/im/common/utils"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"time"
)

// Note 定义
type Note struct {
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

func (m *Note) Decode(buf []byte) *Note {
	_ = json.Unmarshal(buf, m)
	return m
}

func (m *Note) Encode() []byte {
	buf, _ := json.Marshal(m)
	return buf
}

// Iris Note RedisDao
var (
	IrisNoteRedisDao *NoteRedisDao
	NoteKeyPrefix    = "IRIS_NOTE_"
)

func InitIrisNoteRedisDao() {
	IrisNoteRedisDao = &NoteRedisDao{Pool: utils.ImRedisPool()}
}

type NoteRedisDao struct {
	Pool *redis.Pool
}

func (current *NoteRedisDao) Save(note *Note) (err error) {
	_, err = current.Pool.Get().Do("set", NoteKeyPrefix+note.Title, string(note.Encode()))
	return
}

func (current *NoteRedisDao) Select(title string) (note *Note, err error) {
	result, err := redis.String(current.Pool.Get().Do("get", NoteKeyPrefix+title))
	if err != nil {
		if err == redis.ErrNil {
			err = imerror.UserUndefined
		}
		return
	}
	note = &Note{}
	note = note.Decode([]byte(result))
	return
}
