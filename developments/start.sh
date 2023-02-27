#!/bin/bash

INSFRAS=insfras
SERVICES=services

# kind set up
go install sigs.k8s.io/kind@v0.17.0
kind create cluster --quiet

# add repo
helm repo add bitnami https://charts.bitnami.com/bitnami --force-update
helm repo add jaeger-all-in-one https://raw.githubusercontent.com/hansehe/jaeger-all-in-one/master/helm/charts --force-update
helm repo add hashicorp https://helm.releases.hashicorp.com --force-update
helm repo add grafana https://grafana.github.io/helm-charts --force-update
helm repo add nats https://nats-io.github.io/k8s/helm/charts --force-update
helm repo add minio https://charts.min.io/ --force-update

# install insfras
helm upgrade --install postgresql bitnami/postgresql --namespace ${INSFRAS} --create-namespace
helm upgrade --install jaeger-all-in-one jaeger-all-in-one/jaeger-all-in-one --namespace ${INSFRAS} --create-namespace
helm upgrade --install consul hashicorp/consul --namespace ${INSFRAS} --create-namespace
helm upgrade --install grafana grafana/grafana --namespace ${INSFRAS} --create-namespace
helm upgrade --install kafka bitnami/kafka --namespace ${INSFRAS} --create-namespace
helm upgrade --install redis bitnami/redis --namespace ${INSFRAS} --create-namespace
helm upgrade --install nats nats/nats --namespace ${INSFRAS} --create-namespace
helm upgrade --install minio minio/minio --namespace ${INSFRAS} --create-namespace

#install services
