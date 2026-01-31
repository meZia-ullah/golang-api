package routes

import (
	"net/http"
	"strings"

	"goapi/handlers"
)

func RegisterRoutes() {
	// Staff routes
	http.HandleFunc("/staff", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlers.GetStaff(w, r)
			return
		}
		if r.Method == http.MethodPost {
			handlers.CreateStaff(w, r)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
	})

	// Auth routes
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handlers.Register(w, r)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handlers.Login(w, r)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
	})

	// Todo routes
	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlers.GetTodos(w, r)
			return
		}
		if r.Method == http.MethodPost {
			handlers.CreateTodo(w, r)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
	})

	// Todo detail routes (PUT, DELETE)
	http.HandleFunc("/todo/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/todo/") {
			if r.Method == http.MethodPut {
				handlers.UpdateTodo(w, r)
				return
			}
			if r.Method == http.MethodDelete {
				handlers.DeleteTodo(w, r)
				return
			}
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
	})
}
