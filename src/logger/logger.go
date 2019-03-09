package logger

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"utils"
)

const logFile = "hsocks.log"
const logDir = "log"

var Debug *log.Logger

type devNul struct {

}

func (devNul) Write(p []byte) (n int, err error) {
	return 0 ,nil
}

func InitLog(){
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	setLogWriter()
	if utils.Args.Debug{
		Debug = log.New(os.Stderr ,"Debug > " ,log.Llongfile)
	}else{
		nul := devNul{}
		Debug = log.New(nul ,"" ,0)
	}
}

func setLogWriter() {
	dir := filepath.Join(utils.GetAbsPath() ,logDir)
	_ = os.Mkdir(dir, 755)
	filePath := filepath.Join(dir ,logFile)
	if f ,err := os.OpenFile(filePath ,os.O_WRONLY | os.O_CREATE | os.O_APPEND,755);err != nil{
		log.Fatalln("init log file failed : " ,err)
	}else {
		log.SetOutput(io.MultiWriter(f ,os.Stderr))
	}
}
