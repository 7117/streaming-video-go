package main

import (
	"html/template"
	// "encoding/json"
	"log"
	"net/http"
	// "io/ioutil"
	// "net/url"
	// "net/http/httputil"
	"github.com/julienschmidt/httprouter"
	// "io"
)

func homeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// r 里面包含着httpserver httpserver里面含有这cookie方法
	cname, err1 := r.Cookie("username")
	sid, err2 := r.Cookie("session")

	// 无值的时候  没有登录的
	if err1 != nil || err2 != nil {
		// 需要赋值的参数
		p := &HomePage{Name: "孙潇"}
		// 将html认定为模板文件
		t, e := template.ParseFiles("D:/Github/Streaming-video/src/webserver/template/home.html")
		if e != nil {
			log.Printf("parsing template home.html error:%s", e)
			return
		}
		// t为模板 p写入到w传递给模板t
		t.Execute(w, p)
		return
	}

	// 有值得时候  登陆了的
	if len(cname.Value) != 0 && len(sid.Value) != 0 {
		http.Redirect(w, r, "D:/Github/Streaming-video/src/webserver/template/userhome.html", http.StatusFound)
		return
	}
}

// func userHomeHandler(w http.ResponseWriter,r *http.Request,ps httprouter.Params){
// 	// // 获取值
// 	// cname,err:=r.Cookie("username")	//Cookie值用于登陆过的
// 	// fname:=r.FormValue("username")		//填充值用于未登录过的
// 	// var p *UserPage

// 	// if err!=nil  {
// 	// 	http.Redirect(w,r,"/",http.StatusFound)
// 	// 	return
// 	// }

// 	// //Cookie值
// 	// if len(cname.Value) !=0 {
// 	// 	p=&UserPage{Name:cname.Value}
// 	// //填充值
// 	// }else if len(fname)!=0{
// 	// 	p=&UserPage{Name:fname}
// 	// }

// 	// 注册页面
// 	p:=&UserPage{Name:"aaaa"}
// 	t,e:=template.ParseFiles("D:/Github/Streamingmedia/src/webserver/template/userhome.html")

// 	if e !=nil {
// 		log.Printf("parsing template home.html error:%s",e)
// 		return
// 	}
// 	t.Execute(w,p)
// 	return
// }

// func apiHandler(w http.ResponseWriter,r *http.Request,ps httprouter.Params){
// 	// 方式
// 	if r.Method!=http.MethodPost {
// 		re,_:=json.Marshal(ErrorRequestNotRecognized)
// 		io.WriteString(w,string(re))
// 		return
// 	}
// 	// body
// 	res,_:=ioutil.ReadAll(r.Body)
// 	apibody:=&ApiBody{}
// 	if err:=json.Unmarshal(res,apibody);err!=nil{
// 		re,_:=json.Marshal(ErrorRequestBodyParseFailed)
// 		io.WriteString(w,string(re))
// 		return
// 	}
// 	// request
// 	request(apibody,w,r)

// 	defer r.Body.Close()
// }

// func proxyHandler(w http.ResponseWriter,r *http.Request,ps httprouter.Params){
// 	u,_:=url.Parse("http://127.0.0.1:9000/")
// 	proxy:=httputil.NewSingleHostReverseProxy(u)  //替换
// 	proxy.ServeHTTP(w,r)
// }
