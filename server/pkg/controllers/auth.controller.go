package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/takanome-dev/blog-with-astro-golang/pkg/utils"
)

type AuthParams struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct{
	Email    string `json:"email"`
	Username string `json:"username"`
}

type AuthResponse struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	body, err := utils.ReadJSON[AuthParams](r.Body)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}
	
	_, err = db.GetUserByEmail(r.Context(), body.Email)
	if err == nil {
		utils.WriteError(w, errors.New("email already registered"), http.StatusConflict)
		return
	}

	hashedPassword, err := utils.HashPassword(body.Password)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	newUser, err := CreateUser(r.Context(), &AuthParams{
		Username: body.Username,
		Email: body.Email,
		Password: hashedPassword,
	})
	if err != nil {
		log.Println(err)
		utils.WriteError(w, errors.New("username already taken"), http.StatusInternalServerError)
		return
	}

	token, err := utils.GenerateJwt(newUser.ID)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, AuthResponse{
		User: UserResponse{
			Email: newUser.Email,
			Username: newUser.Username,
		},
		Token: token,
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := utils.ReadJSON[LoginParams](r.Body)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}
	
	user, err := db.GetUserByEmail(r.Context(), body.Email)
	if err != nil {
		utils.WriteError(w, errors.New("email or password invalid"), http.StatusBadRequest)
		return
	}

	isPasswordValid := utils.VerifyPassword(body.Password, user.Password)
	if !isPasswordValid {
		utils.WriteError(w, errors.New("email or password invalid"), http.StatusBadRequest)
		return
	}

	token, err := utils.GenerateJwt(user.ID)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	// Maybe save token in cookies??
	
	utils.WriteJSON(w, AuthResponse{
		User: UserResponse{Username: user.Username, Email: user.Email},
		Token: token,
	})
}