Pubsub
==================================================

# NAME

  Config Connector Pubsub

# SYNOPSIS

  To apply pull subscription:

  `kubectl apply -f topic.yaml`

  `kubectl apply -f subscription.yaml`

  To apply push subscription:

  `kubectl apply -f topic.yaml`

  `kustomize build . | kubectl apply -f -`

# Description

  This package contains YAML files for management of a single PubSub topic and
  pull subscription through Config Connector. To build a push subscription YAML
  run `kustomize build .`

  kustomize can be found at https://github.com/kubernetes-sigs/kustomize

