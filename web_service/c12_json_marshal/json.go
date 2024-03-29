package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	post := Post{
		Id:      1,
		Content: "Hello World!",
		Author: Author{
			Id:   2,
			Name: "Sau Sheong",
		},
		Comments: []Comment{
			Comment{
				Id:      1,
				Content: "Have a great day!",
				Author:  "Adam",
			},
			Comment{
				Id:      2,
				Content: "How are you today?",
				Author:  "Betty",
			},
		},
	}
	bytes, _ := jsoniter.Marshal(&post)
	ioutil.WriteFile("web_service/c12_json_marshal/post.json", bytes, 0644)

	bytes, _ = ioutil.ReadFile("web_service/c12_json_marshal/post.json")
	fmt.Println(string(bytes))
	container := Post{}
	jsoniter.Unmarshal(bytes, &container)
	fmt.Println(container)
}
