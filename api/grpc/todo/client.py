import grpc
import todo_pb2_grpc
import todo_pb2


channel = grpc.insecure_channel('172.25.160.1:8080', options=(('grpc.enable_http_proxy', 0),))
stub = todo_pb2_grpc.TodoServiceStub(channel)

item = todo_pb2.TodoItem(title="title 1", short_description="short desc 1", description="desc 1", created_by="created by 1")
res = stub.AddTodo(item)
print(res)