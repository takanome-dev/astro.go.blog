package controllers

import (
	"context"
	"errors"
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

	utils.WriteJSON(w, users)
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

	utils.WriteJSON(w, user)
}

func GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	if username == "" {
		utils.WriteError(w, errors.New("username is required"), http.StatusBadRequest)
		return
	}

	user, err := db.GetUserByUsername(r.Context(), username)
	if err != nil {
		utils.WriteError(w, err, 404)
		return
	}

	utils.WriteJSON(w, user)
}

func CreateUser(ctx context.Context, user *AuthParams) (database.User, error) {
	return db.CreateUser(ctx, database.CreateUserParams{
		ID: uuid.New(),
		Username: user.Username,
		Email: user.Email,
		Password: user.Password,
	})
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {}
func DeleteUser(w http.ResponseWriter, r *http.Request) {}