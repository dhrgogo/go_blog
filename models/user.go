package models

import (
	"fmt"
)

type User struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"created_by"`
	Email    string `json:"modified_by"`
}

// ExistTagByName checks if there is a tag with the same name
func ExistTagByName() (bool, error) {
	var user User
	//err := db.Find(&user, 2)
	err := db.Where("username = ?", "2222222").Find(&user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("----------------", user.Password)

	return false, nil
}
