package data

import (
	"call"
	"log"
	"time"
	"utils/gopool"
)

var packetReturnExecutor *gopool.FixedExecutor

func init(){
	packetReturnExecutor = gopool.NewFixedExecutor(5 ,8 ,time.Second*5)
	packetReturnExecutor.Start()
}

type PacketReturnTask struct {
	PacketRequest call.PacketRequest
	PacketResponse *call.PacketResponse
}


func AsynHandle(packetReq call.PacketRequest,packetRes *call.PacketResponse){
	task := PacketReturnTask{
		PacketRequest:packetReq,
		PacketResponse:packetRes,
	}
	if err := packetReturnExecutor.Submit(task);err != nil{
		log.Println("Submit return packet task error : " ,err)
	}
}

// ________________________________________________

func (task PacketReturnTask) Run() error {
	//responseBytes := task.PacketResponse.Packet
	return nil
}

func RewritePacket(packet []byte) []byte{
	return nil
}

func SendPacket(packet []byte) error{
	return nil
}
