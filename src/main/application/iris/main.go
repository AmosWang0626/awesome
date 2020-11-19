package main

import (
	"amos.wang/awesome/src/main/application/im/common/utils"
	"amos.wang/awesome/src/main/application/iris/note_dao"
	"amos.wang/awesome/src/main/utils/log_utils"
	"fmt"
	"github.com/kataras/iris/v12"
	"time"
)

func main() {
	app := iris.Default()

	pool := utils.ImRedisPool()
	conn := pool.Get()
	note_dao.InitIrisNoteRedisDao()
	defer conn.Close()
	defer pool.Close()

	bytes := 0
	app.Get("/note/{title}", func(ctx iris.Context) {
		title := ctx.Params().Get("title")
		note, err := note_dao.IrisNoteRedisDao.Select(title)
		if err != nil {
			log_utils.Error.Println(err)
		}

		bytes, _ = ctx.Writef(note.Content)
		fmt.Println("written bytes :: ", bytes)
	})

	app.Get("/note/{title}/{content:path}", func(ctx iris.Context) {
		title := ctx.Params().Get("title")
		content := ctx.Params().Get("content")

		note := &note_dao.Note{Title: title, Content: content, CreateTime: time.Now()}
		err := note_dao.IrisNoteRedisDao.Save(note)
		if err != nil {
			log_utils.Error.Println(err)
		}

		message := "[" + title + "]" + " Content :: " + content
		bytes, _ = ctx.WriteString(message)
		fmt.Println("written bytes :: ", bytes)
	})

	_ = app.Listen(":8080")
}
