package main

import (
	"github.com/GoogleContainerTools/kpt-functions-sdk/go/pkg/framework/runners"
	"kpt-samples/iam/pubsub/generator"
)

const usage = "Generates IAM Policies for multiple pubsub topics."

func main() {
	runners.RunFunc(generator.Generate, usage)
}
