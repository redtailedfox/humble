package main

import (
	"context"
	api "decrypt/kitex_gen/api"
)

// DecryptImpl implements the last service interface defined in the IDL.
type DecryptImpl struct{}

// Decrypt implements the DecryptImpl interface.
func (s *DecryptImpl) Decrypt(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	// TODO: Your code here...
	return
}
