package main

import (
	"fmt"

	"context"
	api "hello/kitex_gen/api"
)

// EchoImpl implements the last service interface defined in the IDL.
type EchoImpl struct{}

// Call implements the EchoImpl interface.
func (s *EchoImpl) Call(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	fmt.Println(resp)
	fmt.Println(req)
	return &api.Response{Message: "hello, hong wei"}, nil
}
