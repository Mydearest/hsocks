package ser

import (
	"call"
	"log"
	"net/rpc"
	"os"
	"utils"
)

func StartProxyServer(){
	//call.InitServerProxy()
	proxyRpc := &call.ProxyRpc{}
	if err := rpc.Register(proxyRpc);err != nil{
		log.Fatalln("register rpc server error : " ,err)
	}else{
		if utils.Args.CloudFoundry{
			call.StartRpc(os.Getenv("PORT"))
		}else{
			call.StartRpc(utils.Args.ServerMode)
		}
	}
}
