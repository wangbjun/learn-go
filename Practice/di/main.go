package main

import (
	"github.com/facebookarchive/inject"
	"log"
)

type Logger struct {
	cfg string
}

func (Logger) log() {
	log.Printf("write log")
}

type App struct {
	Logger Logger `inject:""` //依赖logger
}

func (r App) run() {
	r.Logger.log()
}

func main() {
	//传统写法
	app := App{
		Logger: Logger{},
	}
	app.run()

	//依赖注入写法
	app2 := App{}
	var g inject.Graph
	_ = g.Provide(&inject.Object{Value: &app2}, &inject.Object{Value: &Logger{}})
	_ = g.Populate()
	app2.run()
}
