package main

import (
	consts "concat/constants"
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/kitex/server/genericserver"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func main() {
	// Parse IDL with Local Files
	// YOUR_IDL_PATH thrift file path,eg: ./idl/example.thrift
	p, err := generic.NewThriftFileProvider("idl/concat.thrift")
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
		addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", consts.Concat+i))
		if err != nil {
			log.Fatalf("Failed to resolve server address: %v", err)
		}

		svr := genericserver.NewServer(
			new(GenericServiceImpl),
			g,
			server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "concat"}),
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

	first, ok := jsonRequest["message1"].(string)
	if !ok {
		fmt.Println("wrong type")
	}

	second, ok := jsonRequest["message2"].(string)
	if !ok {
		fmt.Println("wrong type")
	}

	jsonRequest["message"] = first + second

	jsonResponse, err := json.Marshal(jsonRequest)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(jsonResponse))

	return string(jsonResponse), nil
}
