
@echo off
REM Run the etcd command
start "" cmd /k "docker run -p 2379:2379 -p 2380:2380 gcr.io/etcd-development/etcd:v3.5.0 /usr/local/bin/etcd --advertise-client-urls http://0.0.0.0:2379 --listen-client-urls http://0.0.0.0:2379"

REM Navigate to each directory and start each service
start "" cmd /k "cd .\APIGateway\ && go mod tidy && go run ."
start "" cmd /k "cd .\RPCconcat\ && go mod tidy && go run ."
start "" cmd /k "cd .\RPCdecrypt\ && go mod tidy && go run ."
start "" cmd /k "cd .\RPCencrypt\ && go mod tidy && go run ."
start "" cmd /k "cd .\RPCServer\ && go mod tidy && go run ."