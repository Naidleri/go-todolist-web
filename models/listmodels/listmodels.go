package listmodels

import (
	"go-todolist-web/config"
	"go-todolist-web/entitites"
)

func GetAll() []entitites.List {
	rows, err := config.DB.Query("select * from todolist")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var categories []entitites.List

	for rows.Next() {
		var list entitites.List
		if err := rows.Scan(&list.Id, &list.Task, &list.Deadline, &list.Completed); err != nil {
			panic(err)
		}

		categories = append(categories, list)
	}
	return categories

}

func Create(todolist entitites.List) bool {
	result, err := config.DB.Exec("INSERT INTO todolist (task,deadline,completed) VALUE (?,?,?)", todolist.Task, todolist.Deadline, todolist.Completed)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastInsertId > 0
}

func Delete(id int) error {
	_, err := config.DB.Exec("DELETE FROM todolist WHERE id =?", id)
	return err
}
