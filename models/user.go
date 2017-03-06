package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//User is a struct for mongodb model
type User struct {
	ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name      string        `json:"name"`
	Email     string        `json:"email"`
	Password  string        `json:"password"`
	Slug      string        `json:"slug"`
	Bio       string        `json:"bio"`
	Avatar    []Avatar      `json:"avatar"`
	Account   Account       `json:"account"`
	Address   Address       `json:"addess"`
	CreatedAt time.Time     `json:"created_at"`
	UpdateAt  time.Time     `json:"updated_at"`
}

//Avatar is a struct for user struct
type Avatar struct {
	Small  string `json:"small"`
	Medium string `json:"medium"`
	Large  string `json:"large"`
}

//Account is a struct for user struct
type Account struct {
	Plan  string `json:"plan"`
	Level int    `json:"level"`
}

//Address is a struct for user struct
type Address struct {
	Country string `json:"country"`
	City    string `json:"city"`
}
