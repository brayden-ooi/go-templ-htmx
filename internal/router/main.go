package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/brayden-ooi/go-templ-htmx/internal/db"
	"github.com/brayden-ooi/go-templ-htmx/internal/front_end"
)

func Exec() {
	http.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("web/static/"))))

	http.Handle("/", templ.Handler(front_end.Root()))

	http.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			Respond(w, r)
		case http.MethodPost:
			todo, err := db.NewTodo(r.FormValue("todo"))
			if err != nil {
				_, err := w.Write([]byte("error adding new item to the list"))
				if err != nil {
					fmt.Println("error returning err: ", err)
				}
				return
			}

			db.Database.Add(todo)

			Respond(w, r)
		case http.MethodPut:
			id := r.URL.Query().Get("id")

			err := db.Database.Done(id)
			if err != nil {
				_, err := w.Write([]byte("error updating item to the list"))
				if err != nil {
					fmt.Println("error returning err: ", err)
				}
				return
			}

			Respond(w, r)
		case http.MethodDelete:
			id := r.URL.Query().Get("id")

			db.Database.Delete(id)

			Respond(w, r)
		}
	})

	http.Handle("/404", templ.Handler(front_end.NotFound(), templ.WithStatus(http.StatusNotFound)))
}

func Respond(w http.ResponseWriter, r *http.Request) {
	switch r.Header.Get("Accept") {
	case "application/json":
		err := json.NewEncoder(w).Encode(db.Database.Todos)
		if err != nil {
			_, err := w.Write([]byte("error returning json payload"))
			if err != nil {
				fmt.Println("error returning err: ", err)
			}
			return
		}
	case "text/html":
		fallthrough
	default:
		todoContainer := front_end.TodoContainer(db.Database.Todos)
		err := todoContainer.Render(r.Context(), w)
		if err != nil {
			_, err := w.Write([]byte("error returning html payload"))
			if err != nil {
				fmt.Println("error returning err: ", err)
			}
			return
		}
	}
}
