package main

import (
	"fmt"
	"os"
	"path/filepath"

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
	b, _ = ioutil.ReadFile("fc.yaml")
	
	rw := &kio.ByteReadWriter{
		Reader:                fc,
		Writer:                os.Stdout,
		KeepReaderAnnotations: true,
		FunctionConfig: yaml.MustParse(string(b)),
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
	if err := yaml.Unmarshal([]byte(f.rw.FunctionConfig.MustString()), &fc); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	for _, i := range in {
		// Generate your objects here
	}
	return in, nil
}
