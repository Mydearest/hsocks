package main

import (
	"fmt"
	"logger"
	"testing"
	"utils"
)


func init(){
	logger.InitLog()
}

func TestMainn(t *testing.T){
	args := utils.ParseCommand()
	if args.Help{
		utils.PrintUsage()
		return
	}else if args.ProxyServer == ""{
		fmt.Println("no proxy server ,use -p")
		return
	}
}
