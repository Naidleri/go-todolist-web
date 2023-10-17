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

func Add(todolist entitites.List) bool {
	result, err := config.DB.Exec("INSERT INTO todolist (task,deadline) VALUE (?,?)", todolist.Task, todolist.Deadline)

	if err != nil {
		panic(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	return rowsAffected > 0
}

func Detail(id int) entitites.List {
	rows := config.DB.QueryRow("SELECT id,task from todolist WHERE id =?", id)

	var list entitites.List
	if err := rows.Scan(&list.Id, &list.Task); err != nil {
		panic(err.Error())
	}
	return list
}

func Update(id int, list entitites.List) bool {
	result, err := config.DB.Exec("UPDATE todolist SET task=? WHERE id=?", list.Task, list.Id)
	if err != nil {
		panic(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	return rowsAffected > 0
}

func Delete(id int) error {
	_, err := config.DB.Exec("DELETE FROM todolist WHERE id =?", id)
	return err
}
