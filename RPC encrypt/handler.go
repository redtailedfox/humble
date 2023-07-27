package main

import (
	"context"
	api "encrypt/kitex_gen/api"
)

// EncryptImpl implements the last service interface defined in the IDL.
type EncryptImpl struct{}

// Encrypt implements the EncryptImpl interface.
func (s *EncryptImpl) Encrypt(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	// TODO: Your code here...
	return
}
