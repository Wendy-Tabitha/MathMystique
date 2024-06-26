package mathskills

import (
	"math"
	"reflect"
	"sort"
	"testing"
)

func TestMedian(t *testing.T) {
	input := []float64{10, 30, 20, 50, 40}
	expected := 30.0

	sort.Float64s(input)
	result := Median(input)
	roundOutput := math.Round(result)

	if !reflect.DeepEqual(roundOutput, expected) {
		t.Errorf("Expected %v, but got %v", expected, roundOutput)
	}
}
