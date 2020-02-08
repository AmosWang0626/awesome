package codec

import (
	"amos.wang/awesome/src/main/utils/log_utils"
)

type Codec interface {
	Decode()
	Encode()
}

func Decode() {
	defer func() {
		if err := recover(); err != nil {
			log_utils.Error.Println("codec decode", err)
		}
	}()
}

func Encode() {
	defer func() {
		if err := recover(); err != nil {
			log_utils.Error.Println("codec encode", err)
		}
	}()
}
