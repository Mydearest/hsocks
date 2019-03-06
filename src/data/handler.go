package data

import (
	"github.com/google/gopacket"
	"log"
	"utils/gopool"
)

func init(){
	packetSendExecutor = gopool.NewSingleExecutor()
	packetSendExecutor.Start()
}

type packetSendTask struct {
	tcpPacket gopacket.Packet
}

func (task packetSendTask) Run() error {

	return nil
}

var packetSendExecutor *gopool.SingleExecutor

func packetHandle(packet gopacket.Packet){
	task := packetSendTask{
		tcpPacket:packet,
	}
	if err := packetSendExecutor.Submit(task);err != nil{
		log.Println("Submit packet task(client->server) error : " ,err)
	}
}
