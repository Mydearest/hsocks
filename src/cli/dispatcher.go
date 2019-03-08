package cli

import (
	"call"
	"log"
	"net/http"
)

func dispatcher(mux *http.ServeMux){
	mux.HandleFunc("/" ,handle)
}

func handle(writer http.ResponseWriter, req *http.Request){
	rpcClient := call.DefaultProxy
	rpcRequest := rpcClient.BuildProxyRequest(*req)
	rpcResponse := &call.ProxyResponse{}
	if err := rpcClient.Invoke(call.ProxyHttp ,rpcRequest ,rpcResponse);err != nil{
		log.Println(err)
		return
	}
	for k ,v := range rpcResponse.Headers{
		writer.Header().Set(k ,v[0])
	}
	writer.WriteHeader(rpcResponse.StatusCode)
	if _ ,err := writer.Write(rpcResponse.Body);err != nil{
		log.Println(err)
	}
}
