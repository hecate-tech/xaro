#!/bin/sh

protoc -I=protobuf --go_out=src/proto protobuf/engo-xaro.proto 
pbjs protobuf/engo-xaro.proto > assets/json/superstellar_proto.json