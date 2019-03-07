package main

// import github.com/google/gopacket || go get github.com/google/gopacket
import (
	"fmt"
	"logger"
	"utils"
)

func init(){
	logger.InitLog()
}

func main(){
	args := utils.ParseCommand()
	if args.Help{
		utils.PrintUsage()
		return
	}else if args.ProxyServer == ""{
		fmt.Println("no proxy server ,use -p")
		return
	}
}