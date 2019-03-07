package call

import (
	"time"
)

type ProxyServer struct {}

type PacketRequest struct {
	Packet []byte
	ProxyTimeout time.Duration	// 这个超时应该由proxy_server来使用，用于控制代理向被墙网站发出请求到收到回复的时间
}

type PacketResponse struct {
	Packet []byte
}

//	远端服务器被本地调用
func (client ProxyServer)ProxyPacket(req PacketRequest ,res *PacketResponse) error{

	return nil
}