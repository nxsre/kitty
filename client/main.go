package main

import (
	"net/http"
	"log"
)

func main() {


	http.HandleFunc("/job",handler)

	http.ListenAndServe(":8081",nil)
}

func handler(response http.ResponseWriter,request *http.Request)  {

	bys :=make([]byte,1024)
	n,_ := request.Body.Read(bys)
	log.Println(string(bys[:n]))
	response.Write([]byte("ok"))
}
