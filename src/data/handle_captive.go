package data

import (
	"call"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"log"
	"utils/gopool"
)

func init(){
	packetSendExecutor = gopool.NewSingleExecutor()
	packetSendExecutor.Start()
}


func (task packetSendTask) Run() error {
	log.Printf("dst:%d ,src:%d\n" ,task.tcpLayer.DstPort ,task.tcpLayer.SrcPort)
	packetReq := call.PacketRequest{
		DstPort:task.tcpLayer.DstPort,
		SrcPort:task.tcpLayer.SrcPort,
		Packet:task.packet.Data(),
	}
	packetRes := &call.PacketResponse{}
	if err := call.Invoke(call.M_ProxyPacket ,packetReq ,packetRes);err != nil{
		return err
	}
	return nil
}


type packetSendTask struct {
	packet gopacket.Packet
	tcpLayer *layers.TCP
}

var packetSendExecutor *gopool.SingleExecutor

func packetHandle(pack gopacket.Packet ,tcp *layers.TCP){
	task := packetSendTask{
		packet:pack,
		tcpLayer:tcp,
	}
	if err := packetSendExecutor.Submit(task);err != nil{
		log.Println("Submit packet task(client->server) error : " ,err)
	}
}
