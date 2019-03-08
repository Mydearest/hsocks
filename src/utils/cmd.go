package utils

import (
	"flag"
	"fmt"
)

type CommandArgs struct {
	Help bool
	ClientMode string
	ServerMode string
	ProxyServer string	// domain <- cloudfoundry
	CloudFoundry bool	// use cf web service as proxy server
	ProxyTimeout int64
}

var Args *CommandArgs

func init(){
	Args = &CommandArgs{}
}

func ParseCommand() *CommandArgs{
	flag.StringVar(&Args.ProxyServer ,"p" ,"" ,"proxy server addr ,start with http(s)")
	flag.StringVar(&Args.ClientMode ,"-client" ,":8080" ,"start with client(local) mode ,--client <LOCAL_ADDR>")
	flag.StringVar(&Args.ServerMode ,"-server" ,"" ,"start with server mode ,--server <RPC_PORT>")
	flag.BoolVar(&Args.CloudFoundry ,"c" ,true ,"use cloudfoundry as proxy service ,that means the value of --server <RPC_PORT> is useless")
	flag.BoolVar(&Args.Help ,"h" ,false ,"help info")
	flag.Int64Var(&Args.ProxyTimeout ,"t" ,-1 ,"set proxy time out -> local server put req/get res from proxy server ,and if this val < 0 ,then proxy would never time out")
	flag.Parse()
	flag.Usage = PrintUsage
	return Args
}

func PrintUsage(){
	fmt.Println("Proxy Client-side")
	flag.PrintDefaults()
}