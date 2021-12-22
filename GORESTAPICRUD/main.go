package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type task struct {

	ID int `json:"id"`
	Name string `json:"name"`
	Content string`json:"content"`

}
type allTasks []task


var  tasks = allTasks{
	{
		ID :1,
		Name:"Tarea 1",
		Content: "Some content",
	},
}


func getTasks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)

}

func createTasks(w http.ResponseWriter, r *http.Request){
	var newTask task
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a valid Task")
	}

	json.Unmarshal(reqBody, &newTask)

	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)

}

func deleteTask (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	//conversor  a entero
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil{
		fmt.Fprintf(w, "Invalid ID")
		return
	}
	for index, task := range tasks{
		if task.ID == taskID{
			tasks = append(tasks[:index], tasks[index + 1:]...)
			fmt.Fprintf(w, "The task with Id %v has been remove succesfully",taskID)
			w.Header().Set("Content-Type", "application/json")

		}
	}

}

func getTask(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	//conversor  a entero
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil{
		fmt.Fprintf(w, "Invalid ID")
		return
	}
	for _, task := range tasks{
		if task.ID == taskID{
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
		}
	}

}

func updateTask(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	//conversor  a entero
	taskID, err := strconv.Atoi(vars["id"])
	var updatedTask task

	if err != nil{
		fmt.Fprintf(w, "Invalid ID")
		return
	}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Please Enter Valid Data")

	}
	json.Unmarshal(reqBody, &updatedTask)

	for index, task := range tasks {
		if task.ID == taskID{
			tasks = append(tasks[:index], tasks[index + 1:]...)
			updatedTask.ID = taskID
			tasks = append(tasks, updatedTask)

			fmt.Fprintf(w, "The task with ID %v has been update succesfully",taskID)
		}
	}

}


func indexRoute (w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Welcome to my API"))

}

func main()  {

	//Initialize the server
	router := mux.NewRouter().StrictSlash(true)

	//Routes
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks", createTasks).Methods("POST")
	router.HandleFunc("/tasks/{id}", getTask).Methods("GET")
	router.HandleFunc("/tasks/{id}",deleteTask).Methods("DELETE")
	router.HandleFunc("/tasks/{id}",updateTask).Methods("PUT")

	//Runserver
	log.Fatal(http.ListenAndServe(":8000",router))

}
