package test

import (
	"data"
	"logger"
	"testing"
)

func TestCapturePacket(t *testing.T){
	logger.InitLog()
	adpter := data.NetAdapter{
		DeviceName:"\\Device\\NPF_{CF9F69E9-4D50-4102-BC01-4FFD254C89FF}",
		SnapShotLen:1024,
		Promiscuous:false,
	}
	adpter.CapturePacket()
}
