package main

import (
	"go-todolist-web/config"
	"go-todolist-web/controller/completecontroller"
	listcontroller "go-todolist-web/controller/listcontoller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDb()

	//homepage
	http.HandleFunc("/", listcontroller.Index)
	http.HandleFunc("/complete", completecontroller.Home)

	//list
	http.HandleFunc("/add", listcontroller.Add)
	http.HandleFunc("/edit", listcontroller.Edit)
	http.HandleFunc("/delete", listcontroller.Delete)

	log.Print("Server berjalan (3000)")
	http.ListenAndServe(":3000", nil)
}
