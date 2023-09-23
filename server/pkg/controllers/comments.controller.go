package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/takanome-dev/blog-with-astro-golang/internal/database"
	"github.com/takanome-dev/blog-with-astro-golang/pkg/utils"
)

type CreateCommentParams struct {
	Body   string    `json:"body"`
	PostID uuid.UUID `json:"post_id"`
}
type UpdateCommentParams struct {
	Body sql.NullString `json:"body"`
}

func GetAllComments(w http.ResponseWriter, r *http.Request) {
	comments, err := db.GetAllComments(r.Context())
	if err != nil {
		utils.WriteError(w, err, 500)
		return
	}

	utils.WriteJSON(w, comments)
}

func GetCommentByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.WriteError(w, err, 400)
		return
	}

	comments, err := db.GetCommentByID(r.Context(), id)
	if err != nil {
		utils.WriteError(w, err, 404)
		return
	}

	utils.WriteJSON(w, comments)
}

func CreateComment(w http.ResponseWriter, r *http.Request) {
	body, err := utils.ReadJSON[CreateCommentParams](r.Body)
	if err != nil {
		utils.WriteError(w, err, 400)
		return
	}

	currentUser, ok := utils.CtxValue[utils.JwtUser](r.Context()); 
	if !ok {
		utils.WriteError(w, fmt.Errorf("something went wrong when retrieving user id from context"), 400)
		return
	}

	comment, err := db.CreateComment(r.Context(), database.CreateCommentParams{
		ID: uuid.New(),
    Body: body.Body,
    UserID: currentUser.UserID,
    PostID: body.PostID,
	})
	if err != nil {
		utils.WriteError(w, err, 500)
		return
	}
	
	utils.WriteJSON(w, comment)
}

func UpdateComment(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.WriteError(w, err, 400)
		return
	}

	body, err := utils.ReadJSON[UpdateCommentParams](r.Body)
	if err != nil {
		utils.WriteError(w, err, 400)
		return
	}

	comment, err := db.UpdateComment(r.Context(), database.UpdateCommentParams{
		ID: id,
		Body: body.Body,
	})
	if err != nil {
		utils.WriteError(w, err, 500)
		return
	}

	utils.WriteJSON(w, comment)
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.WriteError(w, err, 400)
		return
	}

	err = db.DeleteComment(r.Context(), id)
	if err != nil {
		utils.WriteError(w, err, 500)
		return
	}

	utils.WriteJSON(w, fmt.Sprintf("Comment with id %s has been deleted!", id))
}