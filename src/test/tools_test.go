package test

import (
	"log"
	"logger"
	"testing"
	"utils"
)

func TestTools(t *testing.T){
	log.Println(utils.GetAbsPath())
}

func TestLog(t *testing.T){
	logger.SetLog()
	log.Println("1")
}