package call

import (
	"log"
	"net/rpc"
	"utils"
)

const M_ProxyPacket = "ProxyServer.ProxyPacket"

var DefaultRpcClient *rpc.Client

func init(){
	DefaultRpcClient = NewRpcClient(utils.Args.ProxyServer)
}

func NewRpcClient(serverAddr string) *rpc.Client{
	conn, err := rpc.DialHTTP("tcp", serverAddr)
	if err != nil{
		log.Fatalln("Init rpc client error : " ,err)
	}
	return conn
}

func Invoke(method string ,request PacketRequest ,response *PacketResponse) error{
	if err := DefaultRpcClient.Call(method ,request ,response);err != nil{
		log.Println("Invoke remote rpc method error : " ,err)
		return err
	}
	return nil
}