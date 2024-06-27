// Unit tests for the adder package
package adder_test

import (
	"testing"

	"github.com/hannahpullen/actions-demo/pkg/adder"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

func TestAdderPositive(t *testing.T) {
	testCases := []struct {
		name           string
		inputs         []string
		expectedResult float64
	}{
		{
			name:           "two_ints",
			inputs:         []string{"1", "2"},
			expectedResult: 3.0,
		},
		{
			name:           "three_ints",
			inputs:         []string{"1", "1", "1"},
			expectedResult: 3.0,
		},
		{
			name:           "decimal_input",
			inputs:         []string{"0.5", "1.1"},
			expectedResult: 1.6,
		},
		{
			name:           "negative_input",
			inputs:         []string{"10.0", "-20.0"},
			expectedResult: -10.0,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			adder := adder.New(zaptest.NewLogger(tt))
			result, err := adder.Add(tc.inputs...)
			require.NoError(tt, err, "error running Add")
			assert.InDelta(tt, tc.expectedResult, result, 1e-4, "Unexpected reuslt from Add")
		})
	}
}

func TestAdderNegative(t *testing.T) {
	testCases := []struct {
		name   string
		inputs []string
	}{
		{
			name:   "non_number_inputs",
			inputs: []string{"a", "b"},
		},
		{
			name:   "mixed_inputs",
			inputs: []string{"4", "b"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			adder := adder.New(zaptest.NewLogger(tt))
			_, err := adder.Add(tc.inputs...)
			assert.Error(tt, err, "Expected error")
		})
	}
}
