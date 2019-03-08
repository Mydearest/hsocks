package call

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/rpc"
	"network"
	"utils"
)

type Proxy struct {
	RpcClient *rpc.Client
}

type ProxyRequest struct {
	Url string
	Host string
	Method string
	Protocal string
	Headers map[string][]string
	Body []byte
}

type ProxyResponse struct {
	Protocal string
	StatusCode int
	Headers map[string][]string
	Body []byte
}

var DefaultProxy *Proxy

func init(){
	if conn, err := rpc.DialHTTP("tcp", utils.Args.ProxyServer);err != nil{
		log.Fatalln("init rpc client error : " ,err)
	}else{
		DefaultProxy = &Proxy{
			RpcClient:conn,
		}
	}
}

func (p *Proxy)BuildProxyRequest(req http.Request) ProxyRequest{
	proxyRequst := ProxyRequest{}
	data ,err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil{
		log.Println(err)
		return proxyRequst
	}else {
		proxyRequst = ProxyRequest{
			Method:req.Method,
			Url:req.URL.Path,
			Host:req.Host,
			Protocal:req.Proto,
			Headers:req.Header,
			Body:data,
		}
		return proxyRequst
	}
}

func (p *Proxy)Invoke(method RpcMethod ,req ProxyRequest ,res *ProxyResponse) error{
	return p.RpcClient.Call(string(method) ,req ,res)
}

type RpcMethod string

const ProxyHttp RpcMethod = "Proxy.ProxyHttp"

func (p *Proxy)ProxyHttp(req ProxyRequest ,res *ProxyResponse) error{
	httpClient := &network.HttpClient{
		Url:req.Url,
		Method:req.Method,
		Body:req.Body,
		Protocal:req.Protocal,
		Host:req.Host,
		Headers:req.Headers,
	}
	if httpRes ,err := httpClient.DoRequest();err != nil{
		return err
	}else{
		data ,err := ioutil.ReadAll(httpRes.Body)
		defer httpRes.Body.Close()
		if err != nil{
			return err
		}
		res.Body = data
		res.Headers = httpRes.Header
		res.Protocal = httpRes.Proto
		res.StatusCode = httpRes.StatusCode
		return nil
	}
}