#!/bin/bash

JSONPATH="{.contexts[?(@.name==\"$(kubectl config current-context)\")].context.namespace}"
CURRENT_NAMESPACE=$(kubectl config view -o jsonpath=$JSONPATH)

kubectl create serviceaccount loadbalancer-source-ranger -n ${CURRENT_NAMESPACE}
kubectl create clusterrolebinding --clusterrole=cluster-admin --serviceaccount=${CURRENT_NAMESPACE}:loadbalancer-source-ranger loadbalancer-source-range
kubectl apply -f example/deployment.yaml