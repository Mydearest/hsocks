package call

import (
	"log"
	"net/rpc"
)

const M_ProxyPacket = "ProxyServer.ProxyPacket"
const M_ReturnPacket = "ProxyClient.ReturnPacket"

var DefaultRpcClient *rpc.Client

func init(){
	//DefaultRpcClient = NewRpcClient(utils.Args.ProxyServer)
}

func NewRpcClient(remoteAddr string) *rpc.Client{
	conn, err := rpc.DialHTTP("tcp", remoteAddr)
	if err != nil{
		log.Fatalln("Init rpc client error : " ,err)
	}
	return conn
}

func Invoke(request PacketRequest ,response *PacketResponse) error{
	if err := DefaultRpcClient.Call(M_ProxyPacket ,request ,response);err != nil{
		log.Println("Invoke rpc method error : " ,err)
		return err
	}
	return nil
}