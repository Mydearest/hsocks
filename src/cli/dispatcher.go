package cli

import (
	"call"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

func proxySsl(writer http.ResponseWriter, req *http.Request){
	dstCon ,err := net.DialTimeout("tcp", req.Host, time.Second*5)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusServiceUnavailable)
		return
	}
	writer.WriteHeader(http.StatusOK)
	hijacker ,ok := writer.(http.Hijacker)
	if !ok {
		http.Error(writer, "Hijacking not supported", http.StatusInternalServerError)
		return
	}
	cliCon ,_ ,err := hijacker.Hijack()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusServiceUnavailable)
	}
	go transfer(dstCon, cliCon)
	go transfer(cliCon, dstCon)
}

func transfer(destination io.WriteCloser, source io.ReadCloser){
	defer destination.Close()
	defer source.Close()
	if _ ,err := io.Copy(destination, source);err != nil{
		log.Println(err)
	}
}

func proxyHttp(writer http.ResponseWriter, req *http.Request){
	rpcClient := call.DefaultClientProxy
	rpcRequest := rpcClient.BuildProxyRequest(*req)
	rpcResponse := &call.ProxyResponse{}
	if err := rpcClient.Invoke(call.M_ProxyHttp ,rpcRequest ,rpcResponse);err != nil{
		log.Println(err)
		return
	}
	copyHeaders(writer.Header() ,rpcResponse.Headers)
	writer.WriteHeader(rpcResponse.StatusCode)
	if _ ,err := writer.Write(rpcResponse.Body);err != nil{
		log.Println(err)
	}
}

func copyHeaders(dst, src http.Header) {
	for k, arrVal := range src {
		for _, val := range arrVal {
			dst.Add(k, val)
		}
	}
}


