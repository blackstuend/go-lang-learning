package main

import (
	"leaning-test-httpd/pkg/handler"

	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
	"gopkg.in/mgo.v2"
)

func main() {
	// get init.config
	config, err := ini.Load("conf/app.ini")
	if err != nil {
		panic("")
	}

	// get server port
	port := config.Section("runtime").Key("port").MustString(":5000")

	// connect to mgo and get collection
	session, _ := mgo.Dial("127.0.0.1:27017")
	db := session.DB("todo")
	c := db.C("login.info")

	h := &handler.Handler{
		DB: c,
	}

	r := gin.Default()
	r.LoadHTMLGlob("./templates/*.html")

	r.GET("/", h.GetIndex)
	r.POST("/addToDo", h.AddTodo)
	r.DELETE("/removeTodo/:id", h.RemoveTodo)
	r.PUT("updateTodo/:id", h.UpdateCompleted)
	r.Run(port)
}
