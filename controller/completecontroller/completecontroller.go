package completecontroller

import (
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("view/complete/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w,nil)
}
