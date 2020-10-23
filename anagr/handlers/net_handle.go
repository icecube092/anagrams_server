package handlers

import (
	"encoding/json"
	"net/http"
)

func getHandle(handler *HttpHandler, request *http.Request)[]string{
	if request.URL.Path != "/get"{
		return []string{"Use '/get' path for GET request"}
	}
	if word := request.URL.Query().Get("word"); word != ""{
		return getWordHandle(handler, word)
	}else{
		return []string{"Use 'word' key for GET request"}
	}
}

func postHandle(handler *HttpHandler, request *http.Request)string{
	wordsChan := make(chan []string, 1)
	if request.URL.RequestURI() != "/load"{
		return "Use '/load' path for POST request"
	}
	if isValid(request, wordsChan){
		postWordHandle(handler, <-wordsChan)
		return "Success"
	}else{
		return "Not valid request"
	}
}

func isValid(request *http.Request, channel chan []string)(bool){
	contentType := []string{"application/x-www-form-urlencoded"}
	for index := range request.Header["Content-Type"]{
		if request.Header["Content-Type"][index] != contentType[index]{
			return false
		}
	}
	err := request.ParseForm()
	if err != nil{
		return false
	}
	form := request.Form
	var words []string
	for array := range form {
		err = json.Unmarshal([]byte(array), &words)
		if err != nil{
			return false
		}
		break
	}
	if len(words) == 0{
		return false
	}
	channel<-words
	return true
}
