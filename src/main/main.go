package main

import (
	"log"
	"os"
)

func init(){
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(os.Stderr)
}

func main(){

}
