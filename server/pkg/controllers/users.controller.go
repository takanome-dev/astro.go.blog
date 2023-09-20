package controllers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/takanome-dev/blog-with-astro-golang/internal/database"
	"github.com/takanome-dev/blog-with-astro-golang/pkg/config"
	"github.com/takanome-dev/blog-with-astro-golang/pkg/utils"
)

var db *database.Queries

func init() {
	db = config.GetDB()
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := db.GetAllUsers(r.Context())
	if err != nil {
		utils.WriteError(w, err, 500)
		return
	}

	utils.WriteJSON(w, utils.MarshalUsersResponse(users))
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := uuid.Parse(idStr)
	
	if err != nil {
		utils.WriteError(w, err, 400)
		return
	}

	user, err := db.GetUserByID(r.Context(), id)
	if err != nil {
		utils.WriteError(w, err, 404)
		return
	}

	utils.WriteJSON(w, utils.MarshalUserResponse(user))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := utils.ReadJSON[database.CreateUserParams](r.Body)
	if err != nil {
		utils.WriteError(w, err, 400)
		return
	}

	user, err := db.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		Username: body.Username,
		Email: body.Email,
		Password: body.Password,
	})
	if err != nil {
		utils.WriteError(w, err, 500)
		return
	}

	utils.WriteJSON(w, utils.MarshalUserResponse(user))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {}
func DeleteUser(w http.ResponseWriter, r *http.Request) {}