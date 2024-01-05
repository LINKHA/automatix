#!/bin/bash

echo "------------------------------gate-api----------------------------------------"
# go build -o data/server/gate-api  -v app/gate/cmd/api/gate.go
# ./data/server/gate-api -f app/gate/cmd/api/etc/gate.yaml &

echo "------------------------------rolemanager-rpc----------------------------------------"
# go build -o data/server/rolemanager-rpc  -v app/rolemanager/cmd/rpc/rolemanager.go
# ./data/server/rolemanager-rpc -f app/rolemanager/cmd/rpc/etc/rolemanager.yaml &

echo "------------------------------usercenter-rpc----------------------------------------"
go build -o data/server/usercenter-rpc  -v app/usercenter/cmd/rpc/usercenter.go
./data/server/usercenter-rpc -f app/usercenter/cmd/rpc/etc/usercenter.yaml &

echo "------------------------------servermanager-rpc----------------------------------------"
go build -o data/server/servermanager-rpc  -v app/servermanager/cmd/rpc/servermanager.go
./data/server/servermanager-rpc -f app/servermanager/cmd/rpc/etc/servermanager.yaml &

echo "------------------------------roommanager-rpc----------------------------------------"
go build -o data/server/roommanager-rpc  -v app/roommanager/cmd/rpc/roommanager.go
./data/server/roommanager-rpc -f app/roommanager/cmd/rpc/etc/roommanager.yaml &

wait