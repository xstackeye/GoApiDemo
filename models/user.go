package models

import (
	// "errors"
	"log"
)

//User
type User struct {
	Id   int64  `json:"id"`
	Name string `xorm:"unique not null" json:"name"`
	Type string `json:"type" xorm:"not null default user"`
}

func init() {
	// syncdb
	if err := x.Sync(new(User)); err != nil {
		log.Fatalf("Fail to sync database: %v\n", err)
	}
}

func NewUser(name string) (User, error) {
	var user User
	_, err := x.Insert(&User{Name: name, Type: "user"})
	if err != nil {
		log.Printf("Fail to create user: %v\n", err)
		return user, err
	}
	// get User just added
	user.Name = name
	_, err = x.Get(&user)
	return user, err
}

func GetUser() ([]User, error) {
	var user []User = make([]User, 0)
	err := x.Find(&user)
	if err != nil {
		log.Printf("Fail to get user from database: %v\n", err)
		return nil, err
	}
	return user, nil
}
