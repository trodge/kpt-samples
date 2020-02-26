#!/bin/bash
resources=$(kubectl apply -f $1)
resources=$(cut -d' ' -f1 <<< $resources)
for r in $resources; do
  kind=$(cut -d'.' -f1 <<< $r);
  name=${r##*/};
  echo kubectl describe $kind $name
  kubectl describe $kind $name
done
