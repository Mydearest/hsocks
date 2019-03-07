package test

import (
	"data"
	"errors"
	"fmt"
	"log"
	"logger"
	"testing"
	"time"
	"utils/gopool"
)

func TestAllIface(t *testing.T){
	for k ,v := range data.GetAllInterface(){
		log.Println(k ,v)
	}
}

//  client start
func TestCapturePacket(t *testing.T){
	logger.InitLog()
	adpter := data.NetAdapter{
		DeviceName:"{\\Device\\NPF_{634C9C46-2E60-42A4-BED9-096842ADB213}",
		//DeviceName:"\\Device\\NPF_{CF9F69E9-4D50-4102-BC01-4FFD254C89FF}",
		SnapShotLen:1024,
		Promiscuous:false,
	}
	adpter.CapturePacket()
}

type T int

var c = 0

func (T) Run() error {
	c++
	time.Sleep(time.Second)
	return errors.New(fmt.Sprintf("error %d",c))
}

func TestFixedPool(t *testing.T){
	executor := gopool.NewFixedExecutor(1 ,8 ,time.Second*2)
	executor.Start()
	tt := T(10)
	for i:=0;i<100 ;i++  {
		if err := executor.Submit(tt);err != nil{
			log.Println(err)
		}
	}
	time.Sleep(time.Second*5)
}
