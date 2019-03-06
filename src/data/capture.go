package data

import (
	"github.com/google/gopacket"
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
		// 处理数据包
		log.Println(packet)
	}
}
