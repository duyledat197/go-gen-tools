#!/bin/bash

INSFRAS=insfras
SERVICES=services

# kind set up
go install sigs.k8s.io/kind@v0.17.0
kind create cluster --quiet

# add repo
helm repo add bitnami https://charts.bitnami.com/bitnami --no-update
helm repo add jaeger-all-in-one https://raw.githubusercontent.com/hansehe/jaeger-all-in-one/master/helm/charts --no-update
helm repo add hashicorp https://helm.releases.hashicorp.com --no-update
helm repo add grafana https://grafana.github.io/helm-charts --no-update
helm repo add nats https://nats-io.github.io/k8s/helm/charts --no-update
helm repo add minio https://charts.min.io/ --no-update

# install insfras
insfras=("bitnami/postgresql" "jaeger-all-in-one/jaeger-all-in-one" "hashicorp/consul" "grafana/grafana" "bitnami/kafka" "bitnami/redis" "nats/nats" "minio/minio")

for insfra in ${insfras[@]}; do
  name=$(echo ${insfra} | sed -r 's:.*/::')
  helm upgrade --install ${name} ${insfra} --namespace ${INSFRAS} --create-namespace
done

#install services
srvs=("gateway" "order" "inventory" "third_party")
