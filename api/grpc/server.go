package grpc

import (
	pbTodo "reblog-server/api/grpc/todo"

	"reblog-server/service/todo"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcServer struct {
	server *grpcServer
}

func getGRPCServer(svc todo.TodoService) *grpc.Server {
	server := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_prometheus.UnaryServerInterceptor,
			// middleware.MetricInterceptor(cacheRepo),
			// middleware.AddRequestID(),
			// middleware.GRPCLogging(cacheRepo),
			grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandler(func(p interface{}) (err error) {
				return status.Errorf(codes.Unknown, "panic triggered: %v", p)
			})),
		),
	)

	pbTodo.RegisterTodoServiceServer(server, pbTodo.TodoHandler{
		Service: svc,
	})

	pbTodo.RegisterPingServiceServer(server, &pbTodo.PingHandler{})

	// reflection.Register(server)

	return server
}

func NewGRPCServer(svc todo.TodoService) *grpc.Server {
	server := getGRPCServer(svc)
	return server
}
