package controllers

import (
	"errors"
	"net/http"
	"time"

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
		utils.WriteError(w, errors.New("username already taken"), http.StatusInternalServerError)
		return
	}

	exp := time.Now().Add(60*60*24*7*time.Second)

	token, err := utils.GenerateJwt(newUser.ID, exp)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	// cookie := utils.EncodeCookie(token, exp)
	// http.SetCookie(w, cookie)

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

	exp := time.Now().Add(60*60*24*7*time.Second)

	token, err := utils.GenerateJwt(user.ID, exp)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	// TODO: use cookies instead
	// cookie, err := utils.EncodeCookie(token, exp)
	// if err != nil {
	// 	utils.WriteError(w, err, 500)
	// 	return
	// }
	// http.SetCookie(w, cookie)
	
	utils.WriteJSON(w, AuthResponse{
		User: UserResponse{Username: user.Username, Email: user.Email},
		Token: token,
	})
}