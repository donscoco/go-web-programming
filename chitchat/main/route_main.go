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
			generateHTML(writer, threads, "layout", "public.navbar", "index")
		} else {
			generateHTML(writer, threads, "layout", "private.navbar", "index")
		}
	}
}

// GET /err?msg=
// shows the error message page
func err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	_, err := session(writer, request)
	if err != nil {
		generateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		generateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}
