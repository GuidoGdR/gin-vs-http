package adapter

import (
	"context"
)

type Handler func(ctx context.Context, req *Request) (*Response, error)
