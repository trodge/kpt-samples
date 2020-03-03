package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"sigs.k8s.io/kustomize/kyaml/kio"
	"sigs.k8s.io/kustomize/kyaml/kio/filters"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

const usage = "Generates IAM Policies for multiple pubsub topics."

func main() {
	fc, err := os.Open("fc.yaml")
	if err != nil {
		_ = fmt.Errorf("Error opening file: %v\n", err)
	}
	rw := &kio.ByteReadWriter{
		Reader:                fc,
		Writer:                os.Stdout,
		KeepReaderAnnotations: true,
	}
	err = kio.Pipeline{
		Inputs: []kio.Reader{rw},
		Filters: []kio.Filter{
			&filter{rw: rw},       // generate the Resources from the input
			// set Resource filenames
			&filters.FileSetter{FilenamePattern: filepath.Join("config", "%n.yaml")},
			filters.FormatFilter{}, // format the output
		},
		Outputs: []kio.Writer{rw},
	}.Execute()
	if err != nil {
		_ = fmt.Errorf("Pipeline error: %v\n", err)
		os.Exit(1)
	}
}

type filter struct {
	rw *kio.ByteReadWriter
}


type PubSubIAMFunctionConfig struct {
	Spec struct {
		Project string `yaml:"project"`
		PubSubTopics []string `yaml:"pubSubTopics"`
		Bindings []Binding `yaml:"bindings"`
	}
}

type Binding struct {
	Role string `yaml:"role"`
	Members []string `yaml:"members"`
}

func (f *filter) Filter(in []*yaml.RNode) ([]*yaml.RNode, error) {
	var fc PubSubIAMFunctionConfig
	var out []*yaml.RNode
	for _, i := range in {
		// Read function config from in RNode.
		if err := yaml.Unmarshal([]byte(i.MustString()), &fc); err != nil {
			return nil, fmt.Errorf("Error unmarshalling function config: %v\n", err)
		}
		spec := fc.Spec
		for _, pst := range spec.PubSubTopics {
			for _, bdg := range spec.Bindings {
				for i := range bdg.Members {
					mbr := &bdg.Members[i]
					*mbr = strings.ReplaceAll(*mbr, "${PROJECT_ID?}", spec.Project)
				}
				fmt.Printf("PubSub Topic: %v, Binding: %v\n", pst, bdg)
				out = append(out, nil)
			}
		}
	}
	return in, nil
}
