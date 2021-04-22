package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Handler struct {
	// mgo client
	DB *mgo.Collection
}

type ToDo struct {
	Content  string
	Complete bool
	Id       string
}

func (h *Handler) GetIndex(c *gin.Context) {
	// get Todo
	var data []ToDo
	h.DB.Find(bson.M{}).All(&data)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"TodoItems": data,
	})
}

func (h *Handler) AddTodo(c *gin.Context) {

	var data ToDo

	if err := c.Bind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "404",
			"content": "fail",
		})
		return
	}

	// save to db
	h.DB.Insert(&ToDo{
		Content:  data.Content,
		Complete: false,
	})

	// return json
	c.JSON(200, gin.H{
		"status":  "200",
		"content": data.Content,
	})
}

func (h *Handler) DeleteTodo(c *gin.Context) {

}
