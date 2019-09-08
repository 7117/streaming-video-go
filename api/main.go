package main 
import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

// 新属性
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
	// 给一个属性赋值
	m.r = r
	// handle是一个接口  进行接收m结构体
	return m
}

// 我们自己的方法
// serverHTTP方法是在handler接口里面的  必须哟这个接口
func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 进行校验session
	validateUserSession(r)
	// 原来的serverHTTP()  再加上原来的serverHTTP方法
	m.r.ServeHTTP(w, r)
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", CreateUser)
	router.POST("/user/:username", Login)

	// router.GET("/user/:username", GetUserInfo)
	// router.POST("/user/:username/videos", AddNewVideo)

	// router.GET("/user/:username/videos", ListAllVideos)
	// router.POST("/user/:username/videos/:vid-id", DeleteVideo)

	// router.POST("/videos/:vid-id/comments", PostComment)
	// router.GET("/videos/:vid-id/comments", ShowComments)

	return router
}

func main() {
	// router是一个结构体
	// r是router形式的
	r := RegisterHandlers()
	// 将r（路由等信息接口形式）注入到NewMiddleWareHandler里面
	// 返回的是接口形式的数据
	// mh是handle形式的
	mh := NewMiddleWareHandler(r)
	// 字符串   handle形
	http.ListenAndServe("127.0.0.1:8000", mh)
}

