loadbalancer-source-ranger
===============================

# What are admission webhooks??

Let’s take a look at the definition of admission webhooks as it appears in the official Kubernetes documentation. 

>Admission webhooks are HTTP callbacks that receive admission requests and do something with them. You can define two types of admission webhooks, validating admission Webhook and mutating admission webhook. With validating admission Webhooks, you may reject requests to enforce custom admission policies. With mutating admission Webhooks, you may change requests to enforce custom defaults.

# What is `load-banalcer-source-ranger` Admission webhook ?

This webhook automatically sets `loadBalancerSourceRanges` to Service.

# How to build?

## build binary
```
./make.sh
```

## build docker image

# How to deploy？

```
# deploy webhook
kubectl apply -f example/serviceaccount.yaml
kubectl apply -f example/clusterrolebinding.yaml
kubectl apply -f example/deployment.yaml

# apply Service
kubectl apply -f example/sample-service.yaml
```

# TODO
Parameterize IP souces...
