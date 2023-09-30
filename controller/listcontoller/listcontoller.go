package listcontroller

import (
	"go-todolist-web/entitites"
	"go-todolist-web/models/listmodels"
	"html/template"
	"net/http"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	categories := listmodels.GetAll()
	data := map[string]any{
		"list": categories,
	}

	temp, err := template.ParseFiles("view/list/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("view/list/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var todolist entitites.List

		todolist.Task = r.FormValue("todolist")
		deadlineStr := r.FormValue("deadline")

		deadline, err := time.Parse("2006-01-02", deadlineStr)
		if err != nil {
			http.Error(w, "Gagal mengambil tangggal", http.StatusBadRequest)
			return
		}
		todolist.Deadline = deadline

		if time.Now().After(deadline) {
			todolist.Completed = false
		} else {
			todolist.Completed = true
		}
		listmodels.Create(todolist)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {

}

func Delete(w http.ResponseWriter, r *http.Request) {

}
