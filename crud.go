package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Define a struct for the todo item
type TodoItem struct {
    ID        string `json:"id"`
    Title     string `json:"title"`
    Completed bool   `json:"completed"`
}

// Define a slice to hold the todo items
var todoList []TodoItem

// Get all the todo items
func getAllTodos(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(todoList)
}

// Get a single todo item by ID
func getTodoByID(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range todoList {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(nil)
}

// Add a new todo item
func addTodoItem(w http.ResponseWriter, r *http.Request) {
    var newTodo TodoItem
    json.NewDecoder(r.Body).Decode(&newTodo)
    todoList = append(todoList, newTodo)
    json.NewEncoder(w).Encode(newTodo)
}

// Update an existing todo item
func updateTodoItem(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for i, item := range todoList {
        if item.ID == params["id"] {
            var updatedTodo TodoItem
            json.NewDecoder(r.Body).Decode(&updatedTodo)
            updatedTodo.ID = item.ID
            todoList[i] = updatedTodo
            json.NewEncoder(w).Encode(updatedTodo)
            return
        }
    }
    json.NewEncoder(w).Encode(nil)
}

// Delete a todo item by ID
func deleteTodoItem(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for i, item := range todoList {
        if item.ID == params["id"] {
            todoList = append(todoList[:i], todoList[i+1:]...)
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(nil)
}

func main() {
    // Initialize the todo list
    todoList = []TodoItem{
        TodoItem{ID: "1", Title: "Todo item 1", Completed: false},
        TodoItem{ID: "2", Title: "Todo item 2", Completed: true},
        TodoItem{ID: "3", Title: "Todo item 3", Completed: false},
    }

    // Create a new router using the gorilla/mux package
    router := mux.NewRouter()

    // Define the API endpoints
    router.HandleFunc("/todos", getAllTodos).Methods("GET")
    router.HandleFunc("/todos/{id}", getTodoByID).Methods("GET")
    router.HandleFunc("/todos", addTodoItem).Methods("POST")
    router.HandleFunc("/todos/{id}", updateTodoItem).Methods("PUT")
    router.HandleFunc("/todos/{id}", deleteTodoItem).Methods("DELETE")

    // Start the HTTP server
	fmt.Println("server started")
    log.Fatal(http.ListenAndServe(":8080", router))
}
