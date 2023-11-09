#!/bin/bash

kubectl create -f https://download.elastic.co/downloads/eck/2.9.0/crds.yaml

kubectl apply -f https://download.elastic.co/downloads/eck/2.9.0/operator.yaml

kubectl apply -f ./elasticsearch.yaml -n elastic-system
# Check if the pods are ready
while true; do
    STATUS=$(kubectl get pods -n elastic-system | grep 'es' | awk '{print $2}')
    if [ "$STATUS" == "1/1" ]; then
        echo "Pods are ready."
        break
    else
        echo "Waiting for pods to be ready..."
        sleep 5
    fi
done

kubectl apply -f ./kibana.yaml -n elastic-system
# Check if the pods are ready
while true; do
    STATUS=$(kubectl get pods -n elastic-system | grep 'kb' | awk '{print $2}')
    if [ "$STATUS" == "1/1" ]; then
        echo "Pods are ready."
        break
    else
        echo "Waiting for pods to be ready..."
        sleep 5
    fi
done
