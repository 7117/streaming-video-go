package main

import (
	"io"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func CreateUser(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	io.WriteString(w,"cretae user handle")
}

