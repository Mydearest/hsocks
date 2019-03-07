package utils

import (
	"flag"
	"fmt"
	"time"
)

type CommandArgs struct {
	Help bool
	ProxyServer string	// domain <- cloudfoundry
	ProxyTimeout int64
}

var Args *CommandArgs

func init(){
	Args = &CommandArgs{}
}

func ParseCommand() *CommandArgs{
	flag.StringVar(&Args.ProxyServer ,"p" ,"" ,"proxy server addr ,start with http(s)")
	flag.BoolVar(&Args.Help ,"h" ,false ,"help info")
	flag.Int64Var(&Args.ProxyTimeout ,"t" ,int64(time.Second*5) ,"set proxy time out -> local server put req/get res from proxy server")
	flag.Parse()
	flag.Usage = PrintUsage
	return Args
}

func PrintUsage(){
	fmt.Println("Client-side")
	flag.PrintDefaults()
}