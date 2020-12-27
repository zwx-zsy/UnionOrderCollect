#!/usr/bin/env bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor ../main.go
docker build -t registry.cn-shanghai.aliyuncs.com/vincent-k8s/unionordercollectCron:0.0.1 .
docker login --username=vincent321x@gmail.com registry.cn-shanghai.aliyuncs.com
docker push registry.cn-shanghai.aliyuncs.com/vincent-k8s/unionordercollectCron:0.0.1
docker rmi $(docker images registry.cn-shanghai.aliyuncs.com/vincent-k8s/unionordercollectCron)
