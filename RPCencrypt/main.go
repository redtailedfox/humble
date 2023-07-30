package main

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	consts "encrypt/constants"
	"fmt"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/kitex/server/genericserver"
	etcd "github.com/kitex-contrib/registry-etcd"
	"io"
	"log"
	"net"
)

func main() {
	p, err := generic.NewThriftFileProvider("idl/encrypt.thrift")
	if err != nil {
		panic(err)
	}
	g, err := generic.JSONThriftGeneric(p)
	if err != nil {
		panic(err)
	}
	r, err := etcd.NewEtcdRegistry([]string{consts.EtcdAddr}) // r should not be reused.
	if err != nil {
		log.Fatalf("Failed to create etcd registry: %v", err)
	}
	servers := make([]server.Server, consts.NumServers)

	for i := 0; i < consts.NumServers; i++ {
		addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", consts.Encrypt+i))
		if err != nil {
			log.Fatalf("Failed to resolve server address: %v", err)
		}

		impl := &GenericServiceImpl{ServerName: fmt.Sprintf("encrypt%d", i)}
		svr := genericserver.NewServer(
			impl,
			g,
			server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "encrypt"}),
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
	ServerName string
}

func encrypt(stringToEncrypt string, key []byte) (string, error) {
	plaintext := []byte(stringToEncrypt)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	return base64.URLEncoding.EncodeToString(ciphertext), nil
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

	encodedText := base64.StdEncoding.EncodeToString([]byte(dataValue))
	encodedText = base64.StdEncoding.EncodeToString([]byte(encodedText))
	key := []byte("curryisthegoatcurryisthegoatgoat")
	encodedText, err = encrypt(encodedText, key)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Encoded text: %s\n", encodedText)

	jsonRequest["message"] = encodedText

	jsonResponse, err := json.Marshal(jsonRequest)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(jsonResponse))

	return string(jsonResponse), nil
}
