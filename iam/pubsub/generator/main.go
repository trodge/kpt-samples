package generator

import (
	"github.com/GoogleContainerTools/kpt-functions-sdk/go/pkg/framework/types"
	"os"
	"sigs.k8s.io/kustomize/kyaml/kio"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

func Generate(configs *types.Configs) error {
	w := &kio.ByteWriter{Writer: os.Stdout}
	var nodes []*yaml.RNode
	w.Write(nodes)
	return nil
}
