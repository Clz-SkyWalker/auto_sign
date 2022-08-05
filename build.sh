#!/bin/bash

if [ $1 == "build" ]
then
  go build -o ./bin/auto_sign ./cmd/main.go
fi
