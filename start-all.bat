@echo off
REM Run the etcd command
start cmd "docker run -p 2379:2379 -p 2380:2380 gcr.io/etcd-development/etcd:v3.5.0 /usr/local/bin/etcd --advertise-client-urls http://0.0.0.0:2379 --listen-client-urls http://0.0.0.0:2379"

REM Navigate to each directory and start each service
start cmd "cd APIGateway && go run ."
start cmd "cd RPCconcat && go run ."
start cmd "cd RPCdecrypt && go run ."
start cmd "cd RPCencrypt && go run ."
start cmd "cd RPCServer && go run ."

