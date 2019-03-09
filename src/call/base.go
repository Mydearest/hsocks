package call

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func StartRpc(port string){
	rpc.HandleHTTP()
	if listener ,err := net.Listen("tcp" ,":"+port);err != nil{
		log.Println("Start rpc server error : " ,err)
	}else {
		if err := http.Serve(listener ,nil);err != nil{
			log.Println("Http server error : " ,err)
		}
	}
}
