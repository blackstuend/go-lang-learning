package handler

import (
	"encoding/json"
	"fmt"
	"log"
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
	Content  string        `json:"content"`
	Complete bool          `json:"complete"`
	Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
}

type ToDOView struct {
	ID       string `json:"id"`
	Content  string `json:"content"`
	Complete bool   `json:"complete"`
}

func (h *Handler) GetIndex(c *gin.Context) {
	// get Todo
	var (
		data     []ToDo
		viewData []ToDOView
	)

	h.DB.Find(bson.M{}).All(&data)
	fmt.Printf("%+v", data)
	buf, _ := json.Marshal(data)

	_ = json.Unmarshal(buf, &viewData)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"TodoItems": viewData,
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

func (h *Handler) RemoveTodo(c *gin.Context) {

	id := c.Param("id")

	// delete data
	if err := h.DB.Remove(bson.M{"_id": bson.ObjectIdHex(id)}); err != nil {
		// delete fail
		log.Printf("failed remove todo reason : %v \n", err)

		c.JSON(200, gin.H{
			"status":  "400",
			"content": "delete fail",
		})
		return
	}
	// completed delete
	c.JSON(200, gin.H{
		"status":  "200",
		"content": "delete success",
	})
}

func (h *Handler) UpdateCompleted(c *gin.Context) {

	id := c.Param("id")

	var data ToDo

	if err := c.Bind(&data); err != nil {
		c.JSON(200, gin.H{
			"status":  "404",
			"content": "bind failed",
		})
		return
	}
	h.DB.FindId(bson.ObjectIdHex(id)).One(&data)

	data.Complete = true

	// update data
	if err := h.DB.UpdateId(bson.ObjectIdHex(id), data); err != nil {
		c.JSON(200, gin.H{
			"status":  "404",
			"content": "update fail",
		})
		return
	}

	// completed delete
	c.JSON(200, gin.H{
		"status":  "200",
		"content": "update success",
	})
}
