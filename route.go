package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)


type httpHandler struct{
	storage map[string][]string
	mux     sync.RWMutex
}

func (h *httpHandler) ServeHTTP(response http.ResponseWriter, request *http.Request){
	defer request.Body.Close()
	encoder := json.NewEncoder(response)
	response.WriteHeader(200)
	switch request.Method{
	case "GET":
		words := getHandle(h, request)
		err := encoder.Encode(words)
		if err != nil{
			fmt.Println(err)
		}
	case "POST":
		success := postHandle(h, request)
		err := encoder.Encode(success)
		if err != nil{
			fmt.Println(err)
		}
	default:
		err := encoder.Encode("Only POST and GET method supported")
		if err != nil{
			fmt.Println(err)
		}
	}

}
