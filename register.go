// Package statsd registers the extension for output
package statsd

import (
	"github.com/javaducky/xk6-output-statsd/pkg/statsd"
	"go.k6.io/k6/output"
)

func init() {
	output.RegisterExtension("output-statsd", func(p output.Params) (output.Output, error) {
		return statsd.New(p)
	})
}
