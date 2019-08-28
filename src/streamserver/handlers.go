package main 

import (
	"io"
	"os"
	"net/http"
	"html/template"
	"io/ioutil"
	"time"
	"log"
	"github.com/julienschmidt/httprouter"
)

func testPageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, _ := template.ParseFiles("./videos/upload.html")
 
    t.Execute(w, nil)
}

// 获取组装视频路径
// 打开视频
// 设置浏览器格式
// ServeContent  // ServeContent使用provided ReadSeeker进行回复请求 
// defer关闭
func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	vl := VIDEO_DIR + vid

	video, err := os.Open(vl)
	if err != nil {
		log.Printf("Error when try to open file: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}
	// 设置格式  浏览器自动识别为mp4 进行识别视频
	w.Header().Set("Content-Type", "video/mp4")
	http.ServeContent(w, r, "", time.Now(), video)

	defer video.Close()
}

