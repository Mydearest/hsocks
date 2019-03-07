package utils

import (
	"flag"
	"fmt"
)

type CommandArgs struct {
	Help bool
	ProxyServer string	// http(s)

}

var Args *CommandArgs

func init(){
	Args = &CommandArgs{}
}

func ParseCommand() *CommandArgs{
	flag.StringVar(&Args.ProxyServer ,"p" ,"" ,"proxy server addr ,start with http(s)")
	flag.BoolVar(&Args.Help ,"h" ,false ,"help info")
	flag.Parse()
	flag.Usage = PrintUsage
	return Args
}

func PrintUsage(){
	fmt.Println("Client-side")
	flag.PrintDefaults()
}