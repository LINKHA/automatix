#!/bin/bash
go build --trimpath --mod=vendor --buildmode=plugin -o ./data/modules/folders/backend.so
../nakama --config ./nakama-config.yml --database.address postgres:localdb@localhost:5432/nakama