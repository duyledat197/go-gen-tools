#!/bin/bash

go install sigs.k8s.io/kind@v0.17.0
kind create cluster --quiet

helm repo add bitnami https://charts.bitnami.com/bitnami --no-update --namespace insfras

helm upgrade --install postgresql bitnami/postgresql
