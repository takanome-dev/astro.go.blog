package controllers

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/takanome-dev/astro.go.blog/internal/database"
	"github.com/takanome-dev/astro.go.blog/pkg/config"
	"github.com/takanome-dev/astro.go.blog/pkg/utils"
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

	err = utils.WriteJSON(w, users)
	if err != nil {
		utils.WriteError(w, err, 500)
		return
	}
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

	err = utils.WriteJSON(w, user)
	if err != nil {
		utils.WriteError(w, err, 500)
		return
	}
}

func GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	currentUserID, ok := utils.CtxValue[utils.JwtUser](r.Context()); 
	if !ok {
		utils.WriteError(w, fmt.Errorf("something went wrong when retrieving user id from context"), 400)
		return
	}

	user, err := db.GetUserByID(r.Context(), currentUserID.UserID)
	if err != nil {
		utils.WriteError(w, err, 404)
		return
	}

	err = utils.WriteJSON(w, user)
	if err != nil {
		utils.WriteError(w, err, 500)
		return
	}
}

func GetCurrentUserKPIs(w http.ResponseWriter, r *http.Request) {
	currentUserID, ok := utils.CtxValue[utils.JwtUser](r.Context()); 
	if !ok {
		utils.WriteError(w, fmt.Errorf("something went wrong when retrieving user id from context"), 400)
		return
	}

	user, err := db.GetUserKPIs(r.Context(), currentUserID.UserID)
	if err != nil {
		utils.WriteError(w, err, 404)
		return
	}

	var posts interface{}
	err = json.Unmarshal(user.LastThreePosts.([]byte), &posts)
	if err != nil {
			utils.WriteError(w, err, 500)
			return
	}

	var comments interface{}
	err = json.Unmarshal(user.LastThreeComments.([]byte), &comments)
	if err != nil {
			utils.WriteError(w, err, 500)
			return
	}

	result := struct {
		User              database.User        `json:"user"`
		LastThreePosts    interface{} `json:"last_three_posts"`
		LastThreeComments interface{} `json:"last_three_comments"`
	}{
			User:              user.User,
			LastThreePosts:    posts,
			LastThreeComments: comments,
	}

	err = utils.WriteJSON(w, result)
	if err != nil {
		utils.WriteError(w, err, 500)
		return
	}
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

	err = utils.WriteJSON(w, user)
	if err != nil {
		utils.WriteError(w, err, 500)
		return
	}
}

func CreateUser(ctx context.Context, user *AuthParams) (database.User, error) {
	return db.CreateUser(ctx, database.CreateUserParams{
		ID: uuid.New(),
		Username: user.Username,
		Email: user.Email,
		Password: user.Password,
	})
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.WriteError(w, err, 400)
		return
	}

	var image string
	var url string
	image = r.FormValue("image")

	if strings.HasPrefix(image, "data:image") {
		file, fileHeader, err := r.FormFile("image")
		if err != nil {
			utils.WriteError(w, err, http.StatusBadRequest)
			return
		}

		url, err = utils.HandleImage(file, fileHeader)
		if err != nil {
			utils.WriteError(w, err, http.StatusBadRequest)
			return
		}
	} else {
		url = image
	}

	userBio := r.FormValue("bio") 
	userEmail := r.FormValue("email") 
  userFullName := r.FormValue("full_name")
	userGithubUsername := r.FormValue("github_username")
	userLocation := r.FormValue("location")
	userTwitterUsername := r.FormValue("twitter_username")
	userUsername := r.FormValue("username")
	userWebsiteUrl := r.FormValue("website_url")
	
	foundUser, err := db.GetUserByID(r.Context(), id); 
	if err != nil {
		utils.WriteError(w, err, 400)
		return;
	}
	
	var bio sql.NullString
	if userBio != "" {
		bio = sql.NullString{String: userBio, Valid: true}
	} else {
		bio = sql.NullString{String: foundUser.Bio, Valid: true }
	}

	var email sql.NullString
	if userEmail != "" {
		email = sql.NullString{String: userEmail, Valid: true}
	} else {
		email = sql.NullString{String: foundUser.Email, Valid: true}
	}

	var name sql.NullString
	if userFullName != "" {
		name = sql.NullString{String: userFullName, Valid: true}
	} else {
		name = sql.NullString{String: foundUser.Name, Valid: true}
	}

	var github_username sql.NullString
	if userGithubUsername != "" {
		github_username = sql.NullString{String: userGithubUsername, Valid: true}
	} else {
		github_username = sql.NullString{String: foundUser.GithubUsername, Valid: true}
	}

	var location sql.NullString
	if userLocation != "" {
		location = sql.NullString{String: userLocation, Valid: true}
	} else {
		location = sql.NullString{String: foundUser.Location, Valid: true}
	}

	var twitter_username sql.NullString
	if userTwitterUsername != "" {
		twitter_username = sql.NullString{String: userTwitterUsername, Valid: true}
	} else {
		twitter_username = sql.NullString{String: foundUser.TwitterUsername, Valid: true}
	}

	var username sql.NullString
	if userUsername != "" {
		username = sql.NullString{String: userUsername, Valid: true}
	} else {
		username = sql.NullString{String: foundUser.Username, Valid: true}
	}

	var website_url sql.NullString
	if userWebsiteUrl != "" {
		website_url = sql.NullString{String: userWebsiteUrl, Valid: true}
	} else {
		website_url = sql.NullString{String: foundUser.WebsiteUrl, Valid: true}
	}

	updateUser := database.UpdateUserParams{
		ID: uuid.NullUUID{UUID: id, Valid: true},
		Image: sql.NullString{String: url, Valid: true},
		Bio: bio,
		Email: email,
		Name: name,
		GithubUsername: github_username,
		Location: location,
		TwitterUsername: twitter_username,
		Username: username,
		WebsiteUrl: website_url,
	}

	user, err := db.UpdateUser(r.Context(), updateUser)
	if err != nil {
		utils.WriteError(w, err, 500)
		return
	}

	err = utils.WriteJSON(w, user)
	if err != nil {
		utils.WriteError(w, err, 500)
		return
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {}