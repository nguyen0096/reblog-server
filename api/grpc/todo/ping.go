package todo

import (
	"context"
	"log"
)

type PingHandler struct{}

func (c *PingHandler) Ping(ctx context.Context, in *PingRequest) (*PingResponse, error) {
	log.Printf("Received: %v", in.Data)
	return &PingResponse{Data: "Data: " + in.Data}, nil
}
