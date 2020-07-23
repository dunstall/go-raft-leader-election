package elector

import (
  // "context"
)

type Client interface {
  // DialContext(ctx context.Context, addr string) (Connection, error)
  Dial(addr string) (Connection, error)
}
