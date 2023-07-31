// Package template registers the extension for output
package template

import (
	"github.com/grafana/xk6-output-template/pkg/template"
	"go.k6.io/k6/output"
)

func init() {
	output.RegisterExtension("xk6-template", func(p output.Params) (output.Output, error) {
		return template.New(p)
	})
}
