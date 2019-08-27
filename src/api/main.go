package main 
import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

type middleWareHandler struct {
	// router结构体
	// r是router结构体的数据  
	// router结构体包括handle接口 
	// handle接口（A Handler responds to an HTTP request.响应一个HTTP请求）包括serverHTTP方法
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	// m是一个新值  
	m := middleWareHandler{}
	// 让原来的接口r  等于 自己定义的m结构体的r
	// 里面的内容因为一样  所以两个可以相等   ducktype  m的是自定义的   r是main函数的  内容其实是一个东西
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
	// router是一个结构体
	r := RegisterHandlers()
	// 将r（路由等信息接口形式）注入到NewMiddleWareHandler里面
	// 返回的是接口形式的数据
	mh := NewMiddleWareHandler(r)
	http.ListenAndServe("127.0.0.1:8000", mh)
}

