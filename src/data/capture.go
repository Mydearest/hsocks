package data

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"log"
)

type NetAdapter struct {
	DeviceName string	// 物理网卡名
	SnapShotLen int32	// 每个数据包最大大小
	Promiscuous bool	// 是否使用混杂模式
	LogAllPacket bool	// 是否log所有packet
}

func (adapter NetAdapter)CapturePacket(){
	handle, err := pcap.OpenLive(adapter.DeviceName, adapter.SnapShotLen, adapter.Promiscuous, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	packets := packetSource.Packets()
	for {
		select {
		case packet := <-packets:
			if packet == nil {
				log.Println("None packet ,check interface")
				return
			}
			if adapter.LogAllPacket {
				log.Println(packet.String())
			}
			if packet.NetworkLayer() == nil || packet.TransportLayer() == nil || packet.TransportLayer().LayerType() != layers.LayerTypeTCP {
				continue
			}
			tcp := packet.TransportLayer().(*layers.TCP)
			// 只处理80和443的包
			if tcp.DstPort != 80 && tcp.DstPort != 443{
				continue
			}
			packetHandle(packet ,tcp)
		}
	}
}
