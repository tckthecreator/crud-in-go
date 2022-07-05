package models

import "github.com/tckthecreator/go-api-with-pgsql/db"

func GetAll() (todos []Todo, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	rows, err := conn.Query(`SELECT id, name, description, is_done FROM todos`)

	if err != nil {
		return
	}

	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.IsDone)

		if err != nil {
			continue
		}

		todos = append(todos, todo)
	}

	return
}
