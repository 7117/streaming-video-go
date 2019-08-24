package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func RegisterHandles()*httprouter.Router{
	router:=httprouter.New();

	router.POST("/user",CreateUser)

	return router;
}

func main(){
	r:=RegisterHandles()
	http.ListenAndServe("127.0.0.1:8000",r)

}