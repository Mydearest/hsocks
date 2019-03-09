package call

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/rpc"
	"utils"
)

type Proxy struct {
	RpcClient *rpc.Client
}

type ProxyRequest struct {
	Url string
	Host string
	Method string
	Protocol string
	Headers map[string][]string
	Body []byte
}

type ProxyResponse struct {
	Protocol string
	StatusCode int
	Headers map[string][]string
	Body []byte
}


type RpcMethod string

const M_ProxyHttp RpcMethod = "ProxyRpc.ProxyHttp"
const M_ProxySsl RpcMethod = "ProxyRpc.ProxySsl"

var DefaultClientProxy *Proxy

//var DefaultServerProxy *Proxy
//
//func InitServerProxy(){
//	DefaultServerProxy = &Proxy{}
//}

func InitClientProxy(){
	if conn, err := rpc.DialHTTP("tcp", utils.Args.ProxyServer);err != nil{
		log.Fatalln("init rpc client error : " ,err)
	}else{
		DefaultClientProxy = &Proxy{
			RpcClient:conn,
		}
	}
}

func (*Proxy)BuildProxyRequest(req http.Request) ProxyRequest{
	proxyRequest := ProxyRequest{}
	data ,err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil{
		log.Println(err)
		return proxyRequest
	}else {
		proxyRequest = ProxyRequest{
			Method:req.Method,
			Url:req.URL.Path,
			Host:req.Host,
			Protocol:req.Proto,
			Headers:req.Header,
			Body:data,
		}
		return proxyRequest
	}
}

func (p *Proxy)Invoke(method RpcMethod ,req ProxyRequest ,res *ProxyResponse) error{
	return p.RpcClient.Call(string(method) ,req ,res)
}


//___________________________________RPC method___________________________________

type ProxyRpc struct {

}

func (p *ProxyRpc)ProxySsl(){

}

func (p *ProxyRpc)ProxyHttp(req ProxyRequest ,res *ProxyResponse) error{
	httpClient := &HttpClient{
		Url:req.Url,
		Method:req.Method,
		Body:req.Body,
		Protocol:req.Protocol,
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
		res.Protocol = httpRes.Proto
		res.StatusCode = httpRes.StatusCode
		return nil
	}
}