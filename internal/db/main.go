package db

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var Database = DB{
	Todos: []*Todo{},
}

type DB struct {
	Todos []*Todo
}

type Todo struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	IsDone      bool      `json:"isDone"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewTodo(dsc string) (Todo, error) {
	// validation
	if dsc == "" {
		return Todo{}, errors.New("invalid input")
	}

	return Todo{
		ID:          uuid.New(),
		Description: dsc,
		IsDone:      false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func (db *DB) Add(todo Todo) {
	db.Todos = append(db.Todos, &todo)
}

func (db *DB) Get(id string) (*Todo, error) {
	for _, todo := range db.Todos {
		if todo.ID.String() == id {
			return todo, nil
		}
	}

	return nil, errors.New("invalid id")
}

func (db *DB) Done(id string) error {
	todo, err := db.Get(id)
	if err != nil {
		return err
	}

	todo.IsDone = true
	todo.UpdatedAt = time.Now()

	return nil
}

func (db *DB) Delete(id string) {
	nextTodos := []*Todo{}

	for _, todo := range db.Todos {
		if todo.ID.String() != id {
			nextTodos = append(nextTodos, todo)
		}
	}

	db.Todos = nextTodos
}
