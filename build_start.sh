#!/bin/bash
go build -trimpath -mod=vendor
cd example/
go build --trimpath --mod=vendor --buildmode=plugin -o ./data/modules/folders/backend.so
../automatix --config ./nakama-config.yml --database.address postgres:localdb@localhost:5432/nakama