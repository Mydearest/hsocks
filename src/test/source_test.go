package test

import (
	"data"
	"testing"
	"time"
)

func TestCapturePacket(t *testing.T){
	adpter := data.NetAdapter{
		DeviceName:"\\Device\\NPF_{CF9F69E9-4D50-4102-BC01-4FFD254C89FF}",
		SnapShotLen:1024,
		Timeout:time.Second*20,
		Promiscuous:false,
	}
	adpter.CapturePacket()
}
