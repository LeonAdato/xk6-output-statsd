package statsd

//func TestConfig(t *testing.T) {
//	t.Parallel()
//	// TODO: add more cases
//	testCases := map[string]struct {
//		jsonRaw json.RawMessage
//		env     map[string]string
//		arg     string
//		config  config
//		err     string
//	}{
//		"default": {
//			config: config{
//				Addr:         "statsd",
//				PushInterval: 1 * time.Second,
//			},
//		},
//
//		"overwrite": {
//			env: map[string]string{"K6_TEMPLATE_ADDRESS": "else", "K6_TEMPLATE_PUSH_INTERVAL": "4ms"},
//			config: config{
//				Addr:         "else",
//				PushInterval: 4 * time.Millisecond,
//			},
//		},
//
//		"early error": {
//			env: map[string]string{"K6_TEMPLATE_ADDRESS": "else", "K6_TEMPLATE_PUSH_INTERVAL": "4something"},
//			config: config{
//				Addr:         "else",
//				PushInterval: 1 * time.Second,
//			},
//			err: `time: unknown unit "something" in duration "4something"`,
//		},
//	}
//
//	for name, testCase := range testCases {
//		testCase := testCase
//		t.Run(name, func(t *testing.T) {
//			t.Parallel()
//			config := newConfig()
//			require.Equal(t, testCase.config, config)
//		})
//	}
//}
