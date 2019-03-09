package call

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

type HttpClient struct {
	Url string
	Host string
	Method string
	Protocol string
	Headers map[string][]string
	Body []byte
	SSL bool
}

func (client *HttpClient)DoRequest() (*http.Response ,error){
	trueUrl := fmt.Sprintf("http://%s%s" ,client.Host ,client.Url)
	if req ,err := http.NewRequest(client.Method ,trueUrl ,bytes.NewReader(client.Body));err != nil{
		log.Printf("Build request error ,m->%s ,url->%s\n" ,client.Method ,trueUrl)
		return nil ,err
	}else {
		req.Header = client.Headers
		req.Proto = client.Protocol
		req.Host = client.Host
		cli := http.DefaultClient
		if res ,err := cli.Do(req);err != nil{
			log.Printf("Do request error ,m->%s ,url->%s\n" ,client.Method ,trueUrl)
			return nil ,err
		}else {
			return res ,nil
		}
	}
}
