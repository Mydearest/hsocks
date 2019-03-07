package call

import (
	"github.com/google/gopacket/layers"
)

type ProxyServer struct {

}
type ProxyClient struct {

}


type PacketRequest struct {
	Packet []byte
	SrcPort layers.TCPPort
	DstPort layers.TCPPort
}

type PacketResponse struct {
	IsTimeout bool
}

//	远端服务器被调用该方法
//func (client ProxyServer)ProxyPacket(req PacketRequest ,res *PacketResponse) error{
//
//	return nil
//}

// local服务器被远端调用
func (client ProxyClient)ReturnPacket(req PacketRequest ,res *PacketResponse) error{

	return nil
}