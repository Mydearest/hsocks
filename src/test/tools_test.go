package test

import (
	"log"
	"logger"
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
	//req := data.ProxyRequest{
	//	Url:"http://static.shinoha.cn/upload",
	//	Body:[]byte("q=123"),
	//}
	//byts ,_ := req.Post()
	//log.Println(string(byts))
}

func TestMainn(t *testing.T){
	args := utils.ParseCommand()
	if args.Help{
		utils.PrintUsage()
		return
	}else if args.ProxyServer == ""{
		log.Println("no proxy server ,use -p")
		return
	}
}