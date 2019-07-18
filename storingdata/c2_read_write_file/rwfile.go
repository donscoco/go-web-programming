package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

var data = []byte("Hello World!\n")

func main() {
	// write to file and read from file using WriteFile and ReadFile
	err := ioutil.WriteFile("storingdata/c2_read_write_file/data1", data, 0644)
	if err != nil {
		panic(err)
	}
	bytes, _ := ioutil.ReadFile("storingdata/c2_read_write_file/data1")
	// convert bytes to string
	fmt.Println(string(bytes))

	// write to file and read from file using the File struct
	//create or open file by os
	fileStore, err := os.Create("storingdata/c2_read_write_file/data2")
	defer fileStore.Close()
	//store data from memory to driver without using buffer
	finishcount, err := fileStore.Write(data)
	fmt.Printf("Wrote %d bytes to driver\n", finishcount)

	//open file by os
	fileLoad, err := os.Open("storingdata/c2_read_write_file/data2")
	defer fileStore.Close()
	//create a container to save data from driver
	var container = make([]byte, len(data))
	//load data from driver to memory without using buffer
	finishcount, err = fileLoad.Read(container)
	fmt.Printf("Read %d bytes from file\n", finishcount)
	fmt.Println(string(container))

}
