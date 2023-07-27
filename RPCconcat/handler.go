package main

import (
	api "concat/kitex_gen/api"
	"context"
)

// ConcatImpl implements the last service interface defined in the IDL.
type ConcatImpl struct{}

// Concat implements the ConcatImpl interface.
func (s *ConcatImpl) Concat(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	// TODO: Your code here...
	return
}
