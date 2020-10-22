package main


import (
	"fmt"
	"net/http"
)


var addr = "localhost:8080"

func main(){
	server := http.Server{}
	server.Addr = addr
	handler := &httpHandler{}
	handler.storage = make(map[string][]string)
	server.Handler = handler

	defer func(){
		fmt.Println("Exit...")
		err := server.Close()
		if err != nil{
			fmt.Println(err)
		}
	}()
	fmt.Println("Listening...")
	err := server.ListenAndServe()
	if err != nil{
		fmt.Println(err)
	}
}

