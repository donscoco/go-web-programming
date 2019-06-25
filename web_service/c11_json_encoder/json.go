package main

import (
	"encoding/json"
	"io"
	"os"
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
	file, err := os.Create("web_service/c11_json_encoder/post.json")
	if err != nil {
		panic(err)
	}
	jsonWriter := io.Writer(file)
	jsonEncoder := json.NewEncoder(jsonWriter)
	jsonEncoder.Encode(&post)

}
