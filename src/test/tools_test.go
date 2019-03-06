package test

import (
	"log"
	"logger"
	"network"
	"testing"
	"utils"
)

func TestTools(t *testing.T){
	log.Println(utils.GetAbsPath())
}

func TestLog(t *testing.T){
	logger.InitLog()
	log.Println("1")
}

func TestPost(t *testing.T){
	req := network.ProxyRequest{
		Url:"http://static.shinoha.cn/upload",
		Body:[]byte("q=123"),
	}
	byts ,_ := req.Post()
	log.Println(string(byts))
}