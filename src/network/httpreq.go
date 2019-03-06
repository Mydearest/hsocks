package network

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

type ProxyRequest struct {
	Url string
	Body []byte
}

func (proxyReq ProxyRequest)Post() (byts []byte ,err error) {
	req ,err := http.NewRequest(http.MethodPost ,proxyReq.Url ,bytes.NewBuffer(proxyReq.Body))
	if err != nil{
		log.Println(err)
	}
	res ,err := http.DefaultClient.Do(req)
	if err != nil{
		log.Println(err)
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}