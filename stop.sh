#!/bin/bash

echo "------------------------------gate-api----------------------------------------"
pkill -f gate-api

echo "------------------------------rolemanager-rpc----------------------------------------"
pkill -f rolemanager-rpc

echo "------------------------------usercenter-rpc----------------------------------------"
pkill -f usercenter-rpc

echo "------------------------------servermanager-rpc----------------------------------------"
pkill -f servermanager-rpc
