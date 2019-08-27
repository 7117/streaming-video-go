package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func RegisterHandles()*httprouter.Router{
	
	router:=httprouter.New();

	router.POST("/user",CreateUser)
	router.POST("/user/:user_name",Login)

	return router;
}

func main(){
	r:=RegisterHandles()
	// 会堵塞  会注册r函数
	http.ListenAndServe("127.0.0.1:8000",r)
}

