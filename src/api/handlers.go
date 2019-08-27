package main 

import (
	"io"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"github.com/julienschmidt/httprouter"
	"api/defs"
	"api/dbops"
	"api/session"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}
	
	// Unmarshal-json转化为struct
	if err := json.Unmarshal(res, ubody); err != nil {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return 
	}

	// 进行添加新的用户
	if err := dbops.AddUserCredential(ubody.Username, ubody.Pwd); err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	// session保存至map与db中
	id := session.GenerateNewSessionId(ubody.Username)
	// 发送json的成功数据
	su := &defs.SignedUp{Success: true, SessionId: id}

	if resp, err := json.Marshal(su); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp), 201)
	}
}


func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}