#!/bin/bash

kubectl delete -f https://download.elastic.co/downloads/eck/2.9.0/crds.yaml
kubectl delete -f https://download.elastic.co/downloads/eck/2.9.0/operator.yaml
kubectl delete -f ./elasticsearch.yaml
kubectl delete -f ./kibana.yaml
