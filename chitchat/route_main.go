package main

import (
	"github.com/sausheong/gwp/Chapter_2_Go_ChitChat/chitchat/data"
	"net/http"
)

func index(writer http.ResponseWriter, request *http.Request) {
	//get the data from db
	threads, err := data.Threads()
	if err != nil {

	} else {
		// get the session user info
		_, err := session(writer, request)
		if err != nil {
			//
			p(threads)
		} else {
			//

		}
	}
}
