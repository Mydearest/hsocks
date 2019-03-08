package network

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type HttpClient struct {
	Method string
	Url string
	Body []byte
}

func (client *HttpClient)doRequest() ([]byte ,error){
	if req ,err := http.NewRequest(client.Method ,client.Url ,bytes.NewReader(client.Body));err != nil{
		//log.Printf("Build request error ,m->%s ,url->%s\n" ,client.Method ,client.Method)
		return nil,err
	}else {
		cli := http.DefaultClient
		if res ,err := cli.Do(req);err != nil{
			//log.Printf("Do request error ,m->%s ,url->%s\n" ,client.Method ,client.Method)
			return nil,err
		}else {
			defer res.Body.Close()
			if data ,err := ioutil.ReadAll(res.Body);err != nil{
				//log.Println("Read res errror")
				return nil ,err
			}else{
				return data ,nil
			}
		}
	}
}
