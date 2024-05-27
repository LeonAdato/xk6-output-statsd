package statsd

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"gopkg.in/guregu/null.v3"

	"go.k6.io/k6/lib/testutils"
	"go.k6.io/k6/lib/types"
	"go.k6.io/k6/output"
)

func getOutput(
	logger logrus.FieldLogger, addr, namespace null.String, bufferSize null.Int, pushInterval types.NullDuration,
) (*Output, error) {
	return newOutput(
		output.Params{
			Logger: logger,
			JSONConfig: json.RawMessage(fmt.Sprintf(`{
			"addr": "%s",
			"namespace": "%s",
			"bufferSize": %d,
			"pushInterval": "%s"
		}`, addr.String, namespace.String, bufferSize.Int64, pushInterval.Duration.String())),
		})
}




func TestInitWithoutAddressErrors(t *testing.T) {
	t.Parallel()
	c := &Output{
		config: config{},
		logger: testutils.NewLogger(t),
	}
	err := c.Start()
	require.Error(t, err)
}

func TestInitWithBogusAddressErrors(t *testing.T) {
	t.Parallel()
	c := &Output{
		config: config{
			Addr: null.StringFrom("localhost:90000"),
		},
		logger: testutils.NewLogger(t),
	}
	err := c.Start()
	require.Error(t, err)
}

func TestLinkReturnAddress(t *testing.T) {
	t.Parallel()
	bogusValue := "bogus value"
	c := &Output{
		config: config{
			Addr: null.StringFrom(bogusValue),
		},
	}
	require.Equal(t, fmt.Sprintf("statsd (%s)", bogusValue), c.Description())
}
