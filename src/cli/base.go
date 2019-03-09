package cli

import (
	"call"
	"log"
	"logger"
	"net/http"
	"utils"
)


func init(){

}

func StartProxyClient(){
	call.InitClientProxy()
	StartHttpServer()
}

func StartHttpServer(){
	var server = http.Server{
		Handler:http.HandlerFunc(func (writer http.ResponseWriter, req *http.Request){
			logger.Debug.Println(req.Method ,req.Host)
			if req.Method == http.MethodConnect{
				proxySsl(writer ,req)
			}else{
				proxyHttp(writer ,req)
			}
		}),
		Addr:":"+utils.Args.ClientMode,
	}
	if err := server.ListenAndServe();err != nil{
		log.Println(err.Error())
	}
}
