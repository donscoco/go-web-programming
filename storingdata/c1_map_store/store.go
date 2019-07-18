package main

import "fmt"

//target:
// 1.store a slice of struct

type Post struct {
	Id      int
	Author  string
	Content string
}

var post1 = Post{Id: 1, Content: "Hello World!", Author: "Sau Sheong"}
var post2 = Post{Id: 2, Content: "Bonjour Monde!", Author: "Pierre"}
var post3 = Post{Id: 3, Content: "Hola Mundo!", Author: "Pedro"}
var post4 = Post{Id: 4, Content: "Greetings Earthlings!", Author: "Sau Sheong"}

func store(searchById map[int]*Post, searchByAuthor map[string][]*Post, post Post) {
	searchById[post.Id] = &post
	searchByAuthor[post.Author] = append(searchByAuthor[post.Author], &post)
}
func main() {
	var SearchById = make(map[int]*Post)
	var SearchByAuthor = make(map[string][]*Post)

	store(SearchById, SearchByAuthor, post1)
	store(SearchById, SearchByAuthor, post2)
	store(SearchById, SearchByAuthor, post3)
	store(SearchById, SearchByAuthor, post4)

	fmt.Println(SearchById[1])
	fmt.Println(SearchById[2])

	for _, post := range SearchByAuthor["Sau Sheong"] {
		fmt.Println(post)
	}
	for _, post := range SearchByAuthor["Pedro"] {
		fmt.Println(post)
	}
}
