package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
)

type Post struct {
	Id       int
	Content  string
	Author   string
	Comments []Comment
}

type Comment struct {
	Id      int
	Content string
	Author  string
	Post    *Post
}

var Db *sql.DB

func init() {
	fmt.Println("init invoke...")
	var err error
	Db, err = sql.Open("postgres", "port=5432 user=donscoco password=donscoco dbname=donscocodb sslmode=disable")
	if err != nil {
		panic(err)
	}
	fmt.Println("init finished...")

}

//create post
func (post *Post) create() (err error) {
	err = Db.QueryRow("insert into posts (content, author) values ($1, $2) returning id", post.Content, post.Author).Scan(&post.Id)
	return
}

//create comment
func (comment *Comment) create() (err error) {
	if comment.Post == nil {
		return errors.New("Post not found")
	}
	err = Db.QueryRow("insert into comments (content, author, post_id) values ($1, $2, $3) returning id", comment.Content, comment.Author, comment.Post.Id).Scan(&comment.Id)
	return
}

//getpost
func GetPost(id int) (post Post, err error) {
	err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	if err != nil {

	}
	rows, err := Db.Query("select id, content, author from comments where post_id = $1", id)
	for rows.Next() {
		comment := Comment{Post: &post}
		rows.Scan(&comment.Id, &comment.Content, &comment.Author)
		if err != nil {
			return
		}
		post.Comments = append(post.Comments, comment)
	}
	defer rows.Close()
	return
}
func main() {
	post := Post{Content: "Hello World!", Author: "Sau Sheong"}
	post.create()

	// Add a comment
	comment := Comment{Content: "Good post!", Author: "Joe", Post: &post}
	comment.create()
	readPost, _ := GetPost(post.Id)

	fmt.Println(readPost)                  // {1 Hello World! Sau Sheong [{1 Good post! Joe 0xc20802a1c0}]}
	fmt.Println(readPost.Comments)         // [{1 Good post! Joe 0xc20802a1c0}]
	fmt.Println(readPost.Comments[0].Post) // &{1 Hello World! Sau Sheong [{1 Good post! Joe 0xc20802a1c0}]}
}
