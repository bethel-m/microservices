package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func database_setup(database_name string) *gorm.DB {
	// open a connection to postgres and create a database
	dsn := "host=localhost user=postgres password=bravo2000 port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("error connecting to database >>>", err)
	}

	fmt.Println("pass here")
	create_db := fmt.Sprintf("CREATE DATABASE %s", database_name)
	db.Exec(create_db)

	// connect to created the created database and return the connection
	dsn = "host=localhost user=postgres password=bravo2000 dbname=users port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db_con, db_err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if db_err != nil {
		log.Println("error connecting to database >>>", err)
	}

	err = db_con.AutoMigrate(&user{})
	if err != nil {
		log.Println("error auto migrating >>", err)
	}

	return db_con
}
