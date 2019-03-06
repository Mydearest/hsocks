package data

import (
	"github.com/google/gopacket/pcap"
	"log"
)

func AllInterface() []pcap.Interface{
	devices, err := pcap.FindAllDevs()
	if err != nil{
		log.Println(err)
	}
	return devices
}
