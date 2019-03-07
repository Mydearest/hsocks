package call

import (
	"time"
)

type ProxyServer struct {}

type PacketRequest struct {
	Packet []byte
	ProxyTimeout time.Duration
}

type PacketResponse struct {
	Packet []byte
}

//	远端服务器被本地调用
func (client ProxyServer)ProxyPacket(req PacketRequest ,res *PacketResponse) error{

	return nil
}