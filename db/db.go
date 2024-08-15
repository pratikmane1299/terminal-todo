package db

import (
	"database/sql"
	"fmt"
	"path/filepath"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	Id        int
	Todo      string
	Completed bool
	Created   time.Time
}

type DB struct {
	Db      *sql.DB
	dataDir string
}

func (db *DB) tableExists() bool {
	if _, err := db.Db.Query("select  * from todos"); err != nil {
		return false
	}

	return true
}

func (db *DB) createTable() error {
	_, err := db.Db.Exec(`CREATE TABLE "todos" ( "id" INTEGER, "todo" TEXT NOT NULL, "completed" TEXT, "created" DATETIME, PRIMARY KEY("id" AUTOINCREMENT) )`)

	return err
}

func (db *DB) GetAllTodos() ([]Todo, error) {
	var todos []Todo
	rows, err := db.Db.Query("Select * from todos")
	if err != nil {
		return todos, err
	}

	for rows.Next() {
		var todo Todo
		err := rows.Scan(
			&todo.Id,
			&todo.Todo,
			&todo.Completed,
			&todo.Created,
		)

		if err != nil {
			return todos, err
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

func (db *DB) GetTaskById(id int) (Todo, error) {
	var todo Todo
	err := db.Db.QueryRow("select * from todos where id = ?", id).Scan(&todo.Id, &todo.Todo, &todo.Completed, &todo.Created)

	return todo, err

}

func (db *DB) DeleteTodoById(todoId int) error {
	_, err := db.GetTaskById(todoId)
	if err != nil {
		return fmt.Errorf("todo with id %d not found %w", todoId, err)
	}
	_, err = db.Db.Exec("delete from todos where id = ?", todoId)

	return err
}

func (db *DB) AddTodo(todo string) error {
	_, err := db.Db.Exec("insert into todos(todo, completed, created) values(?, ?, ?)", todo, false, time.Now())

	return err
}

func (db *DB) UpdateTodo(todo Todo) error {
	_, err := db.GetTaskById(todo.Id)
	if err != nil {
		return fmt.Errorf("todo with id %d not found %w", todo.Id, err)
	}

	_, err = db.Db.Exec("update todos set todo = ?, completed = ? where id = ?", todo.Todo, todo.Completed, todo.Id)

	return err
}

func OpenDB(path string) (*DB, error) {
	db, err := sql.Open("sqlite3", filepath.Join(path, "terminal.todo.db"))
	if err != nil {
		return nil, err
	}

	terminalDb := DB{Db: db}

	exists := terminalDb.tableExists()
	if !exists {
		err := terminalDb.createTable()
		if err != nil {
			return nil, err
		}
	}

	return &terminalDb, nil
}
