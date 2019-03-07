package data

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"log"
	"time"
	"utils/gopool"
)

type PacketSendTask struct {
	packet gopacket.Packet
	tcpLayer *layers.TCP
	ipv4Layer *layers.IPv4
}

//var packetSendExecutor *gopool.SingleExecutor()
var packetSendExecutor *gopool.FixedExecutor

func init(){
	//packetSendExecutor = gopool.NewSingleExecutor()
	packetSendExecutor = gopool.NewFixedExecutor(2 ,4 ,time.Second/2)
	packetSendExecutor.Start()
}


func (task PacketSendTask)Info() gopool.TaskInfo{
	return gopool.TaskInfo{}
}


func (task PacketSendTask) Run() error {
	tcp := task.tcpLayer
	ipv4 := task.ipv4Layer
	log.Printf("src->[%s:%d] ,dst->[%s:%d]\n" ,ipv4.SrcIP ,tcp.SrcPort ,ipv4.DstIP ,tcp.DstPort)

	//packetReq := call.PacketRequest{
	//	Packet:task.packet.Data(),
	//	ProxyTimeout:time.Duration(utils.Args.ProxyTimeout),
	//}
	//packetRes := &call.PacketResponse{}
	//// 协程池处理
	//AsynHandle(packetReq ,packetRes)
	return nil
}


func packetHandle(pack gopacket.Packet ,tcp *layers.TCP ,ipv4 *layers.IPv4){
	task := PacketSendTask{
		packet:pack,
		tcpLayer:tcp,
		ipv4Layer:ipv4,
	}
	if err := packetSendExecutor.Submit(task);err != nil{
		log.Println("Submit packet task(client->server) error : " ,err)
	}
}
