package main

import (
	"BitcoinConnect/entity"
	"BitcoinConnect/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)
//常量
const RPCURL = "http://127.0.0.1:7001"
const RPCUSER  = "user"
const RPCPASSWORD = "pwd"
/**
 *1.准备要进行rpc通信时的json数据
 *2.使用http连接的post请求，发送json数据
 *3.接收http响应
 *4.根据响应的结果，进行判断处理
 */
func main() {
	//fmt.Println("Hello Word")
	//1.准备发送的json数据
	/**
	 {id,method,jsonrpc,params}
	 */
	//json序列化，反序列化
	rpcReq :=  entity.RPCRequest{}
	rpcReq.Id = time.Now().Unix()
	rpcReq.Jsonrpc = "2.0"
	rpcReq.Method = "getblockcount"//获取当前节点区块的数量
	//对结构体类型进行序列化
	rpcBytes,err :=json.Marshal(&rpcReq)
    if err != nil{
    	fmt.Println(err.Error())
		return
	}
	//fmt.Println("准备好的json格式的数据:",string(rpcBytes))
     //2.发送一个post请求
	//client:客户端
	client := http.Client{}//实例化一个请求客户端

	//实例化一个请求
	request,err := http.NewRequest("POST",RPCURL,bytes.NewBuffer(rpcBytes))
    if err!= nil{
    	fmt.Println(err.Error())
		return
	}

	//设置请求头
	request.Header.Add("Encoding","UTF-8")
	request.Header.Add("Content-Type","application/json")
	//权限认证
	request.Header.Add("Authorization","Basic "+utils.Base64Str(RPCUSER + ":" + RPCPASSWORD))

	//使用客户端发送请求
	response,err :=client.Do(request)
	if err !=nil{
		fmt.Println(err.Error())
		return
	}

	//通过response获取响应的数据
	code := response.StatusCode
	if code == 200{
		fmt.Println("请求成功")
	}else {
		fmt.Println("请求失败")
	}
}
//https://github.com/ZJ-p007/BcRPCCode.git