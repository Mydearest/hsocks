package data

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"log"
	"time"
)

type NetAdapter struct {
	DeviceName string	// 物理网卡名
	SnapShotLen int32	// 每个数据包最大大小
	Promiscuous bool	// 是否使用混杂模式
	Timeout	time.Duration	// 超时时间，超过该时间未捕获到数据包则停止捕获
}

func (adapter NetAdapter)CapturePacket(){
	handle, err := pcap.OpenLive(adapter.DeviceName, adapter.SnapShotLen, adapter.Promiscuous, adapter.Timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// 仅代理tcp数据包非tcp全部不管
		if validTcpPacket(packet){
			packetHandle(packet)
		}
	}
}

// 检查是否为tcp格式
// https://godoc.org/github.com/google/gopacket
func validTcpPacket(packet gopacket.Packet) bool {
	if tcpLayer := packet.Layer(layers.LayerTypeTCP);tcpLayer != nil{
		//tcp ,_ := tcpLayer.(*layers.TCP)
		//log.Printf("Capture packet from src port %d to dst port %d \n" ,tcp.SrcPort ,tcp.DstPort)
		tcpLayer.LayerPayload()
		return true
	}
	return false
}
