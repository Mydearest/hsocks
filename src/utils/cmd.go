package utils

import (
	"flag"
	"fmt"
)

type CommandArgs struct {
	Help bool
	// local proxy client listening addr
	ClientMode string
	// remote proxy server listening addr
	ServerMode string
	// base on http
	ProxyServer string
	// use cf web service as proxy server
	CloudFoundry bool
	ProxyTimeout int64
	Debug bool	// log.setFlag(log.LLongFile) ,if false -> devNul
}

var Args *CommandArgs

func init(){
	Args = &CommandArgs{}
}

func ParseCommand() *CommandArgs{
	flag.StringVar(&Args.ProxyServer ,"p" ,"" ,"proxy server addr ,start with http(s)")
	flag.StringVar(&Args.ClientMode ,"c" ,"" ,"start with client(local) mode ,--client <LOCAL_PORT>")
	flag.StringVar(&Args.ServerMode ,"s" ,"" ,"start with server mode ,--server <RPC_PORT>")
	flag.BoolVar(&Args.CloudFoundry ,"C" ,false ,"use cloudfoundry as proxy service ,that means the value of --server <RPC_PORT> is useless")
	flag.BoolVar(&Args.Help ,"h" ,false ,"help info")
	flag.BoolVar(&Args.Debug ,"d" ,false ,"debug mode")
	flag.Int64Var(&Args.ProxyTimeout ,"t" ,10 ,"set proxy time out (second) -> local server put req/get res from proxy server ,and if this val < 0 ,then proxy would never time out")
	flag.Parse()
	flag.Usage = PrintUsage
	return Args
}

func PrintUsage(){
	fmt.Println("Default Client-side ,use --server to start with server mode")
	flag.PrintDefaults()
}