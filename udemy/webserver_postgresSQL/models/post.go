package models

type Post struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
}

// helper functions to retirn a single or a slice of posts
func GetPost() Post {
	var post Post
	return post
}

func GetPosts() []Post {
	var posts []Post
	return posts
}