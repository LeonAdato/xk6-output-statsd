package statsd

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"go.k6.io/k6/lib/types"
	"go.k6.io/k6/metrics"
	"gopkg.in/guregu/null.v3"
	"testing"
	"time"
)

func TestConfig(t *testing.T) {
	t.Parallel()
	testCases := map[string]struct {
		jsonRaw json.RawMessage
		env     map[string]string
		arg     string
		config  config
		err     string
	}{
		"default": {
			config: config{
				Addr:         null.NewString("localhost:8125", false),
				BufferSize:   null.NewInt(20, false),
				Namespace:    null.NewString("k6.", false),
				PushInterval: types.NewNullDuration(1*time.Second, false),
				TagBlocklist: metrics.SystemTagSet(metrics.TagVU | metrics.TagIter | metrics.TagURL).Map(),
				EnableTags:   null.NewBool(false, false),
			},
		},
		"overwrite-with-env": {
			env: map[string]string{
				"K6_STATSD_ADDR":          "override:8125",
				"K6_STATSD_BUFFER_SIZE":   "1000",
				"K6_STATSD_NAMESPACE":     "TEST.",
				"K6_STATSD_PUSH_INTERVAL": "100ms",
				"K6_STATSD_TAG_BLOCKLIST": "method,group",
				"K6_STATSD_ENABLE_TAGS":   "true",
			},
			config: config{
				Addr:         null.NewString("override:8125", true),
				BufferSize:   null.NewInt(1000, true),
				Namespace:    null.NewString("TEST.", true),
				PushInterval: types.NewNullDuration(100*time.Millisecond, true),
				TagBlocklist: metrics.SystemTagSet(metrics.TagMethod | metrics.TagGroup).Map(),
				EnableTags:   null.NewBool(true, true),
			},
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			config, err := getConsolidatedConfig(testCase.jsonRaw, testCase.env, testCase.arg)
			if testCase.err != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), testCase.err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, testCase.config, config)
		})
	}
}
