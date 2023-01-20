package main

import (
	"log"
	"net/http"
)

type fst struct {
	fs string
}

func main() {
	db := database_setup("users")
	h := BaseHandler{
		usermodel: userModel{db_obj: db},
	}
	loadRoutes(&h)

	log.Println("starting web server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println("error starting server: ", err)
	}

}
