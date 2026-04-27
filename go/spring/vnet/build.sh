#!/usr/bin/env bash
set -e
docker build --no-cache --platform=linux/amd64 -t saichler/spring-vnet:latest .
docker push saichler/spring-vnet:latest
