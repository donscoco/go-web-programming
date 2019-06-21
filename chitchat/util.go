package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sausheong/gwp/Chapter_2_Go_ChitChat/chitchat/data"
	"log"
	"net/http"
	"os"
)

type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
}

var config Configuration

//init func
func init() {
	//init config data
	loadCongig()
}

func loadCongig() {
	// os take the file to the memory
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalln("Cannot open config file", err)
	}
	// get a struct to save the config msg
	decoder := json.NewDecoder(file)
	config = Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalln("Cannot get configuration from file", err)
	}

}

// check if user logined
func session(writer http.ResponseWriter, request *http.Request) (sess data.Session, err error) {
	// get cookie from request
	cookie, err := request.Cookie("_cookie")
	if err == nil {
		sess = data.Session{
			Uuid: cookie.Value,
		}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}
func p(msg interface{}) {
	fmt.Println(msg)
}
