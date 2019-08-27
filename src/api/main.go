package main 
import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

type middleWareHandler struct {
	// router结构体
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	// m是一个新值  
	m := middleWareHandler{}
	// 让原来的接口r  等于自己定义的m结构体的r
	// 两个可以相等
	m.r = r
	// handle是一个接口  进行接收m结构体
	return m
}

// 我们自己的方法
func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 进行校验session
	validateUserSession(r)
	// 原来的serverHTTP()
	m.r.ServeHTTP(w, r)
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", CreateUser)

	router.POST("/user/:user_name", Login)

	return router
}

func main() {
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	http.ListenAndServe("127.0.0.1:8000", mh)
}

