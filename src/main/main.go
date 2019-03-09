package main

import (
	"cli"
	"fmt"
	"logger"
	"ser"
	"utils"
)

func main(){
	//cli.StartSsl()
	args := utils.ParseCommand()
	logger.InitLog()
	if args.Help{
		utils.PrintUsage()
		return
	}else if args.ClientMode == "" && args.ServerMode == ""{
		fmt.Println("select start mode [client|server] ,and set value ,use -h get more help")
		return
	}else if args.ClientMode != "" && args.ServerMode != ""{
		fmt.Println("programs can not run in two modes at the same time")
		return
	}else if args.ClientMode != "" && args.ProxyServer == ""{
		fmt.Println("no proxy server value ,usr -p or -h set proxy server addr")
		return
	}else if args.ClientMode != ""{
		cli.StartProxyClient()
	}else if args.ServerMode != "" || (args.ServerMode == "" && args.CloudFoundry == true){
		ser.StartProxyServer()
	}
}