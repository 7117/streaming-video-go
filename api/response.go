package main 

import (
	"io"
	"encoding/json"
	"net/http"
	"Streaming/api/defs"
)

func sendErrorResponse(w http.ResponseWriter, errResp defs.ErrResponse) {
	// 写入到Header
	w.WriteHeader(errResp.HttpSC)

	// 处理json化
	resStr, _ := json.Marshal(&errResp.Error)
	// 通过io写入数据
	io.WriteString(w, string(resStr))
}

func sendNormalResponse(w http.ResponseWriter, resp string, sc int) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}