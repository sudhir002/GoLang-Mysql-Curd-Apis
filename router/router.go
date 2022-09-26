package router

import (
	"curdapis/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func ApiRouters() string {
	router := mux.NewRouter()
	router.HandleFunc("/allstudents", handler.GetAllStudents).Methods("GET")
	router.HandleFunc("/student/{id}", handler.GetStudents).Methods("GET")
	router.HandleFunc("/createstudent", handler.CreateStudents).Methods("POST")
	router.HandleFunc("/update/{id}", handler.UpdateStudents).Methods("PUT")
	router.HandleFunc("/delete/{id}", handler.DeleteStudents).Methods("DELETE")

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8000", router))
	return "server ended..."
}
