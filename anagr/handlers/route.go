package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)


type HttpHandler struct{
	Storage map[string][]string
}

func (h *HttpHandler) ServeHTTP(response http.ResponseWriter, request *http.Request){
	defer request.Body.Close()
	encoder := json.NewEncoder(response)
	response.WriteHeader(200)
	switch request.Method{
	case "GET":
		log.Println("GET\n", request)
		words := getHandle(h, request)
		err := encoder.Encode(words)
		if err != nil{
			log.Println(err)
		}
	case "POST":
		log.Println("POST\n", request)
		success := postHandle(h, request)
		err := encoder.Encode(success)
		if err != nil{
			log.Println(err)
		}
	default:
		log.Printf("%s request from %s\n", request.Method, request.RemoteAddr)
		err := encoder.Encode("Only POST and GET method supported")
		if err != nil{
			log.Println(err)
		}
	}

}
