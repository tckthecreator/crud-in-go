package models

import "github.com/tckthecreator/go-api-with-pgsql/db"

func Update(id int64, todo Todo) (int64, error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return 0, err
	}

	defer conn.Close()

	res, err := conn.Exec(`UPDATE todos SET name = $1, description = $2, is_done = $3 WHERE id = $4`, todo.Title, todo.Description, todo.IsDone, id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
