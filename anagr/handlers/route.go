package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// структура для обработки запросов на сервер
type HttpHandler struct {
	Storage map[string][]string
}

func (h *HttpHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	encoder := json.NewEncoder(response)
	response.WriteHeader(200)
	switch request.Method {
	case "GET":
		log.Println("GET\n", request)
		words := getHandle(h, request)
		err := encoder.Encode(words) // отдает массив строк
		if err != nil {
			log.Println(err)
		}
	case "POST":
		log.Println("POST\n", request)
		success := postHandle(h, request)
		err := encoder.Encode(success) // отдает строку
		if err != nil {
			log.Println(err)
		}
	default:
		log.Printf("%s request from %s\n", request.Method, request.RemoteAddr)
		err := encoder.Encode("Only POST and GET method supported")
		if err != nil {
			log.Println(err)
		}
	}

}

// обработчик GET запросов
func getHandle(handler *HttpHandler, request *http.Request) []string {
	if request.URL.Path != "/get" {
		return []string{"Use '/get' path for GET request"}
	}
	if word := request.URL.Query().Get("word"); word != "" {
		return getWordHandle(handler, word)
	} else {
		return []string{"Use 'word' key with not empty value for GET request"}
	}
}

// обработчик POST запросов
func postHandle(handler *HttpHandler, request *http.Request) string {
	wordsChan := make(chan []string, 1)
	if request.URL.RequestURI() != "/load" {
		return "Use '/load' path for POST request"
	}
	if isValid(request, wordsChan) {
		postWordHandle(handler, <-wordsChan)
		return "Success"
	} else {
		return "Not valid request"
	}
}

// валидация данных POST запроса
func isValid(request *http.Request, channel chan []string) bool {
	contentType := []string{"application/x-www-form-urlencoded"}
	for index := range request.Header["Content-Type"] {
		if request.Header["Content-Type"][index] != contentType[index] {
			return false
		}
	}
	err := request.ParseForm()
	if err != nil {
		return false
	}
	form := request.Form
	var words []string
	for array := range form {
		err = json.Unmarshal([]byte(array), &words)
		if err != nil {
			return false
		}
		break
	}
	if len(words) == 0 {
		return false
	}
	channel <- words
	return true
}
