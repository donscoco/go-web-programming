package main

import "net/http"

func authenticate(writer http.ResponseWriter, request *http.Request) {
	//parse param
	err := request.ParseForm()
	// get the user's infomation by user's email
	// check user's password
	// if pass set tht info in the cookie
	// redirect
}
