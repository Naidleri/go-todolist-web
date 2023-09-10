package main

import (
	"go-todolist-web/config"
	"go-todolist-web/controller/listcontoller"
	"go-todolist-web/controller/homecontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDb()

	//homepage
	http.HandleFunc("/", homecontroller.Welcome)

	//categories
	http.HandleFunc("/categories", listcontroller.Index)
	http.HandleFunc("/categories/add", listcontroller.Add)
	http.HandleFunc("/categories/edit", listcontroller.Edit)
	http.HandleFunc("/categories/delete", listcontroller.Delete)

	log.Print("Server berjalan (8080)")
	http.ListenAndServe(":8080", nil)
}
