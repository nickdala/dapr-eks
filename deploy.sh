#!/usr/bin/env bash

kubectl apply -f kubernetes/redis.yaml
kubectl apply -f kubernetes/statestore.yaml
kubectl apply -f kubernetes/hellodpr.yaml