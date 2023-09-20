package controllers

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/takanome-dev/blog-with-astro-golang/internal/database"
	"github.com/takanome-dev/blog-with-astro-golang/pkg/utils"
)

// type CreatePostParams struct {
// 	ID          uuid.UUID  `json:"id"`
// 	Title       string     `json:"title"`
// 	Body        string     `json:"body"`
// 	UserID      uuid.UUID  `json:"user_id"`
// 	IsPublished bool       `json:"is_published"`
// 	IsDraft     bool       `json:"is_draft"`
// }

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := db.GetAllPosts(r.Context())
	if err != nil {
		utils.WriteError(w, err, 500)
		return
	}

	// log.Printf("posts retrieved from db: %v", posts)
	// TODO: the results is an empty array if there is no posts
	// TODO: but for some reason null is returned

	utils.WriteJSON(w, utils.MarshalPostsResponse(posts))
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

	utils.WriteJSON(w, utils.MarshalPostResponse(post))
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	body, err := utils.ReadJSON[database.CreatePostParams](r.Body)
	if err != nil {
		utils.WriteError(w, err, 400)
		return
	}

	post, err := db.CreatePost(r.Context(), database.CreatePostParams{
		ID: uuid.New(),
		Title: body.Title,
    Body: body.Body,
    UserID: body.UserID,
    IsPublished: body.IsPublished,
    IsDraft: body.IsDraft,
	})
	if err != nil {
		utils.WriteError(w, err, 500)
		return
	}
	
	utils.WriteJSON(w, utils.MarshalPostResponse(post))
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.WriteError(w, err, 400)
		return
	}

	log.Printf("Request post ID: %v\n", id)
	log.Printf("Unmarshal request body: %v\n", r.Body)

	body, err := utils.ReadJSON[database.UpdatePostParams](r.Body)
	if err != nil {
		utils.WriteError(w, err, 400)
		return
	}

	log.Printf("Marshal request body: %v\n", body)

	post, err := db.UpdatePost(r.Context(), body)
	if err != nil {
			utils.WriteError(w, err, 500)
			return
	}

	log.Printf("Updated post: %v\n", post)

	utils.WriteJSON(w, post)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {}