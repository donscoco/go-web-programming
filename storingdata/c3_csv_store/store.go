package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

//target : store data to csv
type Post struct {
	Id      int
	Content string
	Author  string
}

var allPosts = []Post{
	Post{Id: 1, Content: "Hello World!", Author: "Sau Sheong"},
	Post{Id: 2, Content: "Bonjour Monde!", Author: "Pierre"},
	Post{Id: 3, Content: "Hola Mundo!", Author: "Pedro"},
	Post{Id: 4, Content: "Greetings Earthlings!", Author: "Sau Sheong"},
}

func main() {
	//write
	//open an io file by os (memory to driver)
	file, err := os.Create("storingdata/c3_csv_store/posts.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//get the writer   (memory to memory,workspace to buffer to driver)
	writer := csv.NewWriter(file)
	for _, p := range allPosts {
		line := []string{strconv.Itoa(p.Id), p.Content, p.Author}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	//finish buffer to driver
	writer.Flush()

	//read
	// open an io by od
	file, err = os.Open("storingdata/c3_csv_store/posts.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//get the reader
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	records, err := reader.ReadAll()
	var rdata []Post
	for _, item := range records {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		p := Post{
			Id:      int(id),
			Content: item[1],
			Author:  item[2],
		}
		rdata = append(rdata, p)
	}
	fmt.Println(rdata[0].Id)
	fmt.Println(rdata[0].Content)
	fmt.Println(rdata[0].Author)

}
