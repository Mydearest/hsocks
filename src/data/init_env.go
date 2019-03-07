package data

import (
	"github.com/google/gopacket/pcap"
	"log"
)


func GetAllInterface() []pcap.Interface{
	devices, err := pcap.FindAllDevs()
	if err != nil{
		log.Println(err)
	}
	return devices
}
