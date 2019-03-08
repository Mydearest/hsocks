package network

import (
	"bytes"
	"log"
	"net/http"
)

type HttpClient struct {
	Url string
	Host string
	Method string
	Protocal string
	Headers map[string][]string
	Body []byte
}

func (client *HttpClient)DoRequest() (*http.Response ,error){
	if req ,err := http.NewRequest(client.Method ,client.Url ,bytes.NewReader(client.Body));err != nil{
		log.Printf("Build request error ,m->%s ,url->%s\n" ,client.Method ,client.Method)
		return nil ,err
	}else {
		req.Header = client.Headers
		req.Proto = client.Protocal
		req.Host = client.Host
		cli := http.DefaultClient
		if res ,err := cli.Do(req);err != nil{
			log.Printf("Do request error ,m->%s ,url->%s\n" ,client.Method ,client.Method)
			return nil ,err
		}else {
			return res ,nil
		}
	}
}
