package models

import "github.com/tckthecreator/go-api-with-pgsql/db"

func Insert(todo Todo) (id int64, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	query := `INSERT INTO todos (name, description, is_done) VALUES ($1, $2, $3) RETURNING id`

	err = conn.QueryRow(query, todo.Title, todo.Description, todo.IsDone).Scan(&id)

	return
}
