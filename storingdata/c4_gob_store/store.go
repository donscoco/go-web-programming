package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

var post = Post{Id: 1, Content: "Hello World!", Author: "Sau Sheong"}

//store data
func store(data interface{}, filename string) {
	//create a bytes.Buffer struct
	buffer := new(bytes.Buffer) //can be both a reader and a writer

	//create a encoder by passing the buffer
	encoder := gob.NewEncoder(buffer)

	//use the encoder to encode the data into the buffer
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}
	//write buffer from memory to driver
	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600)
	if err != nil {
		panic(err)
	}
}

//load data
func load(data interface{}, filename string) {
	//read buffer from driver to memory
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	//create a bytes.Buffer struct
	buffer := bytes.NewBuffer(raw)

	//create a decoder from the buffer using the newdecoder
	decoder := gob.NewDecoder(buffer)

	//decode the data
	err = decoder.Decode(data)
	if err != nil {
		panic(err)
	}
}
func main() {
	post := Post{Id: 1, Content: "Hello World!", Author: "Sau Sheong"}
	store(post, "storingdata/c4_gob_store/post")
	var postRead Post
	load(&postRead, "storingdata/c4_gob_store/post")
	fmt.Println(postRead)
}
