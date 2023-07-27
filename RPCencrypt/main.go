package main

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/kitex/server/genericserver"
	etcd "github.com/kitex-contrib/registry-etcd"
	"io"
	"io/ioutil"
	"net"
	"os"
	"strconv"
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
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"}) // r should not be reused.
	if err != nil {
		panic(err)
	}
	filename := "counter.txt"

	// Read the current value from the file
	data, err := ioutil.ReadFile(filename)
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("Error reading the counter:", err)
		return
	}

	// Parse the current value
	var i int
	if len(data) > 0 {
		i, err = strconv.Atoi(string(data))
		if err != nil {
			fmt.Println("Error parsing the counter:", err)
			return
		}
	}
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", 8894+i))
	// Increment the counter
	i++
	i %= 3

	// Write the new value back to the file
	err = ioutil.WriteFile(filename, []byte(strconv.Itoa(i)), 0644)
	if err != nil {
		fmt.Println("Error writing the counter:", err)
		return
	}

	svr := genericserver.NewServer(new(GenericServiceImpl), g, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "encrypt"}), server.WithServiceAddr(addr), server.WithRegistry(r))
	if err != nil {
		panic(err)
	}

	err = svr.Run()
	if err != nil {
		panic(err)
	}
}

type GenericServiceImpl struct {
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

	fmt.Println(m)
	fmt.Println(jsonRequest)

	dataValue, ok := jsonRequest["message"].(string)
	if !ok {
		fmt.Println("data provided is not a string")
		return
	}
	fmt.Println(dataValue)

	encodedText := base64.StdEncoding.EncodeToString([]byte(dataValue))
	encodedText = base64.StdEncoding.EncodeToString([]byte(encodedText))
	key := []byte("curryisthegoatcurryisthegoatgoat")
	encodedText, err = encrypt(encodedText, key)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Encoded text: %s\n", encodedText)

	jsonRequest["message"] = encodedText

	// var respMap map[string]interface{}

	jsonResponse, err := json.Marshal(jsonRequest)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(jsonResponse))
	// fmt.Println(respMap)

	return string(jsonResponse), nil
}
