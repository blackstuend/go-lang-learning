package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type person struct {
	FirstName string
	LastName  string
}

func main() {

	// connect to mgo
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	db := session.DB("user")
	c := db.C("loginInfo")
	err = c.Insert(&person{
		FirstName: "vincent",
		LastName:  "Lin",
	})

	if err != nil {
		panic(err)
	}

	var users []person
	err = c.Find(bson.M{}).All(&users)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", users)
}
