package main 

import (
	"io"
	"os"
	"net/http"
	"html/template"
	"io/ioutil"
	"time"
	"log"
	"fmt"
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
	fmt.Println(vl);

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

func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 能够读取到的最大的缓冲区内容
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	// ParseMultipartForm表单的最大
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "File is too big")
		return 
	}

	// 从表单拿到文件
	// file对应表单的name
	// 第二个是类型
	file, _, err := r.FormFile("file")
	if err != nil {
		log.Printf("Error when try to get file: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return 
	}

	// 读取到我们数据里面 
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file error: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
	}

	// 获取文件名
	fn := p.ByName("vid-id")
	// 保存文件
	err = ioutil.WriteFile(VIDEO_DIR + fn, data, 0777)
	if err != nil {
		log.Printf("Write file error: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	// 反馈code与信息
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Uploaded successfully")
}
