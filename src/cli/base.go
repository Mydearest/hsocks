package cli

import (
	"log"
	"net/http"
	"utils"
)

var Mur *http.ServeMux

func init(){
	Mur = http.NewServeMux()
}

func StartClient(){
	StartHttpServer()
}

func StartHttpServer(){
	var server = http.Server{
		Handler:Mur,
		Addr:utils.Args.ClientMode,
	}
	dispatcher(Mur)
	if err := server.ListenAndServe();err != nil{
		log.Println(err.Error())
	}
}
