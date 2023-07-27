#!/bin/bash

# start etcd in the background
etcd --listen-client-urls http://127.0.0.1:2379 --advertise-client-urls http://127.0.0.1:2379 &

# list of subdirectories
subdirs=("APIGateway" "RPCconcat" "RPCdecrypt" "RPCencrypt" "RPCServer")



# iterate over subdirectories and start each main function in the background
for dir in ${subdirs[@]}; do
    echo "Starting $dir"
    (cd $dir && go run .) &
done

# wait for all background processes to finish
wait