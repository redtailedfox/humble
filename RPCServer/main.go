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
	consts "hello/constants"
	"log"
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
	r, err := etcd.NewEtcdRegistry([]string{consts.EtcdAddr}) // r should not be reused.
	if err != nil {
		panic(err)
	}
	servers := make([]server.Server, consts.NumServers)

	for i := 0; i < consts.NumServers; i++ {
		addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", consts.ServerAddr+i))
		if err != nil {
			log.Fatalf("Failed to resolve server address: %v", err)
		}

		svr := genericserver.NewServer(
			new(GenericServiceImpl),
			g,
			server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "call"}),
			server.WithServiceAddr(addr),
			server.WithRegistry(r),
		)

		if err != nil {
			panic(err)
		}

		servers[i] = svr
	}

	// Start all the servers
	for i := 0; i < consts.NumServers; i++ {
		go func(svr server.Server) {
			err := svr.Run()
			if err != nil {
				log.Fatalf("Failed to start server: %v", err)
			}
		}(servers[i])
	}
	select {}
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

	dataValue, ok := jsonRequest["message"].(string)
	if !ok {
		fmt.Println("wrong type")
		return
	}
	fmt.Println(dataValue)

	jsonRequest["message"] = "Echoed " + dataValue

	jsonResponse, err := json.Marshal(jsonRequest)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(jsonResponse))

	return string(jsonResponse), nil
}
