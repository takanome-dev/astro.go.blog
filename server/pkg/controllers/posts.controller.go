package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/takanome-dev/astro.go.blog/internal/database"
	"github.com/takanome-dev/astro.go.blog/pkg/utils"
)

type CreatePostParams struct {
	Title       string    `json:"title"`
	Body        string    `json:"body"`
	Image       string     `json:"image"`
	IsPublished bool      `json:"is_published,omitempty"`
	IsDraft     bool      `json:"is_draft,omitempty"`
}
type UpdatePostParams struct {
	Title       *string `json:"title"`
	Body        *string `json:"body"`
	IsPublished *bool   `json:"is_published"`
	IsDraft     *bool   `json:"is_draft"`
}
type Comment struct {
	ID        uuid.UUID    `json:"id"`
	Body      string       `json:"body"`
	UserID    uuid.UUID    `json:"user_id"`
	PostID    uuid.UUID    `json:"post_id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
	EditedAt  string `json:"edited_at"`
}


func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := db.GetAllPosts(r.Context())
	if err != nil {
		utils.WriteError(w, err, 500)
		return
	}

	err = utils.WriteJSON(w, posts)
	if err != nil {
		utils.WriteError(w, err, 500)
		return
	}
}

func GetPostByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.WriteError(w, err, 400)
		return
	}

	post, err := db.GetPostByID(r.Context(), id)
	if err != nil {
		utils.WriteError(w, err, 404)
		return
	}

	var comments []struct {
		Comment Comment `json:"comment"`
		User    database.User    `json:"user"`
	}
	type GetPostByIDRow struct {
		Post     database.Post        `json:"post"`
		User     database.User        `json:"user"`
		Comments interface{}          `json:"comments"`
	}

	if post.Comments == nil {
		post.Comments = "[]"
	}

	err = json.Unmarshal([]byte(post.Comments.(string)), &comments)
	if err != nil {
		log.Printf("err when unmarshalling comments: %v", err)
		utils.WriteError(w, err, 500)
		return
	}

	if comments == nil {
		comments = []struct {
			Comment Comment `json:"comment"`
			User    database.User    `json:"user"`
		}{}
	}

	err = utils.WriteJSON(w, GetPostByIDRow{
		Post: post.Post,
		User: post.User,
		Comments: comments,
	})
	if err != nil {
		utils.WriteError(w, err, 500)
		return
	}
}

func GetPostsByUserID(w http.ResponseWriter, r *http.Request) {
	userIdStr := mux.Vars(r)["userId"]
	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		utils.WriteError(w, err, 400)
		return
	}

	posts, err := db.GetPostsByUserID(r.Context(), userId)
	if err != nil {
		utils.WriteError(w, err, 404)
		return
	}

	err = utils.WriteJSON(w, posts)
	if err != nil {
		utils.WriteError(w, err, 500)
		return
	}
}

func GetCurrentUserPosts(w http.ResponseWriter, r *http.Request) {
	currentUser, ok := utils.CtxValue[utils.JwtUser](r.Context()); 
	if !ok {
		utils.WriteError(w, fmt.Errorf("something went wrong when retrieving user id from context"), 400)
		return
	}

	posts, err := db.GetPostsByUserID(r.Context(), currentUser.UserID)
	if err != nil {
		utils.WriteError(w, err, 404)
		return
	}

	err = utils.WriteJSON(w, posts)
	if err != nil {
		utils.WriteError(w, err, 500)
		return
	}
}

func GetCurrentUserDraftPosts(w http.ResponseWriter, r *http.Request) {
	currentUser, ok := utils.CtxValue[utils.JwtUser](r.Context()); 
	if !ok {
		utils.WriteError(w, fmt.Errorf("something went wrong when retrieving user id from context"), 400)
		return
	}

	posts, err := db.GetDraftPostsByUserID(r.Context(), currentUser.UserID)
	if err != nil {
		utils.WriteError(w, err, 404)
		return
	}

	err = utils.WriteJSON(w, posts)
	if err != nil {
		utils.WriteError(w, err, 500)
		return
	}
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	// 10 << 20 -> how much will be stored in memory
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	file, fileHeader, err := r.FormFile("image")
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	url, err := utils.HandleImage(file, fileHeader)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	 title := r.FormValue("title")
	 body := r.FormValue("body")
	 isPublished, _ := strconv.ParseBool(r.FormValue("is_published"))
	 isDraft, _ := strconv.ParseBool(r.FormValue("is_draft"))

	currentUser, ok := utils.CtxValue[utils.JwtUser](r.Context()); 
	if !ok {
		utils.WriteError(w, fmt.Errorf("something went wrong when retrieving user id from context"), 400)
		return
	}

	post, err := db.CreatePost(r.Context(), database.CreatePostParams{
		ID: uuid.New(),
		Title: title,
    Body: body,
		Image: url,
		UserID: currentUser.UserID,
    IsPublished: isPublished,
    IsDraft: isDraft,
	})
	if err != nil {
		utils.WriteError(w, err, 500)
		return
	}
	
	err = utils.WriteJSON(w, post)
	if err != nil {
		utils.WriteError(w, err, 500)
		return
	}
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.WriteError(w, err, 400)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
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
	
	postTitle := r.FormValue("title")
	postBody := r.FormValue("body")
	isPublished, _ := strconv.ParseBool(r.FormValue("is_published"))
	isDraft, _ := strconv.ParseBool(r.FormValue("is_draft"))

	foundPost, err := db.GetPostByID(r.Context(), id)
  if err != nil {
    utils.WriteError(w, err, 404)
    return
  }

  var title sql.NullString
  if postTitle != "" {
      title.String = postTitle
      title.Valid = true
  } else {
      title = sql.NullString{String: foundPost.Post.Title, Valid: true}
  }

  var body sql.NullString
  if postBody != "" {
      body.String = postBody
      body.Valid = true
  } else {
		body = sql.NullString{String: postBody, Valid: true}
  }

  published := sql.NullBool{Bool: isPublished, Valid: true}
	draft := sql.NullBool{Bool: isDraft, Valid: true}

	post, err := db.UpdatePost(r.Context(), database.UpdatePostParams{
		ID: id,
		Title: title,
		Body: body,
		Image: sql.NullString{String: url, Valid: true},
		IsPublished: published,
		IsDraft: draft,
	})
	if err != nil {
		utils.WriteError(w, err, 500)
		return
	}

	err = utils.WriteJSON(w, post)
	if err != nil {
		utils.WriteError(w, err, 500)
		return
	}
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.WriteError(w, err, 400)
		return
	}

	err = db.DeletePost(r.Context(), id)
	if err != nil {
		utils.WriteError(w, err, 500)
		return
	}

	err = utils.WriteJSON(w, fmt.Sprintf("Post with id %s has been deleted!", id))
	if err != nil {
		utils.WriteError(w, err, 500)
		return
	}
}