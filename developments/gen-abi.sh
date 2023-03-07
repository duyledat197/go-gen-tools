#!/bin/sh

ABI_PATH=/app/api/contracts

for f in ${ABI_PATH}/*.sol; do
  name=$(echo ${f} | sed -r 's/\.[^.]*$//')
  echo $name
  echo $f
  abigen --abi $f --pkg abi --type $name --out $name.go

done
