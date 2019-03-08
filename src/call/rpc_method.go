package call

import (
	"errors"
	"fmt"
	"time"
)

type ProxyServer struct {}

type PacketRequest struct {
	Packet []byte
	ProxyTimeout time.Duration	// 这个超时应该由proxy_server来使用，用于控制代理向被墙网站发出请求到收到回复的时间
}

type PacketResponse struct {
	Packet [][]byte	// 服务端将前一次rpc到现在所收到的数据包全部返回
}

//	远端服务器被本地调用
func (client ProxyServer)ProxyPacket(req PacketRequest ,res *PacketResponse) error{
	finishCh := make(chan error)
	go func() {
		finishCh <- handleRpc(req ,res)
	}()
	timer := time.NewTimer(req.ProxyTimeout)
	defer timer.Stop()
	select {
	case <- timer.C:
		return errors.New(fmt.Sprintf("Rpc time out after %d ms" ,req.ProxyTimeout/time.Millisecond))
	case <- finishCh:
		return nil
	}
}

func handleRpc(req PacketRequest ,res *PacketResponse) error{

	return nil
}