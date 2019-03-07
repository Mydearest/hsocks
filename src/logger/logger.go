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

func InitLog(){
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	setLogWriter()
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
