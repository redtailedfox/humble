package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/kitex/server/genericserver"
	etcd "github.com/kitex-contrib/registry-etcd"
	"net"
)

func main() {
	// Parse IDL with Local Files
	// YOUR_IDL_PATH thrift file path,eg: ./idl/example.thrift
	p, err := generic.NewThriftFileProvider("idl/hello.thrift")
	if err != nil {
		panic(err)
	}
	g, err := generic.JSONThriftGeneric(p)
	if err != nil {
		panic(err)
	}
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"}) // r should not be reused.
	if err != nil {
		panic(err)
	}
	//svr := genericserver.NewServer(new(GenericServiceImpl), g)
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", 8888+1))
	svr := genericserver.NewServer(new(GenericServiceImpl), g, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "call"}), server.WithServiceAddr(addr), server.WithRegistry(r))
	if err != nil {
		panic(err)
	}

	err = svr.Run()
	if err != nil {
		panic(err)
	}
	//for i := 0; i < 3; i++ { // adjust the number of instances as needed
	//	go func(i int) {
	//		addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("hellorpc:%d", 8888+i))
	//		if err != nil {
	//			log.Fatalf("Failed to resolve server address: %v", err)
	//		}
	//
	//		svr := genericserver.NewServer(
	//			new(GenericServiceImpl),
	//			g,
	//			server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "Call"}),
	//			server.WithServiceAddr(addr),
	//			server.WithRegistry(r),
	//		)
	//
	//		if err != nil {
	//			panic(err)
	//		}
	//
	//		err = svr.Run()
	//		if err != nil {
	//			panic(err)
	//		}
	//	}(i)
	//}

	// resp is a JSON string
}

type GenericServiceImpl struct {
}

func (g *GenericServiceImpl) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {
	m := request.(string)
	var jsonRequest map[string]interface{}

	err = json.Unmarshal([]byte(m), &jsonRequest)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println(m)
	fmt.Println(jsonRequest)

	dataValue, ok := jsonRequest["message"].(string)
	if !ok {
		fmt.Println("data provided is not a string")
		return
	}
	fmt.Println(dataValue)

	jsonRequest["message"] = "Hello!, " + dataValue

	// var respMap map[string]interface{}

	jsonResponse, err := json.Marshal(jsonRequest)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(jsonResponse))
	// fmt.Println(respMap)

	return string(jsonResponse), nil
}
