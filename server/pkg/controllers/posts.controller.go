package controllers

import (
	"database/sql"
	"fmt"
	"io"
	"log"
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

	utils.WriteJSON(w, posts)
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

	utils.WriteJSON(w, post)
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

	utils.WriteJSON(w, posts)
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

	utils.WriteJSON(w, posts)
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

	defer file.Close()

	fileSize := fileHeader.Size
	log.Printf("the size of the file uploaded is: %v", utils.CalcFileSize(fileSize))

	if fileSize > utils.MAX_UPLOAD_SIZE {
		utils.WriteError(w, fmt.Errorf("file is too big, the maximum allowed is 2MB"), http.StatusBadRequest)
		return
	}

	buff := make([]byte, 512)
	_, err = file.Read(buff)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	filetype := http.DetectContentType(buff)
	if filetype != "image/jpeg" && filetype != "image/png" {
		utils.WriteError(w, fmt.Errorf("the provided file format is not allowed. Please upload a JPEG or PNG image"), http.StatusBadRequest)
		return
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	resp, err := utils.CloudinaryUpload(file, strings.Split(fileHeader.Filename, ".")[0])
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
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
		Image: resp.SecureURL,
		UserID: currentUser.UserID,
    IsPublished: isPublished,
    IsDraft: isDraft,
	})
	if err != nil {
		utils.WriteError(w, err, 500)
		return
	}
	
	utils.WriteJSON(w, post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.WriteError(w, err, 400)
		return
	}

	body, err := utils.ReadJSON[UpdatePostParams](r.Body)
	if err != nil {
		utils.WriteError(w, err, 400)
		return
	}

	foundPost, err := db.GetPostByID(r.Context(), id)
  if err != nil {
    utils.WriteError(w, err, 404)
    return
  }

  var title sql.NullString
  if body.Title != nil {
      title.String = *body.Title
      title.Valid = true
  } else {
      title = sql.NullString{String: foundPost.Title, Valid: true}
  }

  var postBody sql.NullString
  if body.Body != nil {
      postBody.String = *body.Body
      postBody.Valid = true
  } else {
      postBody = sql.NullString{String: foundPost.Body, Valid: true}
  }

  var published sql.NullBool
  if body.IsPublished != nil {
		published.Bool = *body.IsPublished
		published.Valid= true
  } else {
		published = sql.NullBool{Bool: foundPost.IsPublished, Valid: true}
	}

	var draft sql.NullBool
  if body.IsDraft != nil {
		draft.Bool = *body.IsDraft
		draft.Valid = true
  } else {
		draft = sql.NullBool{Bool: foundPost.IsDraft, Valid: true}
	}

	post, err := db.UpdatePost(r.Context(), database.UpdatePostParams{
		ID: id,
		Title: title,
		Body: postBody,
		IsPublished: published,
		IsDraft: draft,
	})
	if err != nil {
		utils.WriteError(w, err, 500)
		return
	}

	utils.WriteJSON(w, post)
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

	utils.WriteJSON(w, fmt.Sprintf("Post with id %s has been deleted!", id))
}