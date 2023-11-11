package model

import "errors"

type Todo = struct {
	Id          int64
	Description string
	IsDone      int
}

func CreateTodo(description string, is_done bool) (Todo, error) {
	res, err := db.Exec("INSERT INTO Todo (description,is_done) VALUES (?,?);", description, is_done)
	if err != nil {
		return Todo{}, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return Todo{}, err
	}
	var done int = 0
	if is_done {
		done = 1
	}
	return Todo{Id: id, Description: description, IsDone: done}, nil
}

func GetAllTodos(limit int) ([]Todo, error) {

	rows, err := db.Query("SELECT * FROM Todo LIMIT (?);", limit)

	if err != nil {
		return nil, err
	}
	var todos []Todo
	defer rows.Close()
	for rows.Next() {
		var todo Todo
		rows.Scan(&todo.Id, &todo.Description, &todo.IsDone)
		todos = append(todos, todo)
	}

	return todos, nil
}

func ToggleTodoDone(id int64) error {
	todo, err := GetTodo(id)
	if err != nil {
		return err
	}
	tx, err := db.Begin()

	if err != nil {
		return err
	}

	var done int = 0
	if todo.IsDone > 0 {
		done = 0
	} else {
		done = 1
	}

	_, err = tx.Exec("UPDATE Todo SET is_done = ? WHERE id = ?;", done, id)
	if err != nil {
		tx.Rollback()
		return err

	}

	tx.Commit()

	return err
}

func UpdateTodo(id int64, description string, is_done bool) error {
	todo, err := GetTodo(id)
	if err != nil {
		return err
	}
	tx, err := db.Begin()

	if err != nil {
		return err
	}

	var done int = 0
	if todo.IsDone > 0 {
		done = 0
	} else {
		done = 1
	}

	_, err = tx.Exec("UPDATE Todo SET is_done = ?, description = ? WHERE id = ?;", done, description, id)
	if err != nil {
		tx.Rollback()
		return err

	}

	tx.Commit()

	return err
}

func DeleteTodo(id int64) error {

	_, err := GetTodo(id)
	if err != nil {
		return err
	}
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("DELETE FROM Todo WHERE id=(?);", id)

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func GetTodo(id int64) (Todo, error) {

	rows, err := db.Query("SELECT * FROM Todo WHERE id=(?);", id)

	if err != nil {
		return Todo{}, err
	}
	var todo Todo = Todo{Id: -1}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&todo.Id, &todo.Description, &todo.IsDone)
		break
	}
	if todo.Id < 1 {
		return todo, errors.New("not found")
	}
	return todo, nil
}
