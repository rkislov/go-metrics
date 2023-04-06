package entity

import "testing"

func TestMetric_NewMetric(t *testing.T) {
	type testCase struct {
		test        string
		types       string
		name        string
		value       float64
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "emtpy name",
			types:       "",
			name:        "",
			value:       0,
			expectedErr: ErrMissingValues,
		},
		{
			test:        "valid value",
			types:       "type",
			name:        "name",
			value:       1.0,
			expectedErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := NewMetric(tc.name, tc.types, tc.value)
			if err != tc.expectedErr {
				t.Errorf("Expexted error: %v, got %v", tc.expectedErr, err)
			}
		})

	}
}
