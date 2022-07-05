package models

import "github.com/tckthecreator/go-api-with-pgsql/db"

func Get(id int64) (todo Todo, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	row := conn.QueryRow(`SELECT id, name, description, is_done FROM todos WHERE id = $1`, id)

	err = row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.IsDone)

	return
}
