import grpc
import os
import ping_pb2_grpc
import ping_pb2

channel = grpc.insecure_channel('0.0.0.0:8080', options=(('grpc.enable_http_proxy', 0),))
stub = ping_pb2_grpc.PingServiceStub(channel)

req = ping_pb2.PingRequest(data="1st ping")
res = stub.Ping(req)
print(res)