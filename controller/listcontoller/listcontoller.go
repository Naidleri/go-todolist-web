package listcontroller

import (
	"go-todolist-web/entitites"
	"go-todolist-web/models/listmodels"
	"html/template"
	"net/http"
	"strconv"
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

		// Tambahkan kode untuk memeriksa apakah todolist.Task kosong
		if todolist.Task == "" {
			http.Error(w, "Task tidak boleh kosong", http.StatusBadRequest)
			return
		}

		// Tambahkan kode untuk memeriksa apakah todolist.Deadline valid
		if !todolist.Deadline.After(time.Now()) {
			http.Error(w, "Deadline harus di masa depan", http.StatusBadRequest)
			return
		}

		// Tambahkan kode untuk set status todolist menjadi false
		todolist.Completed = false

		// Tambahkan kode untuk memanggil fungsi Create()
		if !listmodels.Create(todolist) {
			http.Error(w, "Gagal menambahkan todolist", http.StatusInternalServerError)
			return
		}

		// Tambahkan kode untuk redirect ke halaman("/")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("view/list/edit.html")
		if err != nil {
			panic(err)
		}

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		todolist := listmodels.Detail(id)
		data := map[string]any{
			"todolist": todolist,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var todolist entitites.List
		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		todolist.Task = r.FormValue("edit")

		if !listmodels.Update(id, todolist) {
			http.Error(w, "Gagal mengedit todolist", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := listmodels.Delete(id); err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
