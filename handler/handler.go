package handler

import (
	"curdapis/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	var users []models.Student
	models.Instance.Find(&users)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func GetStudents(w http.ResponseWriter, r *http.Request) {
	userid := mux.Vars(r)["id"]
	var user models.Student
	models.Instance.First(&user, userid)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func CreateStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userdata models.Student
	json.NewDecoder(r.Body).Decode(&userdata)
	models.Instance.Create(&userdata)
	json.NewEncoder(w).Encode(userdata)
}

func UpdateStudents(w http.ResponseWriter, r *http.Request) {
	userid := mux.Vars(r)["id"]
	var user models.Student
	models.Instance.First(&user, userid)
	json.NewDecoder(r.Body).Decode(&user)
	models.Instance.Save(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func DeleteStudents(w http.ResponseWriter, r *http.Request) {
	userid := mux.Vars(r)["id"]
	var user models.Student
	models.Instance.Delete(&user, userid)
	json.NewEncoder(w).Encode("Product Deleted Successfully!")
}
