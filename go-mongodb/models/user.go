package models

import "gopkg.in/mgo.v2/bson"

type User struct{
	Id
	Name
	Gender
	Age
}