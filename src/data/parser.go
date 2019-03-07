package data

import (
	"errors"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func ParseData(data []byte) gopacket.Packet{
	return gopacket.NewPacket(data ,layers.LayerTypeEthernet ,gopacket.Default)
}

func ParseIpv4(packet gopacket.Packet) (*layers.IPv4  ,error){
	if ipv4 := packet.Layer(layers.LayerTypeIPv4);ipv4 == nil{
		return nil ,errors.New("error in parse packet to ipv4")
	}else{
		tcp ,_ := ipv4.(*layers.IPv4)
		return tcp ,nil
	}
}