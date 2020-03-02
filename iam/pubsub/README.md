# Kpt Function Usage

For the purposes of this sample we will assume successful installation of KCC in workload identity mode.


Currently, configurations are in fc.yaml. Only valid k8s objects are supported. User will need to supply a list
of pubsub topics and a map of IAM policy bindings. Data such as project id and PubSub topics list can be specified under
the "spec" object.