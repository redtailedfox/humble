@echo off

REM Navigate to each directory and start each service
start "" cmd /k "cd .\APIGateway\ &&  go get -u github.com/cloudwego/netpoll\"
start "" cmd /k "cd .\RPCconcat\ &&  go get -u github.com/cloudwego/netpoll\"
start "" cmd /k "cd .\RPCdecrypt\ &&  go get -u github.com/cloudwego/netpoll\"
start "" cmd /k "cd .\RPCencrypt\ &&  go get -u github.com/cloudwego/netpoll\"
start "" cmd /k "cd .\RPCServer\ &&  go get -u github.com/cloudwego/netpoll\"