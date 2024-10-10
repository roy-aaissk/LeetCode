package convertthetemperature_test

import (
	convertthetemperature "leet_code/algorithms/2469.ConverttheTemperature"
	"reflect"
	"testing"
)

func Test_ConvertTemperature(t *testing.T) {
	tests := []struct {
		input    float64
		expected []float64
	}{
		{36.50, []float64{309.65000, 97.70000}},
		{122.11, []float64{395.26000, 251.79800}},
	}
	for _, test := range tests {
		if result := convertthetemperature.ConvertTemperature(test.input); !reflect.DeepEqual(result, test.expected) {
			t.Errorf("ConvertTemperature(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}
