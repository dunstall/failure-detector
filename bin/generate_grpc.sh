#!/bin/bash
protoc -I pb/ pb/rpc.proto --go_out=plugins=grpc:pb
