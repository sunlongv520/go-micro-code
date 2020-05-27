set MICRO_REGISTRY=etcd
set MICRO_REGISTRY_ADDRESS=192.168.2.232:2379
set MICRO_CLIENT=grpc
set MICRO_SERVER=grpc
set MICRO_API_NAMESPACE=api.jtthink.com
micro api --handler=api
