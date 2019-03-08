package call

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func RegisterRpc(a interface{}) error{
	return rpc.Register(a)
}

func StartRpc(addr string){
	rpc.HandleHTTP()
	if listener ,err := net.Listen("tcp" ,":"+addr);err != nil{
		log.Println("Start rpc server error : " ,err)
	}else {
		if err := http.Serve(listener ,nil);err != nil{
			log.Println("Http server error : " ,err)
		}
	}
}
