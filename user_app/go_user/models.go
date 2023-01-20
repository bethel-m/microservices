package main

import (
	"fmt"

	"gorm.io/gorm"
)

type user struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type userModel struct {
	db_obj *gorm.DB
}

func (m userModel) create_user(user_info user) {
	new_user_entry := user{FirstName: user_info.FirstName, LastName: user_info.LastName, Email: user_info.Email, Password: user_info.Password}

	fmt.Println("pass here")
	fmt.Println(new_user_entry)

	fmt.Println("here too")
	m.db_obj.Create(&new_user_entry)
	fmt.Println("reach the end")
}

func (m userModel) get_user(name string, db *gorm.DB) {
	var query user
	db.Where("first_name = ?", name).First(&query)
	fmt.Println(query)
}

func (m userModel) delete_user(email string, db *gorm.DB) {
	var query user
	fmt.Println("==========================")
	db.Unscoped().Where("email = ?", email).Delete(&query)
	fmt.Println("query", query)
}
