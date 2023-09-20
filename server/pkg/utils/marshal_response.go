package utils

import (
	"time"

	"github.com/google/uuid"
	"github.com/takanome-dev/blog-with-astro-golang/internal/database"
)

type User struct {
	ID        uuid.UUID  `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func MarshalUserResponse(user database.User) User {
	return User{
		ID: user.ID,
		Username: user.Username,
		Email: user.Username,
		Password: user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func MarshalUsersResponse(users []database.User) []User {
	var marshalledUsers []User

	for _, user := range users {
		marshalledUsers = append(marshalledUsers, MarshalUserResponse(user))
	}

	return marshalledUsers
}

type Post struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Body        string    `json:"body"`
	UserID      uuid.UUID `json:"user_id"`
	IsPublished bool      `json:"is_published"`
	IsDraft     bool      `json:"is_draft"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func MarshalPostResponse(post database.Post) Post {
	return Post{
		ID: post.ID,
		Title: post.Title,
		Body: post.Body,
		UserID: post.UserID,
		IsPublished: post.IsPublished,
		IsDraft: post.IsDraft,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
}

func MarshalPostsResponse(posts []database.Post) []Post {
	var marshalledPosts []Post

	for _, post := range posts {
		marshalledPosts = append(marshalledPosts, MarshalPostResponse(post))
	}

	return marshalledPosts
}