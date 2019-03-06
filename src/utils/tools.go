package utils

import (
	"log"
	"os"
	"path/filepath"
)

func GetAbsPath() string{
	path ,err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil{
		log.Println("get current abs path error : " ,err)
		return ""
	}
	return path
}
