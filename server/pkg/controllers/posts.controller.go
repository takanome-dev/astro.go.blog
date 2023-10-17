package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/takanome-dev/blog-with-astro-golang/internal/database"
	"github.com/takanome-dev/blog-with-astro-golang/pkg/utils"
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

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := db.GetAllPosts(r.Context())
	if err != nil {
		utils.WriteError(w, err, 500)
		return
	}

	// log.Printf("posts retrieved from db: %v", posts)
	// TODO: the results is an empty array if there is no posts
	// TODO: but for some reason null is returned

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

	err = utils.WriteJSON(w, post)
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

func GetPostsForLoggedInUser(w http.ResponseWriter, r *http.Request) {
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
      title = sql.NullString{String: foundPost.Title, Valid: true}
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