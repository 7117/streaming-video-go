package main

import (
	"io"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// goroutine
func CreateUser(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	io.WriteString(w,"cretae user handle")
}

// goroutine
func Login(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	uname:=p.ByName("user_name")
	// w封装的conn.write
	io.WriteString(w,uname)
}
